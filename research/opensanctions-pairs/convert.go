// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"strings"
	"time"
	"unicode"

	"github.com/moov-io/watchman/internal/norm"
	"github.com/moov-io/watchman/pkg/search"
)

const defaultMaxAltNames = 20

var orgLike = map[string]struct{}{
	"company":      {},
	"organization": {},
	"legalentity":  {},
	"publicbody":   {},
	"corporation":  {},
}

var autoMergeSchemas = map[string]struct{}{
	"occupancy":      {},
	"position":       {},
	"succession":     {},
	"family":         {},
	"directorship":   {},
	"ownership":      {},
	"representation": {},
	"membership":     {},
	"associate":      {},
	"employment":     {},
	"unknownlink":    {},
	"address":        {},
	"identification": {},
	"security":       {},
	"cryptowallet":   {},
}

var genericCaptions = map[string]struct{}{
	"occupancy":    {},
	"position":     {},
	"address":      {},
	"succession":   {},
	"company":      {},
	"person":       {},
	"organization": {},
	"legalentity":  {},
	"vessel":       {},
	"airplane":     {},
	"aircraft":     {},
}

type convertOpts struct {
	maxAltNames int
	nameOnly    bool
}

func isAutoMerge(schema string) bool {
	_, ok := autoMergeSchemas[strings.ToLower(schema)]
	return ok
}

func isOrgLike(schema string) bool {
	_, ok := orgLike[strings.ToLower(schema)]
	return ok
}

func entityTypeFor(schema string) search.EntityType {
	switch strings.ToLower(schema) {
	case "person":
		return search.EntityPerson
	case "vessel":
		return search.EntityVessel
	case "airplane", "aircraft":
		return search.EntityAircraft
	case "company", "organization", "legalentity", "publicbody", "corporation":
		return search.EntityBusiness
	default:
		return search.EntityUnknown
	}
}

// queryTypes is the Watchman type fan-out a client would use. FollowTheMoney
// LegalEntity (and other unknown schemas) are not a single Watchman partition;
// production search requires type= and recommends one call per type when the
// query type is unknown. Vessel/aircraft are included only when those IDs exist.
func queryTypes(e FTMEntity) []search.EntityType {
	schema := strings.ToLower(e.Schema)
	if schema == "legalentity" || entityTypeFor(e.Schema) == search.EntityUnknown {
		out := []search.EntityType{search.EntityPerson, search.EntityBusiness}
		if e.firstProp("imoNumber", "mmsi", "callSign") != "" {
			out = append(out, search.EntityVessel)
		}
		if e.firstProp("icaoCode", "serialNumber") != "" {
			out = append(out, search.EntityAircraft)
		}
		return out
	}
	t := entityTypeFor(e.Schema)
	if t == search.EntityUnknown {
		return []search.EntityType{search.EntityPerson, search.EntityBusiness}
	}
	return []search.EntityType{t}
}

func commonTypes(left, right FTMEntity) []search.EntityType {
	ls, rs := queryTypes(left), queryTypes(right)
	seen := make(map[search.EntityType]struct{}, len(ls))
	for _, t := range ls {
		seen[t] = struct{}{}
	}
	var out []search.EntityType
	have := make(map[search.EntityType]struct{})
	for _, t := range rs {
		if _, ok := seen[t]; !ok {
			continue
		}
		if _, dup := have[t]; dup {
			continue
		}
		have[t] = struct{}{}
		out = append(out, t)
	}
	return out
}

func convertPair(p Pair, opts convertOpts) (search.Entity[search.Value], search.Entity[search.Value]) {
	left := convertEntity(p.Left, opts)
	right := convertEntity(p.Right, opts)
	// Company vs Organization would otherwise score 0 from the type filter.
	if left.Type != right.Type && isOrgLike(p.Left.Schema) && isOrgLike(p.Right.Schema) {
		left.Type = search.EntityBusiness
		right.Type = search.EntityBusiness
	}
	return left.Normalize(), right.Normalize()
}

