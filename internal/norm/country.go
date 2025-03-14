//go:build !js

package norm

import (
	"cmp"
	"strings"

	"github.com/moov-io/iso3166"
)

var (
	countryCodeOverrides = map[string]string{
		"CZ": "Czech Republic",
		"GB": "United Kingdom",
		"IR": "Iran",
		"KP": "North Korea",
		"KR": "South Korea",
		"MD": "Moldova",
		"MF": "Saint Martin",
		"RU": "Russia",
		"SX": "Saint Martin",
		"SY": "Syria",
		"TR": "Turkey",
		"TW": "Taiwan",
		"UK": "United Kingdom",
		"US": "United States",
		"VE": "Venezuela",
		"VN": "Vietnam",
		"VG": "Virgin Islands",
		"VI": "Virgin Islands",
	}
)

func Country(input string) string {
	// try input as ISO 3166 code
	if name := iso3166.GetName(input); name != "" {
		over := countryCodeOverrides[strings.ToUpper(input)]
		return cmp.Or(over, name)
	}

	// Try input as Name, find Code and return official Name
	code := iso3166.LookupCode(input)
	if code != "" {
		over := countryCodeOverrides[code]
		return cmp.Or(over, iso3166.GetName(code))
	}

	// return whatever we have
	over := countryCodeOverrides[strings.ToUpper(input)]
	return cmp.Or(over, input)
}
