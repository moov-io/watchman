package postalpool

import (
	"cmp"
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/moov-io/watchman/pkg/address"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// Service spawns a pool of libpostal process instances. It handles running and
// graceful shutdown of worker processes. Each worker process has its own isolated
// copy of libpostal's global state, allowing for concurrent processing without contention.
//
// A Service should be created via NewService() and stopped using the Shutdown() method
// when no longer needed. The Service ensures all worker processes are properly
// terminated on shutdown.
type Service struct {
	conf      Config
	processes []*exec.Cmd
	client    *Client
}

func NewService(logger log.Logger, conf Config) (*Service, error) {
	if !conf.Enabled {
		return nil, nil
	}

	ctx, span := telemetry.StartSpan(context.Background(), "postalpool-setup", trace.WithAttributes(
		attribute.String("binary_path", conf.BinaryPath),
		attribute.Int("workers", conf.Instances),
	))
	defer span.End()

	binPath := cmp.Or(os.Getenv("POSTAL_SERVER_BIN_PATH"), conf.BinaryPath, "./bin/postal-server")
	logger.Info().Logf("starting %v with %v instances starting at %v", binPath, conf.Instances, conf.StartingPort)

	ps := &Service{
		conf:      conf,
		processes: make([]*exec.Cmd, conf.Instances),
	}
	if conf.Instances <= 0 {
		return ps, nil
	}

	var g errgroup.Group
	endpoints := make([]string, conf.Instances)
	for i := 0; i < conf.Instances; i++ {
		idx := i

		port := conf.StartingPort + idx
		endpoints[idx] = fmt.Sprintf("http://127.0.0.1:%d", port)

		g.Go(func() error {
			cmd := exec.Command(binPath) //nolint:gosec
			cmd.Env = append(cmd.Env, fmt.Sprintf("PORT=%d", port))
			cmd.Env = append(cmd.Env, fmt.Sprintf("REQUEST_TIMEOUT=%v", conf.RequestTimeout))

			err := cmd.Start()
			if err != nil {
				ps.Shutdown()

				return fmt.Errorf("failed to start postal instance %d: %w", i, err)
			}

			ps.processes[idx] = cmd

			return nil
		})
	}
	err := g.Wait()
	if err != nil {
		return nil, fmt.Errorf("problem starting postalpool workers: %w", err)
	}

	ps.client = NewClient(conf, endpoints)

	err = ps.client.healthcheck(ctx)
	if err != nil {
		return nil, fmt.Errorf("problem with postalpool healthcheck: %w", err)
	}

	return ps, nil
}

func (ps *Service) Ratio() string {
	return fmt.Sprintf("%d cgo to %d binary", ps.conf.CGOSelfInstances, ps.conf.Instances)
}

func (ps *Service) Shutdown() {
	for _, proc := range ps.processes {
		if proc != nil {
			proc.Process.Kill()
		}
	}
}

func (ps *Service) ParseAddress(ctx context.Context, input string) (search.Address, error) {
	if ps.client == nil {
		return address.ParseAddress(ctx, input), nil
	}
	return ps.client.ParseAddress(ctx, input)
}