func convertEntity(e FTMEntity, opts convertOpts) search.Entity[search.Value] {
	return convertEntityAs(e, opts, entityTypeFor(e.Schema))
}

func convertEntityAs(e FTMEntity, opts convertOpts, typ search.EntityType) search.Entity[search.Value] {
	if opts.maxAltNames <= 0 {
		opts.maxAltNames = defaultMaxAltNames
	}

	out := search.Entity[search.Value]{
		Name: pickPrimaryName(e),
		Type: typ,
	}

	names := uniqueStrings(append([]string{out.Name}, e.prop("name", "alias", "weakAlias", "previousName")...))
	alt := make([]string, 0, len(names))
	for _, n := range names {
		if n == "" || strings.EqualFold(n, out.Name) {
			continue
		}
		alt = append(alt, n)
		if len(alt) >= opts.maxAltNames {
			break
		}
	}

	country := firstCountry(e)
	ids := governmentIDs(e, country)

	if opts.nameOnly {
		attachNames(&out, alt)
		return out
	}

	out.Addresses = addresses(e, country)
	out.Contact = contactInfo(e)
	out.SanctionsInfo = sanctionsInfo(e)

	former := e.prop("previousName")
	for _, name := range former {
		out.HistoricalInfo = append(out.HistoricalInfo, search.HistoricalInfo{
			Type:  "Former Name",
			Value: name,
		})
	}

	switch out.Type {
	case search.EntityPerson:
		out.Person = &search.Person{
			Name:          out.Name,
			AltNames:      alt,
			Gender:        genderOf(e.firstProp("gender")),
			BirthDate:     pickDate(e.prop("birthDate")),
			PlaceOfBirth:  e.firstProp("birthPlace"),
			DeathDate:     pickDate(e.prop("deathDate")),
			Titles:        limitStrings(e.prop("position", "title"), 8),
			GovernmentIDs: ids,
		}
	case search.EntityBusiness:
		out.Business = &search.Business{
			Name:          out.Name,
			AltNames:      alt,
			Created:       pickDate(e.prop("incorporationDate", "createdAt")),
			Dissolved:     pickDate(e.prop("dissolutionDate")),
			GovernmentIDs: ids,
		}
	case search.EntityVessel:
		out.Vessel = &search.Vessel{
			Name:      out.Name,
			AltNames:  alt,
			IMONumber: e.firstProp("imoNumber"),
			MMSI:      e.firstProp("mmsi"),
			CallSign:  e.firstProp("callSign"),
			Flag:      country,
			Type:      search.VesselTypeUnknown,
		}
	case search.EntityAircraft:
		out.Aircraft = &search.Aircraft{
			Name:         out.Name,
			AltNames:     alt,
			Flag:         country,
			ICAOCode:     e.firstProp("icaoCode"),
			SerialNumber: e.firstProp("serialNumber", "registrationNumber"),
			Type:         search.AircraftTypeUnknown,
		}
	default:
		attachNames(&out, alt)
	}

	return out
}

func attachNames(out *search.Entity[search.Value], alt []string) {
	switch out.Type {
	case search.EntityPerson:
		out.Person = &search.Person{Name: out.Name, AltNames: alt}
	case search.EntityBusiness, search.EntityOrganization:
		if out.Type == search.EntityOrganization {
			out.Organization = &search.Organization{Name: out.Name, AltNames: alt}
		} else {
			out.Business = &search.Business{Name: out.Name, AltNames: alt}
		}
	case search.EntityVessel:
		out.Vessel = &search.Vessel{Name: out.Name, AltNames: alt}
	case search.EntityAircraft:
		out.Aircraft = &search.Aircraft{Name: out.Name, AltNames: alt}
	}
}

func pickPrimaryName(e FTMEntity) string {
	caption := strings.TrimSpace(e.Caption)
	names := e.prop("name")
	if caption != "" && !isGenericCaption(caption, e.Schema) {
		return caption
	}
	for _, n := range names {
		if !hasNonLatin(n) {
			return n
		}
	}
	if len(names) > 0 {
		return names[0]
	}
	return caption
}

