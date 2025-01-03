// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_eu

import (
	"strings"
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

func PtrToEntity(record *CSLRecord) search.Entity[CSLRecord] {
	if record != nil {
		return ToEntity(*record)
	}
	return search.Entity[CSLRecord]{}
}

func ToEntity(record CSLRecord) search.Entity[CSLRecord] {
	out := search.Entity[CSLRecord]{
		Source:     search.SourceEUCSL,
		SourceData: record,
	}

	if strings.EqualFold(record.EntitySubjectType, "person") {
		out.Type = search.EntityPerson
		out.Person = &search.Person{}

		if len(record.NameAliasWholeNames) > 0 {
			out.Name = record.NameAliasWholeNames[0]
			out.Person.Name = record.NameAliasWholeNames[0]
			out.Person.AltNames = record.NameAliasWholeNames[1:]
		}
		if len(record.BirthDates) > 0 {
			tt, err := time.Parse("2006-01-02", record.BirthDates[0])
			if err == nil {
				out.Person.BirthDate = &tt
			}
		}
	}

	return out
}
