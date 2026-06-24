package refresh

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/internal/search"
)

// Trigger labels describe what initiated a refresh. They are recorded on Status.
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

// Status is a snapshot of the current or most recent data refresh. It reports the
// refresh operation itself; per-list entity counts are available at GET /v2/listinfo.
type Status struct {
	State   State  `json:"state"`
	Trigger string `json:"trigger,omitempty"`

	StartedAt  *time.Time `json:"startedAt,omitempty"`
	FinishedAt *time.Time `json:"finishedAt,omitempty"` // nil while a refresh is running
	Duration   string     `json:"duration,omitempty"`   // wall-clock of the last completed run

	LastError string `json:"lastError,omitempty"`
}

// Manager orchestrates full data refreshes and tracks their status. A Manager
// serializes refreshes: only one runs at a time regardless of whether it was
// triggered on startup, on a schedule, or on demand.
type Manager interface {
	// RefreshNow runs a full refresh synchronously, blocking until it completes.
	// It returns ErrAlreadyRunning if a refresh is already in progress.
	RefreshNow(ctx context.Context, trigger string) error

	// TriggerAsync starts a full refresh in the background using the Manager's
	// base context. It returns false if a refresh is already in progress.
	TriggerAsync(trigger string) bool

	// Status returns a snapshot of the current or most recent refresh.
	Status() Status
}

func NewManager(baseCtx context.Context, logger log.Logger, downloader download.Downloader, indexedLists index.Lists, searchService search.Service) Manager {
	return &manager{
		baseCtx:       baseCtx,
		logger:        logger,
		downloader:    downloader,
		indexedLists:  indexedLists,
		searchService: searchService,
		status:        Status{State: StateIdle},
	}
}

type manager struct {
	// baseCtx is the long-lived server context used for background refreshes so
	// they outlive the HTTP request that triggered them.
	baseCtx context.Context

	logger        log.Logger
	downloader    download.Downloader
	indexedLists  index.Lists
	searchService search.Service

	mu      sync.Mutex
	running bool
	status  Status
}

func (m *manager) RefreshNow(ctx context.Context, trigger string) error {
	start, ok := m.start(trigger)
	if !ok {
		return ErrAlreadyRunning
	}

	err := m.runRefresh(ctx)
	m.finish(start, err)
	return err
}

func (m *manager) TriggerAsync(trigger string) bool {
	start, ok := m.start(trigger)
	if !ok {
		return false
	}

	go func() {
		err := m.runRefresh(m.baseCtx)
		m.finish(start, err)
		if err != nil {
			m.logger.Error().LogErrorf("background %s data refresh failed: %v", trigger, err)
		}
	}()
	return true
}

func (m *manager) Status() Status {
	m.mu.Lock()
	defer m.mu.Unlock()

	// The *time.Time fields are only ever reassigned (never mutated in place), so
	// the by-value struct copy returned here is a safe snapshot.
	return m.status
}

// start atomically claims the single refresh slot. It returns the start time and
// true if claimed, or the zero time and false if a refresh is already running.
func (m *manager) start(trigger string) (time.Time, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.running {
		return time.Time{}, false
	}

	m.running = true
	now := time.Now().In(time.UTC)
	m.status = Status{
		State:     StateRunning,
		Trigger:   trigger,
		StartedAt: &now,
	}
	return now, true
}

// finish records the result of a refresh and releases the slot.
func (m *manager) finish(start time.Time, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	end := time.Now().In(time.UTC)
	m.running = false
	m.status.FinishedAt = &end
	m.status.Duration = end.Sub(start).String()

	if err != nil {
		m.status.State = StateFailed
		m.status.LastError = err.Error()
		return
	}

	m.status.State = StateSucceeded
	m.status.LastError = ""
}

// runRefresh performs the download, in-memory index swap, and embedding rebuild.
func (m *manager) runRefresh(ctx context.Context) error {
	ctx, span := telemetry.StartSpan(ctx, "refresh-all-sources")
	defer span.End()

	stats, err := m.downloader.RefreshAll(ctx)
	if err != nil {
		return err
	}

	m.logger.Info().Logf("data refreshed - %v entities from %v lists took %v (using %.2fGB)",
		len(stats.Entities), len(stats.Lists), stats.EndedAt.Sub(stats.StartedAt), getCurrentMemoryUsed())

	// Replace in-mem entities
	m.indexedLists.Update(stats)

	// Rebuild embedding index if enabled
	if m.searchService != nil {
		if err := m.searchService.RebuildEmbeddingIndex(ctx); err != nil {
			return fmt.Errorf("failed to rebuild embedding index: %w", err)
		}
	}

	return nil
}

func getCurrentMemoryUsed() float64 {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return float64(mem.Alloc) / 1024.0 / 1024.0 / 1024.0 // divide by 1GB
}
