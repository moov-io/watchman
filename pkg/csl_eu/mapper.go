package csl

import (
	"strings"
	"time"

	"github.com/moov-io/watchman/pkg/csl"
	"github.com/moov-io/watchman/pkg/search"
)

func PtrToEntity(record *csl.EUCSLRecord) search.Entity[csl.EUCSLRecord] {
	if record != nil {
		return ToEntity(*record)
	}
	return search.Entity[csl.EUCSLRecord]{}
}

func ToEntity(record csl.EUCSLRecord) search.Entity[csl.EUCSLRecord] {
	out := search.Entity[csl.EUCSLRecord]{
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
