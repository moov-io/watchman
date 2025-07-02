package linksim

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
)

// SortKeys returns a slice of sortable keys from the given buckets.
//
// Sortable keys can be put into a database where the closest neighbors would be records most similar to the target.
func SortKeys(buckets map[string][]int) []string {
	var out []string

	// Type key
	out = makeKey(buckets, "EntityType", typeKey, out)

	// Source key
	source, foundSource := buckets["Source"]
	sourceID, foundSourceID := buckets["SourceID"]
	if foundSource && foundSourceID {
		out = append(out, makeSourceKey(source, sourceID))
	}

	// Name keys
	out = makeKey(buckets, "Name", nameKey, out)
	out = makeKey(buckets, "AltNames", nameKey, out)

	// GovernmentID keys
	out = makeGovernmentIDKeys(buckets, out)

	// Address keys
	out = makeAddressKeys(buckets, out)

	// Contact keys // TODO(adam):
	// "Contact.EmailAddresses"
	// "Contact.PhoneNumbers"

	return out
}

func makeKey(buckets map[string][]int, key string, f func([]int) string, out []string) []string {
	if v, found := buckets[key]; found {
		return append(out, f(v))
	}
	return out
}

func typeKey(vs []int) string {
	slices.Sort(vs)

	var buf bytes.Buffer
	buf.WriteString("TYPE:")
	for idx := range vs {
		if idx > 1 {
			buf.WriteString("|")
		}
		buf.WriteString(fmt.Sprintf("%4.4d", vs[idx]))
	}
	return buf.String()
}

func makeSourceKey(source []int, sourceID []int) string {
	slices.Sort(source)
	slices.Sort(sourceID)

	var buf bytes.Buffer
	buf.WriteString("SOURCE:")
	for idx := range source {
		if idx > 1 {
			buf.WriteString("|")
		}
		buf.WriteString(fmt.Sprintf("L:%4.4d", source[idx]))
	}

	for idx := range sourceID {
		if idx > 1 {
			buf.WriteString("|")
		}
		buf.WriteString(fmt.Sprintf("X%4.4d", sourceID[idx]))
	}

	return buf.String()
}

func nameKey(vs []int) string {
	slices.Sort(vs)

	var buf bytes.Buffer
	buf.WriteString("NAME:")
	for idx := range vs {
		if idx > 1 {
			buf.WriteString("|")
		}
		buf.WriteString(fmt.Sprintf("%4.4d", vs[idx]))
	}
	return buf.String()
}

func makeGovernmentIDKeys(buckets map[string][]int, out []string) []string {
	// Look for Person and Business GovernmentIDs
	out = readGovernmentIDs(buckets, "Person.GovernmentIDs[%d].%s", out)
	out = readGovernmentIDs(buckets, "Business.GovernmentIDs[%d].%s", out)

	return out
}

func readGovernmentIDs(buckets map[string][]int, pattern string, out []string) []string {
	idx := 0
	for {
		countryKey := fmt.Sprintf(pattern, idx, "Country")
		if country, found := buckets[countryKey]; found {
			// Look for Identifier and Type now
			identifier, found := buckets[fmt.Sprintf(pattern, idx, "Identifier")]
			if !found {
				continue
			}
			tpe, found := buckets[fmt.Sprintf(pattern, idx, "Type")]
			if !found {
				continue
			}

			out = append(out, makeGovernmentIDKey(country, tpe, identifier))

			idx++
		} else {
			break
		}
	}
	return out
}

func makeGovernmentIDKey(country, tpe, identifier []int) string {
	slices.Sort(country)
	slices.Sort(tpe)
	slices.Sort(identifier)

	var buf bytes.Buffer
	buf.WriteString("GOVID:")

	if len(country) == 0 || len(tpe) == 0 || len(identifier) == 0 {
		return fmt.Sprintf("INVALID: %v / %v / %v", country, tpe, identifier)
	}

	// only take the first
	buf.WriteString(fmt.Sprintf("C%4.4d", country[0]))
	buf.WriteString("|")
	buf.WriteString(fmt.Sprintf("T%4.4d", tpe[0]))
	buf.WriteString("|")
	buf.WriteString(fmt.Sprintf("X%4.4d", identifier[0]))

	return buf.String()
}

func makeAddressKeys(buckets map[string][]int, out []string) []string {
	for i := 0; i < 20; i++ {
		var buf bytes.Buffer
		buf.WriteString("ADDR:")

		var fieldsCollected int
		pattern := fmt.Sprintf("Addresses[%d].%%s", i)

		// Look for the fields from broadest to most precise (then extra)
		// starting with Country.
		if v, found := buckets[fmt.Sprintf(pattern, "Country")]; found && len(v) > 0 {
			fieldsCollected++
			buf.WriteString(fmt.Sprintf("C%4.4d|", v[0]))
		}
		if v, found := buckets[fmt.Sprintf(pattern, "State")]; found && len(v) > 0 {
			fieldsCollected++
			buf.WriteString(fmt.Sprintf("S%4.4d|", v[0]))
		}
		if v, found := buckets[fmt.Sprintf(pattern, "PostalCode")]; found && len(v) > 0 {
			fieldsCollected++
			buf.WriteString(fmt.Sprintf("P%4.4d|", v[0]))
		}
		if v, found := buckets[fmt.Sprintf(pattern, "CityFields")]; found && len(v) > 0 {
			fieldsCollected++
			buf.WriteString("Y")

			for idx := range v {
				if idx > 0 {
					buf.WriteString(",")
				}
				buf.WriteString(fmt.Sprintf("%4.4d", v[idx]))
			}
		}
		buf.WriteString("|")
		if v, found := buckets[fmt.Sprintf(pattern, "Line1Fields")]; found && len(v) > 0 {
			fieldsCollected++
			buf.WriteString("L")

			for idx := range v {
				if idx > 0 {
					buf.WriteString(",")
				}
				buf.WriteString(fmt.Sprintf("%4.4d", v[idx]))
			}
		}
		buf.WriteString("|")
		if v, found := buckets[fmt.Sprintf(pattern, "Line2Fields")]; found && len(v) > 0 {
			fieldsCollected++
			buf.WriteString("E")

			for idx := range v {
				if idx > 0 {
					buf.WriteString(",")
				}
				buf.WriteString(fmt.Sprintf("%4.4d", v[idx]))
			}
		}

		if fieldsCollected > 0 {
			addr := strings.TrimSuffix(buf.String(), "|")
			out = append(out, addr)
		} else {
			break // nothing found for this index, so no further indexes will have anything
		}
	}

	return out
}
