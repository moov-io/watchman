// Copyright Bloomfielddev Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_un

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

var (
	websiteRE    = regexp.MustCompile(`https?://[^\s;]+|www\.[^\s;]+`)
	emailLabelRE = regexp.MustCompile(`(?i)email:\s*(.+)(?:\s*\.\s*|$)`)
	phoneLabelRE = regexp.MustCompile(`(?i)(?:telephone|phone number|phone):\s*(.+)(?:\s*\.\s*|$)`)
)

type contactInfo struct {
	EmailAddresses []string
	PhoneNumbers   []string
	Websites       []string
}

// parseContactInfo scans a text blob for email addresses and phone numbers.
// It returns a contactInfo structure containing any matches. Duplicates are
// removed so callers can safely append results without checking for repeats.
func parseContactInfo(s string) contactInfo {
	ci := contactInfo{}
	if s == "" {
		return ci
	}

	// Extract emails from labeled sections
	for _, match := range emailLabelRE.FindAllStringSubmatch(s, -1) {
		if len(match) > 1 {
			emailsStr := strings.TrimSpace(match[1])
			emails := strings.Split(emailsStr, ";")
			for _, e := range emails {
				e = strings.TrimSpace(e)
				if e != "" {
					// Take first part if has space
					if idx := strings.Index(e, " "); idx > 0 {
						e = e[:idx]
					}
					ci.EmailAddresses = append(ci.EmailAddresses, e)
				}
			}
		}
	}

	// Extract phones from labeled sections
	for _, match := range phoneLabelRE.FindAllStringSubmatch(s, -1) {
		if len(match) > 1 {
			phonesStr := strings.TrimSpace(match[1])
			if phonesStr != "" {
				parts := strings.Fields(phonesStr)
				phoneParts := []string{}
				for _, p := range parts {
					if strings.Contains(p, ":") {
						break
					}
					phoneParts = append(phoneParts, p)
				}
				phone := strings.Join(phoneParts, " ")
				phone = strings.TrimSuffix(phone, ".")
				if phone != "" {
					ci.PhoneNumbers = append(ci.PhoneNumbers, phone)
				}
			}
		}
	}

	ci.Websites = websiteRE.FindAllString(s, -1)

	// deduplicate
	uniq := map[string]struct{}{}
	filtered := []string{}
	for _, e := range ci.EmailAddresses {
		if _, ok := uniq[e]; !ok {
			uniq[e] = struct{}{}
			filtered = append(filtered, e)
		}
	}
	ci.EmailAddresses = filtered

	uniq = map[string]struct{}{}
	filtered = []string{}
	for _, p := range ci.PhoneNumbers {
		if _, ok := uniq[p]; !ok {
			uniq[p] = struct{}{}
			filtered = append(filtered, p)
		}
	}
	ci.PhoneNumbers = filtered

	uniq = map[string]struct{}{}
	filtered = []string{}
	for _, w := range ci.Websites {
		if _, ok := uniq[w]; !ok {
			uniq[w] = struct{}{}
			filtered = append(filtered, w)
		}
	}
	ci.Websites = filtered

	return ci
}

