// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

//go:build !js

package prepare

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/abadojack/whatlanggo"
	"github.com/bbalet/stopwords"
	"github.com/pariz/gountries"
	"golang.org/x/text/unicode/norm"
)

//go:generate go run gen_union.go

const (
	minConfidence = 0.50
)

var (
	keepStopwords = func(raw string) bool {
		if raw == "" {
			raw = "false"
		}
		keep, _ := strconv.ParseBool(raw)
		return keep
	}(os.Getenv("KEEP_STOPWORDS"))
)

// switch {
// case in.sdn != nil && !strings.EqualFold(in.sdn.SDNType, "individual"):
// 	in.Processed = removeStopwords(in.Processed, detectLanguage(in.Processed, in.addrs))
// case in.ssi != nil && !strings.EqualFold(in.ssi.Type, "individual"):
// 	in.Processed = removeStopwords(in.Processed, detectLanguage(in.Processed, nil))
// case in.alt != nil:
// 	in.Processed = removeStopwords(in.Processed, detectLanguage(in.Processed, nil))
// }

var (
	numberRegex = regexp.MustCompile(`([\d\.\,\-]{1,}[\d]{1,})`)
)

// wordSegmenter mirrors the default token segmenter used by
// github.com/bbalet/stopwords, so we can tell whether CleanString would reshape a
// token rather than only drop a stopword.
var wordSegmenter = regexp.MustCompile(`[\pL\p{Mc}\p{Mn}-_']+`)

// trivialNormalization returns the normalized text and true when the detected
// language cannot affect the result, allowing RemoveStopwords to skip the
// (expensive) whatlanggo.Detect call.
//
// stopwords.CleanString(token, lang) leaves a whitespace-delimited token unchanged
// under any language when either:
//   - it matches numberRegex (removeStopwords never cleans numbers), or
//   - it is not a stopword in any language (absent from stopwordUnion) and is
//     already NFC and a single intact segment, so the supported-language branch's
//     NFC and segmentation are no-ops and match the unsupported-language branch,
//     which does no segmentation at all.
//
// When every token is trivial the output is just the lowercased tokens rejoined,
// identical to the detect-then-clean path for any language.
func trivialNormalization(input string) (string, bool) {
	words := strings.Fields(strings.ToLower(input))
	for _, w := range words {
		if numberRegex.MatchString(w) {
			continue
		}
		if _, isStopword := stopwordUnion[w]; isStopword {
			return "", false
		}
		if norm.NFC.String(w) != w {
			return "", false
		}
		if seg := wordSegmenter.FindAllString(w, -1); len(seg) != 1 || seg[0] != w {
			return "", false
		}
	}
	return strings.Join(words, " "), true
}

func RemoveStopwords(input string) string {
	if keepStopwords {
		return input
	}
	if out, ok := trivialNormalization(input); ok {
		return out
	}

	info := whatlanggo.Detect(input)

	return removeStopwords(input, info.Lang)
}

func RemoveStopwordsCountry(input string, countryName string) string {
	lang := detectLanguage(input, countryName)

	return removeStopwords(input, lang)
}

func removeStopwords(input string, lang whatlanggo.Lang) string {
	if keepStopwords {
		return input
	}

	var out []string
	words := strings.Fields(strings.ToLower(input))
	for i := range words {
		cleaned := strings.TrimSpace(words[i])

		// When the word is a number leave it alone
		if !numberRegex.MatchString(cleaned) {
			cleaned = strings.TrimSpace(stopwords.CleanString(cleaned, lang.Iso6391(), false))
		}
		if cleaned != "" {
			out = append(out, cleaned)
		}
	}
	return strings.Join(out, " ")
}

// detectLanguage will return a guess as to the appropriate language a given SDN's name
// is written in. The addresses must be linked to the SDN whose name is detected.
func detectLanguage(in string, countryName string) whatlanggo.Lang {
	info := whatlanggo.Detect(in)
	if info.IsReliable() {
		// Return the detected language if whatlanggo is confident enough
		return info.Lang
	}

	if countryName == "" {
		// If no addresses are associated to this text blob then fallback to English
		return whatlanggo.Eng
	}

	// Return the countries primary language associated to the primary address for this SDN.
	//
	// TODO(adam): Should we do this only if there's one address? If there are multiple should we
	// fallback to English or a mixed set?
	country, err := gountries.New().FindCountryByName(countryName)
	if len(country.Languages) == 0 || err != nil {
		return whatlanggo.Eng
	}

	// If the language is spoken in the country and we're somewhat confident in the original detection
	// then return that language.
	if info.Confidence > minConfidence {
		for key := range country.Languages {
			if strings.EqualFold(key, info.Lang.Iso6393()) {
				return info.Lang
			}
		}
	}
	if len(country.Languages) == 1 {
		for key := range country.Languages {
			return whatlanggo.CodeToLang(key)
		}
	}

	// How should we pick the language for countries with multiple languages? A hardcoded map?
	// What if we found the language whose name is closest to the country's name and returned that?
	//
	// Should this fallback be the mixed set that contains stop words from several popular languages
	// in the various data sets?

	return whatlanggo.Eng
}
