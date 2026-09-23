// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package index

// qwertyNeighbors maps a lowercase ASCII rune to adjacent US-QWERTY keys.
// Used for typed identifier queries (IMO, email, …), not name scoring.
var qwertyNeighbors = map[rune]string{
	'1': "2q", '2': "13qw", '3': "24we", '4': "35er", '5': "46rt",
	'6': "57ty", '7': "68yu", '8': "79ui", '9': "80io", '0': "9op",
	'q': "12wa", 'w': "23qesa", 'e': "34wrsd", 'r': "45etdf", 't': "56ryfg",
	'y': "67tugh", 'u': "78yihj", 'i': "89uojk", 'o': "90ipkl", 'p': "0ol",
	'a': "qwsz", 's': "weadzx", 'd': "erfcxs", 'f': "rtgvcd", 'g': "tyhbvf",
	'h': "yujnbg", 'j': "uikmnh", 'k': "iolmj", 'l': "opk",
	'z': "asx", 'x': "zsdc", 'c': "xdfv", 'v': "cfgb", 'b': "vghn",
	'n': "bhjm", 'm': "njk",
}

// qwertyVariants returns q plus single adjacent-key substitutions and
// adjacent transpositions. q itself is first.
func qwertyVariants(q string) []string {
	if q == "" {
		return nil
	}
	seen := map[string]struct{}{q: {}}
	out := []string{q}
	runes := []rune(q)
	add := func(s string) {
		if _, ok := seen[s]; ok {
			return
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	for i, r := range runes {
		for _, n := range qwertyNeighbors[r] {
			runes[i] = n
			add(string(runes))
		}
		runes[i] = r
	}
	for i := 0; i < len(runes)-1; i++ {
		runes[i], runes[i+1] = runes[i+1], runes[i]
		add(string(runes))
		runes[i], runes[i+1] = runes[i+1], runes[i]
	}
	return out
}
