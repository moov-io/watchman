// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/abadojack/whatlanggo"
	"github.com/bbalet/stopwords"
	"github.com/pariz/gountries"
)

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

type stopwordsStep struct{}

func (s *stopwordsStep) apply(in *Name) error {
	if in == nil {
		return nil
	}

	switch {
	case in.sdn != nil && !strings.EqualFold(in.sdn.SDNType, "individual"):
		in.Processed = removeStopwords(in.Processed, detectLanguage(in.Processed, in.addrs))
	case in.ssi != nil && !strings.EqualFold(in.ssi.Type, "individual"):
		in.Processed = removeStopwords(in.Processed, detectLanguage(in.Processed, nil))
	case in.alt != nil:
		in.Processed = removeStopwords(in.Processed, detectLanguage(in.Processed, nil))
	}
	return nil
}

var (
	numberRegex = regexp.MustCompile(`([\d\.\,\-]{1,}[\d]{1,})`)
)

func removeStopwords(in string, lang whatlanggo.Lang) string {
	if keepStopwords {
		return in
	}

	var out []string
	words := strings.Fields(strings.ToLower(in))
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
func detectLanguage(in string, addrs []*ofac.Address) whatlanggo.Lang {
	info := whatlanggo.Detect(in)
	if info.IsReliable() {
		// Return the detected language if whatlanggo is confident enough
		return info.Lang
	}

	if len(addrs) == 0 {
		// If no addresses are associated to this text blob then fallback to English
		return whatlanggo.Eng
	}

	// Return the countries primary language associated to the primary address for this SDN.
	//
	// TODO(adam): Should we do this only if there's one address? If there are multiple should we
	// fallback to English or a mixed set?
	country, err := gountries.New().FindCountryByName(addrs[0].Country)
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
