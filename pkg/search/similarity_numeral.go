package search

import (
	"strconv"
	"strings"
	"unicode"
)

var nameNumeralConflictPenaltyMultiplier = readFloat("NAME_NUMERAL_CONFLICT_PENALTY_MULTIPLIER", 0.70)

// Tokens that are a hull number, SPV ordinal, or roman/spelled numeral.
var spelledNameNumerals = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
	"six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10,
	"eleven": 11, "twelve": 12, "thirteen": 13, "fourteen": 14, "fifteen": 15,
	"sixteen": 16, "seventeen": 17, "eighteen": 18, "nineteen": 19, "twenty": 20,
	"i": 1, "ii": 2, "iii": 3, "iv": 4, "v": 5, "vi": 6, "vii": 7, "viii": 8, "ix": 9,
	"x": 10, "xi": 11, "xii": 12,
}

func hasNameNumeralConflict[Q any, I any](query Entity[Q], index Entity[I]) bool {
	// Use the prepared name string, not NameFields: English number words
	// ("three", "five") are stopwords and would otherwise disappear.
	qNums := nameNumerals(strings.Fields(query.PreparedFields.Name))
	iNums := nameNumerals(strings.Fields(index.PreparedFields.Name))
	if len(qNums) == 0 || len(iNums) == 0 {
		return false
	}
	for n := range qNums {
		if iNums[n] {
			return false
		}
	}
	return true
}

func nameNumerals(fields []string) map[int]bool {
	out := make(map[int]bool)
	for _, tok := range fields {
		if n, ok := tokenNumeral(tok); ok {
			out[n] = true
		}
	}
	return out
}

func tokenNumeral(tok string) (int, bool) {
	tok = strings.ToLower(strings.TrimSpace(tok))
	if tok == "" || tok == "no" {
		return 0, false
	}
	if n, ok := spelledNameNumerals[tok]; ok {
		return n, true
	}
	if !allDigits(tok) {
		return 0, false
	}
	n, err := strconv.Atoi(tok)
	if err != nil {
		return 0, false
	}
	return n, true
}

func allDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
