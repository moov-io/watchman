package postalpool

import (
	"cmp"
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
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
		processes: make([]*exec.Cmd, conf.Instances),
	}

	endpoints := make([]string, conf.Instances)
	for i := 0; i < conf.Instances; i++ {
		port := conf.StartingPort + i

		cmd := exec.Command(binPath)
		cmd.Env = append(cmd.Env, fmt.Sprintf("PORT=%d", port))

		err := cmd.Start()
		if err != nil {
			ps.Shutdown()

			return nil, fmt.Errorf("failed to start postal instance %d: %w", i, err)
		}

		ps.processes[i] = cmd

		endpoints[i] = fmt.Sprintf("http://localhost:%d", port)
	}

	ps.client = NewClient(conf, endpoints)

	err := ps.client.healthcheck(ctx)
	if err != nil {
		return nil, fmt.Errorf("problem with postalpool healthcheck: %w", err)
	}

	return ps, nil
}

func (ps *Service) Shutdown() {
	for _, proc := range ps.processes {
		if proc != nil {
			proc.Process.Kill()
		}
	}
}

func (ps *Service) ParseAddress(ctx context.Context, input string) (search.Address, error) {
	return ps.client.ParseAddress(ctx, input)
}
