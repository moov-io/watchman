#!/usr/bin/env bash
# Collect live Watchman memory artifacts from the admin server (default :9094).
#
# Usage:
#   ./internal/corpusfile/collect_profiles.sh [admin_base_url] [out_dir]
#
# Example:
#   ./internal/corpusfile/collect_profiles.sh http://127.0.0.1:9094 ./profiles
#
# What helps most for corpus memory work:
#   1. heap inuse_space after refresh settles (what is retained)
#   2. allocs alloc_space during a refresh (what refresh costs)
#   3. /debug/memory?gc=1 JSON snapshot (entity counts + MemStats)
#   4. optional: goroutine profile if RSS looks stuck on non-heap

set -euo pipefail

ADMIN="${1:-http://127.0.0.1:9094}"
OUT="${2:-./watchman-profiles-$(date +%Y%m%d-%H%M%S)}"
mkdir -p "$OUT"

echo "Collecting from $ADMIN -> $OUT"

curl -fsS "$ADMIN/debug/memory?gc=1" -o "$OUT/memory.json"
curl -fsS "$ADMIN/debug/pprof/heap" -o "$OUT/heap.pb.gz"
curl -fsS "$ADMIN/debug/pprof/allocs" -o "$OUT/allocs.pb.gz"
curl -fsS "$ADMIN/debug/pprof/goroutine" -o "$OUT/goroutine.pb.gz"
# Short CPU sample; increase seconds under load if needed.
curl -fsS "$ADMIN/debug/pprof/profile?seconds=15" -o "$OUT/cpu.pb.gz" || true

# Text top for quick glance without opening the UI
if command -v go >/dev/null 2>&1; then
  go tool pprof -top -inuse_space "$OUT/heap.pb.gz" >"$OUT/heap_inuse_top.txt" || true
  go tool pprof -top -alloc_space "$OUT/allocs.pb.gz" >"$OUT/allocs_top.txt" || true
fi

cat >"$OUT/README.txt" <<EOF
Watchman memory dump
admin=$ADMIN
collected=$(date -u +%Y-%m-%dT%H:%M:%SZ)

View interactively:
  go tool pprof -http=:8081 $OUT/heap.pb.gz
  # UI: Top, then sample=inuse_space (retained) vs alloc_space (cumulative)

What to look for:
  - search.Entity / PreparedFields / ofac.SDN / map[string][]int (indexes)
  - libpostal / postal if enabled
  - embeddings vectorData if enabled
  - unexpected large buffers from download/parse retained after refresh

Share memory.json + heap_inuse_top.txt + heap.pb.gz for corpus work.
EOF

echo "Done. Files:"
ls -la "$OUT"
