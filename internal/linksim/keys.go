// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package linksim

import (
	"slices"
	"strings"

	"github.com/moov-io/watchman/internal/norm"
	"github.com/moov-io/watchman/internal/stringscore"
	"github.com/moov-io/watchman/pkg/search"
)

// Kind prefixes for composite blocking keys. The kind is the only cleartext
// in a key; every field value is a truncated digest.
const (
	KindType    = "TYPE"
	KindName    = "NAME"
	KindGovID   = "GOVID"
	KindAddr    = "ADDR"
	KindContact = "CONTACT"
	KindIMO     = "IMO"
	KindMMSI    = "MMSI"
	KindAir     = "AIR"
)

// Keys returns PII-safe composite blocking keys for e.
//
// Each key is KIND: then hashed field segments, coarse to fine:
//
//	TYPE:<digest(type)>
//	NAME:<digest(soundex(token))>          // one key per significant name token
//	GOVID:C<digest(country)>|T<digest(type)>|X<digest(identifier)>
//	ADDR:C<digest(country)>|S<digest(state)>|P<digest(postal)>|Y<digest(city)>|L<digest(line1)>
//	CONTACT:E<digest(email)>   CONTACT:P<digest(phone)>
//	IMO:<digest(imo)>   MMSI:<digest(mmsi)>   AIR:<digest(serial)>
//
// Segment boundaries are "|". Prefixes(key) walks them so a SQL index can
// filter ADDR:C…|S…|P… without storing the postal code, city, or street.
//
// Entities should be Normalize()'d first so prepared names and addresses are
// populated. Keys never contain the original name, identifier, address, email,
// or phone; see digest.
func Keys(e search.Entity[search.Value]) []string {
	var out []string

	if t := strings.TrimSpace(string(e.Type)); t != "" {
		out = append(out, joinKind(KindType, digest("type", strings.ToLower(t))))
	}

	out = appendNameKeys(out, e)
	out = appendGovIDKeys(out, e)
	out = appendAddressKeys(out, e)
	out = appendContactKeys(out, e)
	out = appendVesselAircraftKeys(out, e)

	slices.Sort(out)
	return slices.Compact(out)
}

// Prefixes returns the coarse-to-fine segment prefixes of a composite key,
// including the key itself. A key with no "|" has a single prefix (itself).
//
//	ADDR:Caa|Sbb|Pcc → [ADDR:Caa, ADDR:Caa|Sbb, ADDR:Caa|Sbb|Pcc]
func Prefixes(key string) []string {
	kind, payload, ok := strings.Cut(key, ":")
	if !ok || payload == "" {
		return nil
	}
	parts := strings.Split(payload, "|")
	out := make([]string, 0, len(parts))
	prefix := kind + ":"
	for i, p := range parts {
		if i == 0 {
			prefix += p
		} else {
			prefix += "|" + p
		}
		out = append(out, prefix)
	}
	return out
}

func joinKind(kind, payload string) string {
	if payload == "" {
		return ""
	}
	return kind + ":" + payload
}

func joinSegments(kind string, segs []string) string {
	if len(segs) == 0 {
		return ""
	}
	return kind + ":" + strings.Join(segs, "|")
}

func appendNameKeys(out []string, e search.Entity[search.Value]) []string {
	seen := make(map[string]struct{})
	addTokens := func(tokens []string) {
		for _, tok := range tokens {
			code := stringscore.EncodeSoundex(tok)
			if code == "" {
				continue
			}
			key := joinKind(KindName, digest("name-soundex", code))
			if key == "" {
				continue
			}
			if _, dup := seen[key]; dup {
				continue
			}
			seen[key] = struct{}{}
			out = append(out, key)
		}
	}

	addTokens(e.PreparedFields.NameFields)
	for _, alt := range e.PreparedFields.AltNameFields {
		addTokens(alt)
	}
	for _, hist := range e.PreparedFields.HistoricalNameFields {
		addTokens(hist)
	}

	// Fallback when Normalize() was not called: tokenize the raw name.
	if len(seen) == 0 {
		name := e.PreparedFields.Name
		if name == "" {
			name = e.Name
		}
		if name != "" {
			addTokens(strings.Fields(strings.ToLower(name)))
		}
	}
	return out
}