func isGenericCaption(caption, schema string) bool {
	c := strings.ToLower(strings.TrimSpace(caption))
	if c == "" {
		return true
	}
	if strings.EqualFold(c, schema) {
		return true
	}
	_, ok := genericCaptions[c]
	return ok
}

func governmentIDs(e FTMEntity, country string) []search.GovernmentID {
	var ids []search.GovernmentID
	add := func(values []string, idType search.GovernmentIDType, name string) {
		for _, v := range values {
			ids = append(ids, search.GovernmentID{
				Name:       name,
				Type:       idType,
				Country:    country,
				Identifier: v,
			})
		}
	}

	add(e.prop("passportNumber"), search.GovernmentIDPassport, "passport")
	add(e.prop("idNumber"), search.GovernmentIDNational, "national-id")
	add(e.prop("innCode", "taxNumber", "vatCode", "kppCode"), search.GovernmentIDTax, "tax-id")
	add(e.prop("registrationNumber", "ogrnCode", "okpoCode", "leiCode", "dunsCode"), search.GovernmentIDBusinessRegisration, "registration")
	return ids
}

func addresses(e FTMEntity, country string) []search.Address {
	raw := e.prop("address")
	if len(raw) == 0 {
		if country == "" {
			return nil
		}
		return []search.Address{{Country: country}}
	}
	out := make([]search.Address, 0, len(raw))
	for _, line := range raw {
		out = append(out, search.Address{
			Line1:   line,
			Country: country,
		})
	}
	return out
}

func contactInfo(e FTMEntity) search.ContactInfo {
	return search.ContactInfo{
		EmailAddresses: e.prop("email"),
		PhoneNumbers:   e.prop("phone"),
		Websites:       e.prop("website"),
	}
}

func sanctionsInfo(e FTMEntity) *search.SanctionsInfo {
	programs := e.prop("programId")
	if len(programs) == 0 {
		return nil
	}
	return &search.SanctionsInfo{Programs: programs}
}

func genderOf(raw string) search.Gender {
	switch strings.ToLower(strings.TrimSpace(raw)) {
	case "male", "m":
		return search.GenderMale
	case "female", "f":
		return search.GenderFemale
	default:
		return search.GenderUnknown
	}
}

func firstCountry(e FTMEntity) string {
	for _, v := range e.prop("nationality", "country", "jurisdiction", "citizenship", "birthCountry") {
		if name := norm.Country(v); name != "" {
			return name
		}
	}
	return ""
}

func pickDate(values []string) *time.Time {
	var best string
	for _, v := range values {
		v = strings.TrimSpace(v)
		if len(v) > len(best) {
			best = v
		}
	}
	if best == "" {
		return nil
	}
	layouts := []string{
		time.RFC3339,
		"2006-01-02",
		"2006-01",
		"2006",
		"02 Jan 2006",
		"January 2006",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, best); err == nil {
			return &t
		}
	}
	if len(best) >= 4 && isDigits(best[:4]) {
		if t, err := time.Parse("2006", best[:4]); err == nil {
			return &t
		}
	}
	return nil
}

func isDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return len(s) > 0
}

func uniqueStrings(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	var out []string
	for _, v := range in {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		key := strings.ToLower(v)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, v)
	}
	return out
}

func limitStrings(in []string, n int) []string {
	if len(in) <= n {
		return in
	}
	return in[:n]
}

func hasNonLatin(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r) {
			return true
		}
	}
	return false
}

func pairCrossScript(p Pair) bool {
	leftLatin := !hasNonLatin(p.Left.Caption) && !hasNonLatin(strings.Join(p.Left.prop("name"), " "))
	rightLatin := !hasNonLatin(p.Right.Caption) && !hasNonLatin(strings.Join(p.Right.prop("name"), " "))
	return leftLatin != rightLatin
}
