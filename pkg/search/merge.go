package search

import (
	"cmp"
	"fmt"
	"strings"

	"github.com/adamdecaf/merge"
)

func Merge[T Value](entities []Entity[T]) []Entity[T] {
	// Make a copy so we're free to modify
	out := make([]Entity[T], len(entities))
	copy(out, entities)

	return merge.Slices(getMergeKey, mergeEntities, out)
}

func getMergeKey[T Value](entity Entity[T]) string {
	return strings.ToLower(fmt.Sprintf("%s/%s/%s", entity.Source, entity.SourceID, entity.Type))
}

func mergeEntities[T Value](e1, e2 *Entity[T]) {
	*e1 = e1.merge(e2)
}

func (e *Entity[T]) merge(other *Entity[T]) Entity[T] {
	var out Entity[T]

	if e == nil && other == nil {
		return out
	}
	if e == nil {
		return *other
	}
	if other == nil {
		return *e
	}

	// Combine Name fields
	var altNames []string
	switch {
	case e.Name != "" && other.Name == "":
		out.Name = e.Name

	case e.Name == "" && other.Name != "":
		out.Name = other.Name

	default:
		// both populated or both empty
		out.Name = e.Name
		altNames = mergeStrings(altNames, []string{other.Name})
	}

	// Merge Basic fields
	out.Type = cmp.Or(e.Type, other.Type)
	out.Source = cmp.Or(e.Source, other.Source)
	out.SourceID = cmp.Or(e.SourceID, other.SourceID)

	// Merge type fields
	switch {
	case e.Person != nil && other.Person != nil:
		out.Person = &Person{
			Name:          e.Name,
			AltNames:      mergeStrings(altNames, e.Person.AltNames, other.Person.AltNames),
			Gender:        cmp.Or(e.Person.Gender, other.Person.Gender),
			BirthDate:     cmp.Or(e.Person.BirthDate, other.Person.BirthDate),
			PlaceOfBirth:  cmp.Or(e.Person.PlaceOfBirth, other.Person.PlaceOfBirth),
			DeathDate:     cmp.Or(e.Person.DeathDate, other.Person.DeathDate),
			Titles:        mergeStrings(e.Person.Titles, other.Person.Titles),
			GovernmentIDs: mergeGovernmentIDs(e.Person.GovernmentIDs, other.Person.GovernmentIDs),
		}

	case e.Business != nil && other.Business != nil:
		out.Business = &Business{
			Name:          e.Name,
			AltNames:      mergeStrings(altNames, e.Business.AltNames, other.Business.AltNames),
			Created:       cmp.Or(e.Business.Created, other.Business.Created),
			Dissolved:     cmp.Or(e.Business.Dissolved, other.Business.Dissolved),
			GovernmentIDs: mergeGovernmentIDs(e.Business.GovernmentIDs, other.Business.GovernmentIDs),
		}

	case e.Organization != nil && other.Organization != nil:
		out.Organization = &Organization{
			Name:          e.Name,
			AltNames:      mergeStrings(altNames, e.Organization.AltNames, other.Organization.AltNames),
			Created:       cmp.Or(e.Organization.Created, other.Organization.Created),
			Dissolved:     cmp.Or(e.Organization.Dissolved, other.Organization.Dissolved),
			GovernmentIDs: mergeGovernmentIDs(e.Organization.GovernmentIDs, other.Organization.GovernmentIDs),
		}

	case e.Aircraft != nil && other.Aircraft != nil:
		out.Aircraft = &Aircraft{
			Name:         e.Name,
			AltNames:     mergeStrings(altNames, e.Aircraft.AltNames, other.Aircraft.AltNames),
			Type:         cmp.Or(e.Aircraft.Type, other.Aircraft.Type),
			Flag:         cmp.Or(e.Aircraft.Flag, other.Aircraft.Flag),
			Built:        cmp.Or(e.Aircraft.Built, other.Aircraft.Built),
			ICAOCode:     cmp.Or(e.Aircraft.ICAOCode, other.Aircraft.ICAOCode),
			Model:        cmp.Or(e.Aircraft.Model, other.Aircraft.Model),
			SerialNumber: cmp.Or(e.Aircraft.SerialNumber, other.Aircraft.SerialNumber),
		}

	case e.Vessel != nil && other.Vessel != nil:
		out.Vessel = &Vessel{
			Name:                   e.Name,
			AltNames:               mergeStrings(altNames, e.Vessel.AltNames, other.Vessel.AltNames),
			IMONumber:              cmp.Or(e.Vessel.IMONumber, other.Vessel.IMONumber),
			Type:                   cmp.Or(e.Vessel.Type, other.Vessel.Type),
			Flag:                   cmp.Or(e.Vessel.Flag, other.Vessel.Flag),
			Built:                  cmp.Or(e.Vessel.Built, other.Vessel.Built),
			Model:                  cmp.Or(e.Vessel.Model, other.Vessel.Model),
			Tonnage:                cmp.Or(e.Vessel.Tonnage, other.Vessel.Tonnage),
			MMSI:                   cmp.Or(e.Vessel.MMSI, other.Vessel.MMSI),
			CallSign:               cmp.Or(e.Vessel.CallSign, other.Vessel.CallSign),
			GrossRegisteredTonnage: cmp.Or(e.Vessel.GrossRegisteredTonnage, other.Vessel.GrossRegisteredTonnage),
			Owner:                  cmp.Or(e.Vessel.Owner, other.Vessel.Owner),
		}
	}

	// Merge Contact
	out.Contact.EmailAddresses = mergeStrings(e.Contact.EmailAddresses, other.Contact.EmailAddresses)
	out.Contact.PhoneNumbers = mergeStrings(e.Contact.PhoneNumbers, other.Contact.PhoneNumbers)
	out.Contact.FaxNumbers = mergeStrings(e.Contact.FaxNumbers, other.Contact.FaxNumbers)
	out.Contact.Websites = mergeStrings(e.Contact.Websites, other.Contact.Websites)

	out.Addresses = mergeAddresses(e.Addresses, other.Addresses)
	out.CryptoAddresses = mergeCryptoAddresses(e.CryptoAddresses, other.CryptoAddresses)

	return out.Normalize()
}

