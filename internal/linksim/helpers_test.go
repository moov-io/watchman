// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

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

	jane = (search.Entity[search.Value]{
		Name:   "Jane Doe",
		Type:   search.EntityPerson,
		Source: search.SourceUSOFAC,
		Person: &search.Person{
			Name: "Jane Doe",
			GovernmentIDs: []search.GovernmentID{
				{
					Type:       search.GovernmentIDPassport,
					Country:    "GB",
					Identifier: "999888777",
				},
			},
		},
		Addresses: []search.Address{
			{
				Line1:      "10 Downing Street",
				City:       "London",
				PostalCode: "SW1A 2AA",
				Country:    "GB",
			},
		},
	}).Normalize()
)
