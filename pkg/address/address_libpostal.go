//go:build libpostal

package address

import (
	"strings"

	"github.com/moov-io/watchman/pkg/search"

	postal "github.com/openvenues/gopostal/parser"
)

func ParseAddress(input string) search.Address {
	parts := postal.ParseAddress(input)

	return organizeLibpostalComponents(parts)
}

func organizeLibpostalComponents(parsed []postal.ParsedComponent) search.Address {
	// Convert the slice of ParsedComponents into a map for easy access
	components := make(map[string]string)
	for _, c := range parsed {
		components[c.Label] = c.Value
	}

	var addr search.Address

	var houseParts []string
	var line2Parts []string

	// If building name (house) is present, include it in line1
	if val, ok := components["house"]; ok && val != "" {
		houseParts = append(houseParts, val)
	}

	// Add house_number + road to line1
	if val, ok := components["house_number"]; ok && val != "" {
		houseParts = append(houseParts, val)
	}
	if val, ok := components["road"]; ok && val != "" {
		houseParts = append(houseParts, val)
	}

	addr.Line1 = joinNonEmpty(houseParts, " ")

	// Append unit, level, staircase, entrance to line2 if present
	secondaryLabels := []string{"unit", "level", "staircase", "entrance"}
	for _, label := range secondaryLabels {
		if val, ok := components[label]; ok && val != "" {
			line2Parts = append(line2Parts, val)
		}
	}
	addr.Line2 = joinNonEmpty(line2Parts, ", ")

	// City: prefer city, if not present fallback to city_district or suburb
	if val, ok := components["city"]; ok && val != "" {
		addr.City = val
	} else if val, ok := components["city_district"]; ok && val != "" {
		addr.City = val
	} else if val, ok := components["suburb"]; ok && val != "" {
		addr.City = val
	}

	// PostalCode
	if val, ok := components["postcode"]; ok && val != "" {
		addr.PostalCode = val
	}

	// State: prefer state if present, else state_district
	if val, ok := components["state"]; ok && val != "" {
		addr.State = val
	} else if val, ok := components["state_district"]; ok && val != "" {
		addr.State = val
	}

	// Country
	if val, ok := components["country"]; ok && val != "" {
		addr.Country = val
	}

	// Latitude/Longitude not provided by libpostal parsing
	// If you have them from another source, set them here.

	return addr
}

func joinNonEmpty(parts []string, sep string) string {
	var nonEmpty []string
	for _, p := range parts {
		if p != "" {
			nonEmpty = append(nonEmpty, p)
		}
	}
	return strings.Join(nonEmpty, sep)
}
