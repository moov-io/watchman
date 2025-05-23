// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

func TestMapPersonFeaturesFromCSV(t *testing.T) {
	// Create a sample SanctionsEntry based on the CSV structure
	entry := SanctionsEntry{
		ID:           "9640",
		Name:         "Mohammed ABU TEIR",
		Type:         "Individual",
		DatesOfBirth: "1951",
		IDs:          "Gender, Male",
	}

	person := mapPerson(entry)

	// Test birth date
	if person.BirthDate == nil {
		t.Fatal("BirthDate is nil, want non-nil")
	}
	wantYear := 1951
	if person.BirthDate.Year() != wantYear {
		t.Errorf("BirthDate.Year() = %d, want %d", person.BirthDate.Year(), wantYear)
	}

	// Test gender
	if person.Gender != search.GenderMale {
		t.Errorf("Gender = %v, want %v", person.Gender, search.GenderMale)
	}

	// Note: PlaceOfBirth is not supported in search.Person; commented-out test is ignored
}

func TestMapPersonFeaturesEdgeCases(t *testing.T) {
	tests := []struct {
		name       string
		entry      SanctionsEntry
		wantDate   *time.Time
		wantGender search.Gender
	}{
		{
			name: "male gender",
			entry: SanctionsEntry{
				IDs: "Gender, Male",
			},
			wantGender: search.GenderMale,
		},
		{
			name: "female gender",
			entry: SanctionsEntry{
				IDs: "Gender, Female",
			},
			wantGender: search.GenderFemale,
		},
		{
			name: "full date",
			entry: SanctionsEntry{
				DatesOfBirth: "1951-01-01",
			},
			wantDate: func() *time.Time {
				t, _ := time.Parse("2006-01-02", "1951-01-01")
				return &t
			}(),
		},
		{
			name: "invalid date should be ignored",
			entry: SanctionsEntry{
				DatesOfBirth: "invalid-date",
			},
			wantDate: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			person := mapPerson(tt.entry)

			// Check date
			if tt.wantDate == nil {
				if person.BirthDate != nil {
					t.Errorf("BirthDate = %v, want nil", person.BirthDate)
				}
			} else if person.BirthDate == nil || !person.BirthDate.Equal(*tt.wantDate) {
				t.Errorf("BirthDate = %v, want %v", person.BirthDate, tt.wantDate)
			}

			// Check gender
			if person.Gender != tt.wantGender {
				t.Errorf("Gender = %v, want %v", person.Gender, tt.wantGender)
			}
		})
	}
}
