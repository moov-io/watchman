package linksim

import (
	"github.com/moov-io/watchman/pkg/search"
)

var (
	john = (search.Entity[search.Value]{
		Name:   "John Smith",
		Type:   search.EntityPerson,
		Source: search.SourceUSOFAC,
		Person: &search.Person{
			Name: "John Smith",
			GovernmentIDs: []search.GovernmentID{
				{
					Type:       search.GovernmentIDPassport,
					Country:    "US",
					Identifier: "1234567890",
				},
			},
		},
		Contact: search.ContactInfo{
			EmailAddresses: []string{"john.smith123@example.com"},
		},
		Addresses: []search.Address{
			{
				Line1:      "541 First St",
				City:       "Anytown",
				State:      "CA",
				PostalCode: "90210",
				Country:    "US",
			},
		},
	}).Normalize()

	johnathon = (search.Entity[search.Value]{
		Name:   "Johnathon Smith",
		Type:   search.EntityPerson,
		Source: search.SourceUSOFAC,
		Person: &search.Person{
			Name: "Johnathon Smith",
			GovernmentIDs: []search.GovernmentID{
				{
					Type:       search.GovernmentIDPassport,
					Country:    "US",
					Identifier: "1234567890",
				},
			},
		},
		Contact: search.ContactInfo{
			EmailAddresses: []string{"johnathon.smith123@example.com"},
		},
		Addresses: []search.Address{
			{
				Line1:      "541 First St",
				Line2:      "Apt 301",
				City:       "Anytown",
				State:      "CA",
				PostalCode: "90210",
				Country:    "US",
			},
		},
	}).Normalize()
)
