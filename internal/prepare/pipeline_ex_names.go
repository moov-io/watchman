// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Longer markers first so "formerly known as" wins over "formerly".
var formerMarkers = []string{
	"formerly known as",
	"f.k.a.",
	"f/k/a",
	"f.k.a",
	"formerly",
}

var (
	exChainSplit  = regexp.MustCompile(`(?i)[,;]\s*ex[\-\s\x{2013}\x{2014}]+`)
	formerDateRe  = regexp.MustCompile(`(?i)(?:,\s*|\s+until\s+)\d{4}\s*$`)
	formerLeadSep = regexp.MustCompile(`^[\s\-\x{2013}\x{2014}:,]+`)
)

// SplitExNames pulls former names out of a display name and returns the
// remaining current name. It understands:
//
//   - "(ex-Cape Diamond)" / "(ex Cape Diamond)"
//   - chained "(ex A, ex B)"
//   - a trailing year on the former name (", 2020" / "until 2019")
//   - "(f.k.a. …)" / "(f/k/a …)" / "(formerly …)" / "(formerly known as …)"
//   - the same markers inline, and an "ex" token with a separator
//     ("Kalliopi Dawn ex Orion Tanager", "EX PALOMERA")
//
// "Ocean Pioneer (example)" is not a former name: "ex" needs a separator.
func SplitExNames(s string) (cleaned string, former []string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", nil
	}

	var b strings.Builder
	i := 0
	for i < len(s) {
		start := strings.IndexByte(s[i:], '(')
		if start < 0 {
			b.WriteString(s[i:])
			break
		}
		start += i
		b.WriteString(s[i:start])
		endRel := strings.IndexByte(s[start:], ')')
		if endRel < 0 {
			b.WriteString(s[start:])
			break
		}
		end := start + endRel
		inner := strings.TrimSpace(s[start+1 : end])
		if names, ok := cutFormerInner(inner); ok {
			former = append(former, names...)
			i = end + 1
			continue
		}
		b.WriteString(s[start : end+1])
		i = end + 1
	}

	cleaned = strings.Join(strings.Fields(b.String()), " ")
	cleaned, extra := splitInlineFormer(cleaned)
	former = append(former, extra...)
	return cleaned, former
}

func cutFormerInner(inner string) ([]string, bool) {
	inner = strings.TrimSpace(inner)
	if inner == "" {
		return nil, false
	}
	if rest, ok := splitExPrefix(inner); ok {
		if rest == "" {
			return nil, false
		}
		return splitFormerList(rest), true
	}
	lower := strings.ToLower(inner)
	for _, m := range formerMarkers {
		if strings.HasPrefix(lower, m) {
			rest := strings.TrimSpace(inner[len(m):])
			rest = formerLeadSep.ReplaceAllString(rest, "")
			rest = strings.TrimSpace(rest)
			if rest == "" {
				return nil, false
			}
			return splitFormerList(rest), true
		}
	}
	return nil, false
}

func splitInlineFormer(s string) (string, []string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", nil
	}
	start, end, ok := indexFormerMarker(s)
	if !ok {
		return s, nil
	}
	current := strings.TrimSpace(s[:start])
	current = strings.TrimRight(current, ",;")
	current = strings.TrimSpace(current)
	rest := formerLeadSep.ReplaceAllString(s[end:], "")
	rest = strings.TrimSpace(rest)
	if rest == "" {
		return s, nil
	}
	return current, splitFormerList(rest)
}

func indexFormerMarker(s string) (start, end int, ok bool) {
	lower := strings.ToLower(s)
	bestStart := -1
	bestEnd := 0
	for _, m := range formerMarkers {
		i := strings.Index(lower, m)
		if i < 0 {
			continue
		}
		j := i + len(m)
		if bestStart < 0 || i < bestStart || (i == bestStart && j > bestEnd) {
			bestStart = i
			bestEnd = j
		}
	}
	if i, j, found := indexExToken(s); found {
		if bestStart < 0 || i < bestStart {
			return i, j, true
		}
	}
	if bestStart < 0 {
		return 0, 0, false
	}
	return bestStart, bestEnd, true
}

func indexExToken(s string) (start, end int, ok bool) {
	lower := strings.ToLower(s)
	for i := 0; i+2 <= len(lower); {
		j := strings.Index(lower[i:], "ex")
		if j < 0 {
			return 0, 0, false
		}
		j += i
		if j > 0 {
			r, _ := utf8.DecodeLastRuneInString(s[:j])
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				i = j + 2
				continue
			}
		}
		rest := s[j+2:]
		if rest == "" {
			return 0, 0, false
		}
		if isExSeparatorPrefix(rest) {
			return j, j + 2, true
		}
		i = j + 2
	}
	return 0, 0, false
}

func splitFormerList(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	raw := exChainSplit.Split(s, -1)
	out := make([]string, 0, len(raw))
	for _, p := range raw {
		p = stripFormerDate(p)
		p = strings.TrimSpace(p)
		p = strings.Trim(p, ",;")
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func stripFormerDate(s string) string {
	return strings.TrimSpace(formerDateRe.ReplaceAllString(strings.TrimSpace(s), ""))
}

func splitExPrefix(inner string) (string, bool) {
	inner = strings.TrimSpace(inner)
	if len(inner) < 3 {
		return "", false
	}
	if !strings.EqualFold(inner[:2], "ex") {
		return "", false
	}
	switch inner[2] {
	case '-', ' ', '\t':
		return strings.TrimSpace(inner[3:]), true
	}
	if strings.HasPrefix(inner[2:], "–") {
		return strings.TrimSpace(inner[2+len("–"):]), true
	}
	if strings.HasPrefix(inner[2:], "—") {
		return strings.TrimSpace(inner[2+len("—"):]), true
	}
	return "", false
}

func isExSeparatorPrefix(s string) bool {
	if s == "" {
		return false
	}
	switch s[0] {
	case '-', ' ', '\t':
		return true
	}
	return strings.HasPrefix(s, "–") || strings.HasPrefix(s, "—")
}

// ExNames returns former names encoded in s.
func ExNames(s string) []string {
	_, former := SplitExNames(s)
	return former
}
