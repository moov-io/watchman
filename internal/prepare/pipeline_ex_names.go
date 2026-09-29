// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import "strings"

// SplitExNames pulls former names out of "(ex-…)" / "(ex …)" parentheticals
// and returns the name with those parentheticals removed.
// "Ocean Pioneer (ex-Cape Diamond)" → "Ocean Pioneer", ["Cape Diamond"].
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
		if rest, ok := cutExPrefix(inner); ok {
			if rest != "" {
				former = append(former, rest)
			}
			i = end + 1
			continue
		}
		b.WriteString(s[start : end+1])
		i = end + 1
	}

	cleaned = strings.Join(strings.Fields(b.String()), " ")
	return cleaned, former
}

func cutExPrefix(inner string) (string, bool) {
	if len(inner) < 2 {
		return "", false
	}
	if !strings.EqualFold(inner[:2], "ex") {
		return "", false
	}
	rest := strings.TrimLeft(inner[2:], "-–— \t")
	return strings.TrimSpace(rest), true
}

// ExNames returns former names encoded as "(ex-…)" in s.
func ExNames(s string) []string {
	_, former := SplitExNames(s)
	return former
}
