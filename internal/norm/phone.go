//go:build !js

package norm

import (
	"cmp"
	"strings"

	"github.com/dongri/phonenumber"
)

var (
	phoneNumberCleaner = strings.NewReplacer("+", "", " ", "", "-", "", "(", "", ")", "")
)

// PhoneNumber strips all non-numeric characters and normalizes phone numbers
func PhoneNumber(input string) string {
	country := phonenumber.GetISO3166ByNumber(input, false)
	if country.Alpha2 == "" {
		country = phonenumber.GetISO3166ByNumber(input, true)
	}
	if country.Alpha2 == "" {
		out := phonenumber.GetISO3166ByMobileNumber(input)
		if len(out) > 0 {
			country = out[0]
		}
	}
	if country.Alpha2 == "" {
		return phoneNumberCleaner.Replace(input) // can't normalize
	}

	// Parse and reformat
	number := phonenumber.Parse(input, country.Alpha2)

	return phoneNumberCleaner.Replace(cmp.Or(number, input))
}
