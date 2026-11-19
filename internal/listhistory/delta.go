// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/moov-io/watchman/pkg/search"
)

func hashEntity(entity search.Entity[search.Value]) (string, error) {
	_, hash, err := marshalEntity(entity)
	return hash, err
}

func marshalEntity(entity search.Entity[search.Value]) ([]byte, string, error) {
	raw, err := json.Marshal(entity)
	if err != nil {
		return nil, "", fmt.Errorf("marshal entity %s: %w", entity.SourceID, err)
	}
	sum := sha256.Sum256(raw)
	return raw, hex.EncodeToString(sum[:]), nil
}

func decodeEntity(raw string) (search.Entity[search.Value], error) {
	var entity search.Entity[search.Value]
	if err := json.Unmarshal([]byte(raw), &entity); err != nil {
		return entity, fmt.Errorf("decode entity: %w", err)
	}
	return entity.Normalize(), nil
}

func computeDelta(oldHashes map[string]string, incoming []search.Entity[search.Value]) (Delta, error) {
	byID := make(map[string]search.Entity[search.Value], len(incoming))
	order := make([]string, 0, len(incoming))
	for i := range incoming {
		id := incoming[i].SourceID
		if id == "" || len(id) > 255 {
			continue
		}
		if _, exists := byID[id]; !exists {
			order = append(order, id)
		}
		byID[id] = incoming[i]
	}

	var delta Delta
	seen := make(map[string]struct{}, len(byID))
	for _, id := range order {
		entity := byID[id]
		seen[id] = struct{}{}
		hash, err := hashEntity(entity)
		if err != nil {
			return Delta{}, err
		}
		oldHash, exists := oldHashes[id]
		switch {
		case !exists:
			delta.Added = append(delta.Added, entity)
		case oldHash != hash:
			delta.Changed = append(delta.Changed, entity)
		}
	}
	for id := range oldHashes {
		if _, ok := seen[id]; !ok {
			delta.Removed = append(delta.Removed, id)
		}
	}
	slices.Sort(delta.Removed)
	return delta, nil
}

func truncateName(s string) string {
	if utf8.RuneCountInString(s) <= maxNameRunes {
		return s
	}
	runes := []rune(s)
	return string(runes[:maxNameRunes])
}

const likeEscape = "!"

func likeContains(q string) string {
	q = strings.ReplaceAll(q, likeEscape, likeEscape+likeEscape)
	q = strings.ReplaceAll(q, "%", likeEscape+"%")
	q = strings.ReplaceAll(q, "_", likeEscape+"_")
	return "%" + strings.ToLower(q) + "%"
}

func validEntityType(s string) bool {
	if s == "" {
		return true
	}
	switch search.EntityType(s) {
	case search.EntityUnknown, search.EntityPerson, search.EntityBusiness, search.EntityOrganization, search.EntityAircraft, search.EntityVessel:
		return true
	default:
		return false
	}
}

func matchesQuery(name, sourceID, q string) bool {
	if q == "" {
		return true
	}
	q = strings.ToLower(q)
	return strings.Contains(strings.ToLower(name), q) || strings.Contains(strings.ToLower(sourceID), q)
}
