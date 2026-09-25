// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package recordlink emits PII-safe composite blocking keys for Watchman entities.
//
// Use Keys to bucket similar records before Similarity scoring (SQL prefix
// filters or an external index). Prefixes walks a key from coarse to fine.
// The keys never contain names, identifiers, addresses, emails, or phones.
package recordlink

import (
	"github.com/moov-io/watchman/internal/linksim"
	"github.com/moov-io/watchman/pkg/search"
)

// Kind prefixes. The kind is the only cleartext in a key; field values are truncated digests.
const (
	KindType    = linksim.KindType
	KindName    = linksim.KindName
	KindGovID   = linksim.KindGovID
	KindAddr    = linksim.KindAddr
	KindContact = linksim.KindContact
	KindIMO     = linksim.KindIMO
	KindMMSI    = linksim.KindMMSI
	KindAir     = linksim.KindAir
)

// Keys returns hashed blocking keys for e. The entity is Normalize()'d first.
func Keys(e search.Entity[search.Value]) []string {
	return linksim.Keys(e.Normalize())
}

// Prefixes returns the coarse-to-fine segment prefixes of a composite key,
// including the key itself.
//
//	ADDR:Caa|Sbb|Pcc → [ADDR:Caa, ADDR:Caa|Sbb, ADDR:Caa|Sbb|Pcc]
func Prefixes(key string) []string {
	return linksim.Prefixes(key)
}
