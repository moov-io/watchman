// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"time"

	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/pkg/search"
)

func testPerson(id, name string) search.Entity[search.Value] {
	return search.Entity[search.Value]{
		Name:     name,
		Type:     search.EntityPerson,
		Source:   search.SourceUSOFAC,
		SourceID: id,
		Person:   &search.Person{Name: name},
	}
}

func testBusiness(id, name string) search.Entity[search.Value] {
	return search.Entity[search.Value]{
		Name:     name,
		Type:     search.EntityBusiness,
		Source:   search.SourceUSOFAC,
		SourceID: id,
		Business: &search.Business{Name: name},
	}
}

func testStats(hash string, at time.Time, ents ...search.Entity[search.Value]) download.Stats {
	return download.Stats{
		Entities:   ents,
		Lists:      map[string]int{string(search.SourceUSOFAC): len(ents)},
		ListHashes: map[string]string{string(search.SourceUSOFAC): hash},
		EndedAt:    at,
	}
}

func testTime(sec int) time.Time {
	return time.Date(2024, 6, 1, 12, 0, sec, 0, time.UTC)
}
