package postalpool

import (
	"cmp"
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/moov-io/watchman/internal/postalpool/coder"
	"github.com/moov-io/watchman/pkg/address"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
)

// Client provides thread-safe access to a pool of libpostal processes,
// distributing requests across them in a round-robin fashion. Each request
// is automatically routed to the next available worker process.
//
// Requests will be queued up on workers and start to block if every worker
// has an active request. libpostal is single threaded.
//
// Client maintains an address pool to the worker processes. It provides a
// simple API that hides the complexity of working with multiple postal processes.
type Client struct {
	conf      Config
	endpoints []string
	next      atomic.Uint32

	httpClient *http.Client
}

func NewClient(conf Config, endpoints []string) *Client {
	return &Client{
		conf:       conf,
		endpoints:  endpoints,
		httpClient: defaultHttpClient(conf),
	}
}

func defaultHttpClient(conf Config) *http.Client {
	dialer := cmp.Or(conf.Dialer, &net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 60 * time.Second,
	})
	transport := cmp.Or(conf.Transport, &http.Transport{
		MaxIdleConnsPerHost: 20,
		MaxIdleConns:        200,
		IdleConnTimeout:     5 * time.Minute,
		DialContext:         dialer.DialContext,
	})
	return &http.Client{
		Transport: transport,
		Timeout:   cmp.Or(conf.RequestTimeout, 10*time.Second),
	}
}

func (c *Client) ParseAddress(ctx context.Context, input string) (search.Address, error) {
	return c.parseAddress(ctx, input, true)
}

func (c *Client) parseAddress(ctx context.Context, input string, includeCGOSelf bool) (search.Address, error) {
	ctx, span := telemetry.StartSpan(ctx, "postalpool-parse-address")
	defer span.End()

	if len(c.endpoints) == 0 {
		span.SetAttributes(attribute.String("postalpool.method", "cgo-zero"))
		return address.ParseAddress(ctx, input), nil
	}

	// Simple round-robin including self or not
	var offset int
	if includeCGOSelf {
		offset += c.conf.CGOSelfInstances
	}
	idx := int(c.next.Add(1)) % (len(c.endpoints) + offset)

	// If idx equals last position, use local instance
	if idx >= len(c.endpoints) {
		span.SetAttributes(attribute.String("postalpool.method", "cgo-self"))
		return address.ParseAddress(ctx, input), nil
	}

	endpoint := c.endpoints[idx]
	span.SetAttributes(
		attribute.String("postalpool.endpoint", endpoint),
		attribute.String("postalpool.method", "binary"),
	)

	var addr search.Address

	req, err := http.NewRequestWithContext(ctx, "GET", endpoint+"/parse?address="+url.QueryEscape(input), nil)
	if err != nil {
		return addr, fmt.Errorf("creating postal-server request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return addr, fmt.Errorf("HTTP parse to postal-server: %w", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	dec := coder.GetDecoder()
	dec.Reset(resp.Body)
	defer coder.SaveDecoder(dec)

	err = dec.Decode(&addr)
	if err != nil {
		return addr, fmt.Errorf("reading postal-server response: %w", err)
	}
	return addr, nil
}

const (
	healthCheckAddress = "123 1st st anytown ca 90210"
)

func (c *Client) healthcheck(ctx context.Context) error {
	ctx, span := telemetry.StartSpan(ctx, "postalpool-client-healthcheck")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, c.conf.StartupTimeout)
	defer cancel()

	results := make(chan error, len(c.endpoints))

	// Try each endpoint every 250ms until success
	for _, endpoint := range c.endpoints {
		go func(ep string) {
			ticker := time.NewTicker(250 * time.Millisecond)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					results <- fmt.Errorf("healthcheck timed out for %s", ep)
					return

				case <-ticker.C:
					addr, err := c.parseAddress(ctx, healthCheckAddress, false) // force network connections
					if addr.Format() != "" && err == nil {
						results <- nil
						return
					}
				}
			}
		}(endpoint)
	}

	// Wait for first success or all failures
	var lastErr error
	for i := 0; i < len(c.endpoints); i++ {
		if err := <-results; err == nil {
			return nil // Found a working endpoint
		} else {
			lastErr = err
		}
	}

	return fmt.Errorf("all endpoints failed healthcheck: %w", lastErr)
}
