package download

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
)

// Trigger labels describe what initiated a refresh.
const (
	TriggerStartup   = "startup"
	TriggerScheduled = "scheduled"
	TriggerManual    = "manual"
)

// ErrAlreadyRunning is returned when a refresh is requested while another is in progress.
var ErrAlreadyRunning = errors.New("a data refresh is already running")

// State describes the lifecycle of the current or most recent refresh.
type State string

const (
	StateIdle      State = "idle"
	StateRunning   State = "running"
	StateSucceeded State = "succeeded"
	StateFailed    State = "failed"
)

// Status is a snapshot of the current or most recent data refresh.
type Status struct {
	State   State  `json:"state"`
	Trigger string `json:"trigger,omitempty"`

	StartedAt  *time.Time `json:"startedAt,omitempty"`
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	Duration   string     `json:"duration,omitempty"`
	LastError  string     `json:"lastError,omitempty"`
}

// statsUpdater receives download stats after a successful refresh (e.g. index.Lists).
type statsUpdater interface {
	Update(Stats)
}

// embeddingRebuilder rebuilds optional embedding indexes after refresh.
type embeddingRebuilder interface {
	RebuildEmbeddingIndex(ctx context.Context) error
}

// Refresher coordinates full data refreshes with single-flight semantics and status tracking.
type Refresher struct {
	baseCtx context.Context
	logger  log.Logger

	dl            Downloader
	indexedLists  statsUpdater
	searchService embeddingRebuilder

	mu      sync.Mutex
	running bool
	status  Status
}

func NewRefresher(baseCtx context.Context, logger log.Logger, dl Downloader, su statsUpdater, es embeddingRebuilder) *Refresher {
	return &Refresher{
		baseCtx:       baseCtx,
		logger:        logger,
		dl:            dl,
		indexedLists:  su,
		searchService: es,
		status:        Status{State: StateIdle},
	}
}

// RefreshNow runs a full refresh synchronously. Returns ErrAlreadyRunning if one is in progress.
func (r *Refresher) RefreshNow(ctx context.Context, trigger string) error {
	start, ok := r.start(trigger)
	if !ok {
		return ErrAlreadyRunning
	}
	err := r.run(ctx)
	r.finish(start, err)
	return err
}

// TriggerAsync starts a refresh in the background. Returns false if one is already running.
func (r *Refresher) TriggerAsync(trigger string) bool {
	start, ok := r.start(trigger)
	if !ok {
		return false
	}
	go func() {
		err := r.run(r.baseCtx)
		r.finish(start, err)
		if err != nil {
			r.logger.Error().LogErrorf("background %s data refresh failed: %v", trigger, err)
		}
	}()
	return true
}

// Status returns a snapshot of the current or most recent refresh.
func (r *Refresher) Status() Status {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.status
}

func (r *Refresher) start(trigger string) (time.Time, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.running {
		return time.Time{}, false
	}
	r.running = true
	now := time.Now().UTC()
	r.status = Status{
		State:     StateRunning,
		Trigger:   trigger,
		StartedAt: &now,
	}
	return now, true
}

func (r *Refresher) finish(start time.Time, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	end := time.Now().UTC()
	r.running = false
	r.status.FinishedAt = &end
	r.status.Duration = end.Sub(start).String()
	if err != nil {
		r.status.State = StateFailed
		r.status.LastError = err.Error()
		return
	}
	r.status.State = StateSucceeded
	r.status.LastError = ""
}

func (r *Refresher) run(ctx context.Context) error {
	ctx, span := telemetry.StartSpan(ctx, "refresh-all-sources")
	defer span.End()

	stats, err := r.dl.RefreshAll(ctx)
	if err != nil {
		return err
	}

	r.logger.Info().Logf("data refreshed - %v entities from %v lists took %v (using %.2fGB)",
		len(stats.Entities), len(stats.Lists), stats.EndedAt.Sub(stats.StartedAt), getCurrentMemoryUsed())

	if r.indexedLists != nil {
		r.indexedLists.Update(stats)
	}
	if r.searchService != nil {
		if err := r.searchService.RebuildEmbeddingIndex(ctx); err != nil {
			return fmt.Errorf("failed to rebuild embedding index: %w", err)
		}
	}
	return nil
}

func getCurrentMemoryUsed() float64 {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return float64(mem.Alloc) / 1024 / 1024 / 1024
}