func appendGovIDKeys(out []string, e search.Entity[search.Value]) []string {
	for _, id := range governmentIDs(e) {
		ident := normalizeIdentifier(id.Identifier)
		if ident == "" {
			continue
		}
		var segs []string
		if c := normalizeCountry(id.Country); c != "" {
			if d := digest("gov-country", c); d != "" {
				segs = append(segs, "C"+d)
			}
		}
		if t := strings.ToLower(strings.TrimSpace(string(id.Type))); t != "" {
			if d := digest("gov-type", t); d != "" {
				segs = append(segs, "T"+d)
			}
		}
		if d := digest("gov-id", ident); d != "" {
			segs = append(segs, "X"+d)
		}
		if key := joinSegments(KindGovID, segs); key != "" {
			out = append(out, key)
		}
	}
	return out
}

func governmentIDs(e search.Entity[search.Value]) []search.GovernmentID {
	var ids []search.GovernmentID
	if e.Person != nil {
		ids = append(ids, e.Person.GovernmentIDs...)
	}
	if e.Business != nil {
		ids = append(ids, e.Business.GovernmentIDs...)
	}
	if e.Organization != nil {
		ids = append(ids, e.Organization.GovernmentIDs...)
	}
	return ids
}

func appendAddressKeys(out []string, e search.Entity[search.Value]) []string {
	if len(e.PreparedFields.Addresses) > 0 {
		for _, addr := range e.PreparedFields.Addresses {
			if key := addressKey(addr.Country, addr.State, addr.PostalCode, addr.City, addr.Line1, addr.Line2); key != "" {
				out = append(out, key)
			}
		}
		return out
	}
	for _, addr := range e.Addresses {
		if key := addressKey(addr.Country, addr.State, addr.PostalCode, addr.City, addr.Line1, addr.Line2); key != "" {
			out = append(out, key)
		}
	}
	return out
}

func addressKey(country, state, postal, city, line1, line2 string) string {
	var segs []string
	if d := digest("addr-country", normalizeCountry(country)); d != "" {
		segs = append(segs, "C"+d)
	}
	if d := digest("addr-state", strings.ToLower(strings.TrimSpace(state))); d != "" {
		segs = append(segs, "S"+d)
	}
	if d := digest("addr-postal", normalizePostal(postal)); d != "" {
		segs = append(segs, "P"+d)
	}
	if d := digest("addr-city", strings.ToLower(strings.TrimSpace(city))); d != "" {
		segs = append(segs, "Y"+d)
	}
	if d := digest("addr-line1", strings.ToLower(strings.TrimSpace(line1))); d != "" {
		segs = append(segs, "L"+d)
	}
	if d := digest("addr-line2", strings.ToLower(strings.TrimSpace(line2))); d != "" {
		segs = append(segs, "E"+d)
	}
	return joinSegments(KindAddr, segs)
}

func appendContactKeys(out []string, e search.Entity[search.Value]) []string {
	for _, email := range e.Contact.EmailAddresses {
		email = strings.ToLower(strings.TrimSpace(email))
		if d := digest("email", email); d != "" {
			out = append(out, KindContact+":E"+d)
		}
	}

	phones := e.PreparedFields.Contact.PhoneNumbers
	if len(phones) == 0 {
		phones = e.Contact.PhoneNumbers
	}
	for _, phone := range phones {
		phone = strings.TrimSpace(phone)
		if d := digest("phone", phone); d != "" {
			out = append(out, KindContact+":P"+d)
		}
	}
	return out
}

func appendVesselAircraftKeys(out []string, e search.Entity[search.Value]) []string {
	if e.Vessel != nil {
		if d := digest("imo", normalizeIdentifier(e.Vessel.IMONumber)); d != "" {
			out = append(out, joinKind(KindIMO, d))
		}
		if d := digest("mmsi", normalizeIdentifier(e.Vessel.MMSI)); d != "" {
			out = append(out, joinKind(KindMMSI, d))
		}
	}
	if e.Aircraft != nil {
		if d := digest("aircraft-serial", normalizeIdentifier(e.Aircraft.SerialNumber)); d != "" {
			out = append(out, joinKind(KindAir, d))
		}
	}
	return out
}

func normalizeCountry(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	return strings.ToLower(norm.Country(s))
}

func normalizePostal(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.NewReplacer(" ", "", "-", "").Replace(s)
	return s
}

func normalizeIdentifier(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.NewReplacer(" ", "", "-", "").Replace(s)
	return s
}
