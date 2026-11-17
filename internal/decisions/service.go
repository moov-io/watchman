package decisions

import (
	"context"
	"fmt"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Service annotates search hits with a Jev-style decision model.
// Calls are fail-open: a model error leaves Decision unset and search still returns.
type Service struct {
	logger log.Logger
	config Config
	client *client
}

// NewService returns nil when the decision model is disabled.
func NewService(logger log.Logger, config Config) (*Service, error) {
	config.LoadFromEnv()

	if !config.Enabled {
		return nil, nil
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("decisions: invalid config: %w", err)
	}

	logger.Info().Logf("decisions: using model %s at %s", config.Model, systemOneURL(config.BaseURL))

	return &Service{
		logger: logger,
		config: config,
		client: newClient(config),
	}, nil
}

// Annotate attaches decision-model answers onto results. It never returns an
// error to the caller; failures are recorded on the span and logged.
func (s *Service) Annotate(ctx context.Context, query search.Entity[search.Value], results []search.SearchedEntity[search.Value]) {
	if s == nil || len(results) == 0 {
		return
	}

	selected := selectHits(results, s.config)
	if len(selected) == 0 {
		return
	}

	ctx, span := telemetry.StartSpan(ctx, "search-decisions", trace.WithAttributes(
		attribute.String("decisions.model", s.config.Model),
		attribute.Int("decisions.entity_count", len(selected)),
		attribute.Int("decisions.question_count", len(selected)*3),
	))
	defer span.End()

	hits := make([]search.SearchedEntity[search.Value], len(selected))
	for i, idx := range selected {
		hits[i] = results[idx]
	}

	resp, err := s.client.decide(ctx, Request{
		Model:     s.config.Model,
		State:     buildState(query, hits),
		Questions: screeningQuestions(len(hits)),
	})
	if err != nil {
		s.logger.Warn().Logf("decision model request failed, returning search without decisions: %v", err)
		span.SetAttributes(attribute.Bool("decisions.failed", true))
		return
	}

	model := resp.Model
	if model == "" {
		model = s.config.Model
	}

	var annotated int
	for i, idx := range selected {
		decision := applyAnswers(resp.Answers, i, model)
		if decision == nil {
			continue
		}
		results[idx].Decision = decision
		annotated++
	}
	span.SetAttributes(attribute.Int("decisions.annotated_count", annotated))
}

func selectHits(results []search.SearchedEntity[search.Value], conf Config) []int {
	limit := conf.maxEntitiesOrDefault()
	out := make([]int, 0, min(len(results), limit))
	for i, result := range results {
		if result.Match < conf.MinMatch {
			continue
		}
		out = append(out, i)
		if len(out) >= limit {
			break
		}
	}
	return out
}

func applyAnswers(answers map[string]Answer, index int, model string) *search.EntityDecision {
	if len(answers) == 0 {
		return nil
	}

	same, hasSame := answers[sameEntityKey(index)]
	action, hasAction := answers[actionKey(index)]
	quality, hasQuality := answers[qualityKey(index)]
	if !hasSame && !hasAction && !hasQuality {
		return nil
	}

	out := &search.EntityDecision{Model: model}
	if hasSame {
		out.SameEntity = same.Noul
	}
	if hasAction {
		out.Action = action.Choice
		out.ActionProbabilities = action.Probabilities
		out.Confidence = action.Confidence
	}
	if hasQuality {
		out.MatchQuality = quality.Score
		if out.Confidence == 0 {
			out.Confidence = quality.Confidence
		}
	}
	return out
}
