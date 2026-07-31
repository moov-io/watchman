package index

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"

	"github.com/moov-io/base/admin"
)

// MemoryReport is exposed on the admin server for live RSS / corpus sizing.
// Pair with pprof heap profiles from the same admin port for allocation detail.
type MemoryReport struct {
	Time    time.Time `json:"time"`
	Version string    `json:"version,omitempty"`

	// Process memory (Go runtime)
	HeapAlloc  uint64 `json:"heapAlloc"`
	HeapSys    uint64 `json:"heapSys"`
	HeapInuse  uint64 `json:"heapInuse"`
	HeapIdle   uint64 `json:"heapIdle"`
	StackInuse uint64 `json:"stackInuse"`
	Sys        uint64 `json:"sys"`
	NumGC      uint32 `json:"numGC"`
	GCCPUFraction float64 `json:"gcCPUFraction"`

	// Corpus summary from the last successful refresh
	EntityCount int            `json:"entityCount"`
	Lists       map[string]int `json:"lists,omitempty"`
	RefreshStartedAt time.Time `json:"refreshStartedAt,omitempty"`
	RefreshEndedAt   time.Time `json:"refreshEndedAt,omitempty"`

	// How to pull deeper profiles from this admin server
	PprofHints PprofHints `json:"pprofHints"`
}

// PprofHints documents the always-on admin pprof routes (moov-io/base/admin).
type PprofHints struct {
	Heap      string `json:"heap"`
	Allocs    string `json:"allocs"`
	Goroutine string `json:"goroutine"`
	Profile   string `json:"cpuProfile"`
	Note      string `json:"note"`
}

// RegisterMemoryAdmin adds GET /debug/memory with live MemStats + corpus counts.
func RegisterMemoryAdmin(svc *admin.Server, lists Lists) {
	if svc == nil || lists == nil {
		return
	}
	svc.AddHandler("/debug/memory", memoryHandler(lists))
}

func memoryHandler(lists Lists) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Optional GC so heap numbers are less noisy when debugging.
		if r.URL.Query().Get("gc") == "1" {
			runtime.GC()
		}

		var ms runtime.MemStats
		runtime.ReadMemStats(&ms)
		stats := lists.LatestStats()

		// Prefer corpus length when built; fall back to summing list counts.
		entityCount := 0
		if n := EntityCount(lists); n > 0 {
			entityCount = n
		} else {
			for _, n := range stats.Lists {
				entityCount += n
			}
		}

		report := MemoryReport{
			Time:    time.Now().UTC(),
			Version: stats.Version,
			HeapAlloc:     ms.HeapAlloc,
			HeapSys:       ms.HeapSys,
			HeapInuse:     ms.HeapInuse,
			HeapIdle:      ms.HeapIdle,
			StackInuse:    ms.StackInuse,
			Sys:           ms.Sys,
			NumGC:         ms.NumGC,
			GCCPUFraction: ms.GCCPUFraction,
			EntityCount:   entityCount,
			Lists:         stats.Lists,
			RefreshStartedAt: stats.StartedAt,
			RefreshEndedAt:   stats.EndedAt,
			PprofHints: PprofHints{
				Heap:      "/debug/pprof/heap",
				Allocs:    "/debug/pprof/allocs",
				Goroutine: "/debug/pprof/goroutine",
				Profile:   "/debug/pprof/profile?seconds=30",
				Note: "curl -o heap.pb.gz http://ADMIN/debug/pprof/heap && go tool pprof -http=:8081 heap.pb.gz. " +
					"Most useful: inuse_space after a refresh settles; alloc_space for refresh cost. " +
					"Also GET /debug/memory?gc=1 for a post-GC snapshot.",
			},
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		_ = enc.Encode(report)
	}
}