func mergeStrings(ss ...[]string) []string {
	return merge.Slices(
		func(s string) string {
			return strings.ToLower(s)
		},
		nil, // don't merge, just unique
		ss...,
	)
}

func mergeGovernmentIDs(ids1, ids2 []GovernmentID) []GovernmentID {
	return merge.Slices(
		func(id GovernmentID) string {
			return strings.ToLower(fmt.Sprintf("%s/%s/%s", id.Country, id.Type, id.Identifier))
		},
		nil, // don't merge, just unique
		ids1, ids2,
	)
}

func mergeAddresses(a1, a2 []Address) []Address {
	// We're assuming that within two entities Line1 + Line2 is unique enough to be a unique address.
	// We want different Line1's and Line2's to be different addresses
	return merge.Slices(
		func(addr Address) string {
			return strings.ToLower(fmt.Sprintf("%s/%s", addr.Line1, addr.Line2))
		},
		func(a1, a2 *Address) {
			a1.Line1 = cmp.Or(a1.Line1, a2.Line1)
			a1.Line2 = cmp.Or(a1.Line2, a2.Line2)
			a1.City = cmp.Or(a1.City, a2.City)
			a1.PostalCode = cmp.Or(a1.PostalCode, a2.PostalCode)
			a1.State = cmp.Or(a1.State, a2.State)
			a1.Country = cmp.Or(a1.Country, a2.Country)
		},
		a1, a2,
	)
}

func mergeCryptoAddresses(c1, c2 []CryptoAddress) []CryptoAddress {
	return merge.Slices(
		func(addr CryptoAddress) string {
			return strings.ToLower(fmt.Sprintf("%s/%s", addr.Currency, addr.Address))
		},
		nil, // don't merge, just unique
		c1, c2,
	)
}
