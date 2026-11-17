package decisions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const maxDecisionBody = 1 << 20 // 1 MiB

// Request is a Jev / Ollama System One payload.
type Request struct {
	Model     string              `json:"model"`
	State     any                 `json:"state"`
	Questions map[string]Question `json:"questions"`
}

// Question is one typed decision about State.
type Question struct {
	Type         string `json:"type"`
	Instructions string `json:"instructions"`
	Criteria     any    `json:"criteria,omitempty"`
}

// Response is a System One result.
type Response struct {
	Model   string            `json:"model"`
	Answers map[string]Answer `json:"answers"`
	Usage   Usage             `json:"usage"`
}

// Answer is one typed System One answer.
type Answer struct {
	Type          string             `json:"type"`
	Choice        string             `json:"choice,omitempty"`
	Noul          float64            `json:"noul"`
	Score         float64            `json:"score"`
	Confidence    float64            `json:"confidence,omitempty"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
}

// Usage is token accounting from the decision model.
type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type client struct {
	http    *http.Client
	baseURL string
	apiKey  string
	model   string
}

func newClient(conf Config) *client {
	return &client{
		http: &http.Client{
			Timeout: conf.timeoutOrDefault(),
		},
		baseURL: conf.BaseURL,
		apiKey:  conf.APIKey,
		model:   conf.Model,
	}
}

func (c *client) decide(ctx context.Context, req Request) (*Response, error) {
	if req.Model == "" {
		req.Model = c.model
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return nil, fmt.Errorf("encoding decision request: %w", err)
	}

	ctx, span := telemetry.StartSpan(ctx, "decision-model-request", trace.WithAttributes(
		attribute.String("decisions.model", req.Model),
		attribute.Int("decisions.question_count", len(req.Questions)),
	))
	defer span.End()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, systemOneURL(c.baseURL), &buf)
	if err != nil {
		return nil, telemetry.RecordError(ctx, fmt.Errorf("creating decision request: %w", err))
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if c.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, telemetry.RecordError(ctx, fmt.Errorf("posting decision request: %w", err))
	}
	defer resp.Body.Close()

	span.SetAttributes(attribute.Int("http.status_code", resp.StatusCode))

	body, err := io.ReadAll(io.LimitReader(resp.Body, maxDecisionBody))
	if err != nil {
		return nil, telemetry.RecordError(ctx, fmt.Errorf("reading decision response: %w", err))
	}
	if resp.StatusCode != http.StatusOK {
		return nil, telemetry.RecordError(ctx, fmt.Errorf("%w: %d", ErrUnexpectedStatus, resp.StatusCode))
	}

	var out Response
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, telemetry.RecordError(ctx, fmt.Errorf("%w: %w", ErrInvalidResponse, err))
	}
	span.SetAttributes(
		attribute.Int("decisions.input_tokens", out.Usage.InputTokens),
		attribute.Int("decisions.output_tokens", out.Usage.OutputTokens),
		attribute.Int("decisions.answer_count", len(out.Answers)),
	)
	return &out, nil
}

func systemOneURL(baseURL string) string {
	base := strings.TrimRight(strings.TrimSpace(baseURL), "/")
	switch {
	case strings.HasSuffix(base, "/v1/systemone"):
		return base
	case strings.HasSuffix(base, "/v1"):
		return base + "/systemone"
	default:
		return base + "/v1/systemone"
	}
}