// ToEntity converts a UNIndividual to the Moov search.Entity format.
func (p UNIndividual) ToEntity() search.Entity[search.Value] {
	// Construct full name from all parts
	nameParts := []string{p.FirstName, p.SecondName, p.ThirdName, p.FourthName}
	fullName := strings.TrimSpace(strings.Join(nameParts, " "))

	entity := search.Entity[search.Value]{
		SourceID:   fmt.Sprintf("UN-%s", p.DataID), // Prefix to avoid collisions
		Name:       fullName,
		Type:       search.EntityPerson,
		Source:     search.SourceUNCSL,
		SourceData: p,
	}

	// Ensure Person pointer is initialized before appending to its fields
	if entity.Person == nil {
		entity.Person = &search.Person{}
	}
	//set the full name in the Person struct as well for easier access
	entity.Person.Name = fullName

	// map gender if available
	if p.Gender != "" {
		entity.Person.Gender = search.Gender(p.Gender)
	} else if p.Comments != "" {
		// try to extract Gender: Male/Female from comments1 if not present
		if strings.Contains(p.Comments, "Gender:") {
			// simple heuristic: look for 'Gender:' and a word after it
			parts := strings.Split(p.Comments, "Gender:")
			if len(parts) > 1 {
				tok := strings.Fields(parts[1])
				if len(tok) > 0 {
					entity.Person.Gender = search.Gender(tok[0])
				}
			}
		}
	}

	// Map Aliases
	for _, alias := range p.Aliases {
		if alias.Name != "" {
			entity.Person.AltNames = append(entity.Person.AltNames, alias.Name)
		}
	}

	// Map Addresses
	for _, addr := range p.Addresses {
		entity.Addresses = append(entity.Addresses, search.Address{
			City:    addr.City,
			Country: addr.Country,
		})
	}

	// extract contact info from any address notes
	for _, addr := range p.Addresses {
		if addr.Note != "" {
			ci := parseContactInfo(addr.Note)
			entity.Contact.EmailAddresses = append(entity.Contact.EmailAddresses, ci.EmailAddresses...)
			entity.Contact.PhoneNumbers = append(entity.Contact.PhoneNumbers, ci.PhoneNumbers...)
			entity.Contact.Websites = append(entity.Contact.Websites, ci.Websites...)
		}
	}

	// parse comments for contacts as well
	if p.Comments != "" {
		ci := parseContactInfo(p.Comments)
		entity.Contact.EmailAddresses = append(entity.Contact.EmailAddresses, ci.EmailAddresses...)
		entity.Contact.PhoneNumbers = append(entity.Contact.PhoneNumbers, ci.PhoneNumbers...)
		entity.Contact.Websites = append(entity.Contact.Websites, ci.Websites...)
	}

	// extract birth date (prefer exact date, fall back to year)
	for _, bd := range p.BirthDates {
		if bd.Date != "" {
			if t, err := time.Parse("2006-01-02", bd.Date); err == nil {
				entity.Person.BirthDate = &t
				break
			}
		}
		if bd.Year != "" {
			if y, err := strconv.Atoi(bd.Year); err == nil {
				t := time.Date(y, time.January, 1, 0, 0, 0, 0, time.UTC)
				entity.Person.BirthDate = &t
				break
			}
		}
	}

	// extract birth place (first non-empty)
	for _, bp := range p.BirthPlaces {
		parts := []string{}
		if bp.City != "" {
			parts = append(parts, bp.City)
		}
		if bp.State != "" {
			parts = append(parts, bp.State)
		}
		if bp.Country != "" {
			parts = append(parts, bp.Country)
		}
		if len(parts) > 0 {
			entity.Person.PlaceOfBirth = strings.Join(parts, ", ")
			break
		}
	}

	// extract nationalities
	for _, n := range p.Nationalities {
		if n.Text != "" {
			entity.Person.Titles = append(entity.Person.Titles, n.Text)
		}
	}

	return entity.Normalize()
}

// ToEntity converts a UNEntity to the Moov search.Entity format.
func (e UNEntity) ToEntity() search.Entity[search.Value] {

	entity := search.Entity[search.Value]{
		SourceID:   fmt.Sprintf("UN-%s", e.DataID),
		Name:       e.FirstName, // Entities store their name in FIRST_NAME
		Type:       search.EntityBusiness,
		Source:     search.SourceUNCSL,
		SourceData: e,
	}

	// Ensure Business pointer is initialized before appending to its fields
	if entity.Business == nil {
		entity.Business = &search.Business{}
	}

	//set business name in the Business struct as well for easier access
	entity.Business.Name = e.FirstName
	for _, alias := range e.Aliases {
		if alias.Name != "" {
			entity.Business.AltNames = append(entity.Business.AltNames, alias.Name)
		}
	}

	// Map Addresses
	for _, addr := range e.Addresses {
		entity.Addresses = append(entity.Addresses, search.Address{
			City:    addr.City,
			Country: addr.Country,
		})
		if addr.Note != "" {
			ci := parseContactInfo(addr.Note)
			entity.Contact.EmailAddresses = append(entity.Contact.EmailAddresses, ci.EmailAddresses...)
			entity.Contact.PhoneNumbers = append(entity.Contact.PhoneNumbers, ci.PhoneNumbers...)
			entity.Contact.Websites = append(entity.Contact.Websites, ci.Websites...)
		}
	}

	// parse comments field for contact information
	if e.Comments != "" {
		ci := parseContactInfo(e.Comments)
		entity.Contact.EmailAddresses = append(entity.Contact.EmailAddresses, ci.EmailAddresses...)
		entity.Contact.PhoneNumbers = append(entity.Contact.PhoneNumbers, ci.PhoneNumbers...)
		entity.Contact.Websites = append(entity.Contact.Websites, ci.Websites...)
	}

	return entity.Normalize()
}
