// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package linksim

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

const (
	// hashBytes is the truncated SHA-256 length used in keys (8 hex chars).
	// This is a blocking digest, not a unique identifier: collisions only add
	// extra candidates, which scoring still filters.
	hashBytes = 4

	// domain separates Watchman blocking hashes from other SHA-256 uses and
	// from each other across fields (country "us" vs name token "us").
	domain = "watchman/linksim/v1/"
)

// digest returns an 8-character lowercase hex truncation of SHA-256 over a
// domain-separated field/value pair. Empty or whitespace-only input yields "".
//
// The digest is one-way enough that keys do not contain names, ID numbers, or
// addresses. It is not a password hash: anyone who can guess the input can
// recompute the digest. Field names are mixed in so the same string hashes
// differently as a country vs an identifier.
func digest(field, value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	sum := sha256.Sum256([]byte(domain + field + "\x00" + value))
	return hex.EncodeToString(sum[:hashBytes])
}
