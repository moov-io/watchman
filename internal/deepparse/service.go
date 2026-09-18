package deepparse

import (
	"cmp"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/pkg/search"

	deepparsego "github.com/adamdecaf/deepparse-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	defaultBaseURL = "http://localhost:8000"
	defaultModel   = deepparsego.ModelBPEmbAttention
	healthAddress  = "123 First St Anytown CA 90210"
)

// Service talks to an external GRAAL-Research/deepparse FastAPI over HTTP.
type Service struct {
	client deepparsego.Client
	model  deepparsego.Model
}

// NewService returns nil when deepparse is disabled.
func NewService(logger log.Logger, conf Config) (*Service, error) {
	if !conf.Enabled {
		return nil, nil
	}

	baseURL := cmp.Or(os.Getenv("DEEPPARSE_URL"), conf.BaseURL, defaultBaseURL)
	model := deepparsego.Model(cmp.Or(conf.Model, string(defaultModel)))
	timeout := cmp.Or(conf.Timeout, 10*time.Second)

	logger.Info().Logf("using deepparse at %s model=%s", baseURL, model)

	svc := &Service{
		client: deepparsego.NewClient(&http.Client{Timeout: timeout}, baseURL),
		model:  model,
	}

	ctx, span := telemetry.StartSpan(context.Background(), "deepparse-setup", trace.WithAttributes(
		attribute.String("deepparse.base_url", baseURL),
		attribute.String("deepparse.model", string(model)),
	))
	defer span.End()

	if _, err := svc.ParseAddress(ctx, healthAddress); err != nil {
		return nil, fmt.Errorf("deepparse healthcheck: %w", err)
	}

	return svc, nil
}

func (s *Service) ParseAddress(ctx context.Context, input string) (search.Address, error) {
	ctx, span := telemetry.StartSpan(ctx, "parse-address", trace.WithAttributes(
		attribute.String("address.method", "deepparse"),
		attribute.String("deepparse.model", string(s.model)),
	))
	defer span.End()

	resp, err := s.client.ParseAddresses(ctx, s.model, []string{input})
	if err != nil {
		return search.Address{}, err
	}
	if len(resp.Addresses) == 0 {
		return search.Address{}, fmt.Errorf("deepparse returned no addresses")
	}

	return mapParsed(resp.Addresses[0]), nil
}

func mapParsed(in deepparsego.ParsedAddress) search.Address {
	line1 := joinNonEmpty([]string{in.StreetNumber, in.StreetName, in.Orientation}, " ")
	line2 := joinNonEmpty([]string{in.Unit, in.GeneralDelivery}, ", ")

	return search.Address{
		Line1:      line1,
		Line2:      line2,
		City:       in.Municipality,
		State:      in.Province,
		PostalCode: in.PostalCode,
	}
}

func joinNonEmpty(parts []string, sep string) string {
	var nonEmpty []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			nonEmpty = append(nonEmpty, p)
		}
	}
	return strings.Join(nonEmpty, sep)
}
