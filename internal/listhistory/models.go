// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

const (
	RefLatest = "latest"

	ChangeAdded   = "added"
	ChangeRemoved = "removed"
	ChangeChanged = "changed"

	DetailSummary = "summary"
	DetailFull    = "full"

	maxNameRunes = 1024
	insertChunk  = 100
	closeChunk   = 200
)

// Ref is one stored snapshot of a downloaded list. Ref is the list file hash.
type Ref struct {
	ID           int64     `json:"-"`
	ListName     string    `json:"listName"`
	Ref          string    `json:"ref"`
	DownloadedAt time.Time `json:"downloadedAt"`
	EntityCount  int       `json:"entityCount"`
	AddedCount   int       `json:"addedCount"`
	RemovedCount int       `json:"removedCount"`
	ChangedCount int       `json:"changedCount"`
}

// EntitySummary is a list row at a ref without the full entity body.
type EntitySummary struct {
	SourceID   string            `json:"sourceID"`
	Name       string            `json:"name"`
	EntityType search.EntityType `json:"entityType"`
	EntityHash string            `json:"entityHash,omitempty"`
}

// EntityRecord is a list row at a ref, including the prepared entity.
type EntityRecord struct {
	EntitySummary
	Entity search.Entity[search.Value] `json:"entity"`
}

// EntityFilter selects rows when listing a snapshot.
type EntityFilter struct {
	Query      string
	EntityType string
	After      string
	Limit      int
}

// Delta is the per-entity change set from the previous snapshot to a new one.
type Delta struct {
	Added   []search.Entity[search.Value]
	Changed []search.Entity[search.Value]
	Removed []string
}

// ApplyParams is one list's new snapshot and its delta against the previous ref.
type ApplyParams struct {
	ListName     string
	ListHash     string
	DownloadedAt time.Time
	EntityCount  int
	Delta        Delta
}

// ListsResponse is GET /v2/lists.
type ListsResponse struct {
	Lists []Ref `json:"lists"`
}

// RefsResponse is GET /v2/lists/{listName}/refs.
type RefsResponse struct {
	ListName string `json:"listName"`
	Refs     []Ref  `json:"refs"`
}

// EntitiesResponse is GET /v2/lists/{listName}/entities.
type EntitiesResponse struct {
	ListName string         `json:"listName"`
	Ref      Ref            `json:"ref"`
	Entities []EntityRecord `json:"entities"`
	After    string         `json:"after,omitempty"`
}

// Change is one added, removed, or changed entity in a diff.
type Change struct {
	ChangeType string                       `json:"changeType"`
	SourceID   string                       `json:"sourceID"`
	Name       string                       `json:"name"`
	EntityType search.EntityType            `json:"entityType"`
	Before     *search.Entity[search.Value] `json:"before,omitempty"`
	After      *search.Entity[search.Value] `json:"after,omitempty"`
}

// DiffCounts is the full (unpaginated) size of a diff.
type DiffCounts struct {
	Added   int `json:"added"`
	Removed int `json:"removed"`
	Changed int `json:"changed"`
}

// DiffResponse is GET /v2/lists/{listName}/diff.
type DiffResponse struct {
	ListName string     `json:"listName"`
	From     Ref        `json:"from"`
	To       Ref        `json:"to"`
	Counts   DiffCounts `json:"counts"`
	Added    []Change   `json:"added"`
	Removed  []Change   `json:"removed"`
	Changed  []Change   `json:"changed"`
	After    string     `json:"after,omitempty"`
}
