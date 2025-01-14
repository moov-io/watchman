// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/csl_us/gen/ENHANCED_XML"
	"github.com/moov-io/watchman/pkg/search"
)

func TestMapPersonFeaturesFromXML(t *testing.T) {
	featuresXML := `
        <features>
            <feature id="4698">
                <type featureTypeId="8">Birthdate</type>
                <versionId>1612</versionId>
                <value>1951</value>
                <valueDate id="757">
                    <fromDateBegin>1951-01-01</fromDateBegin>
                    <fromDateEnd>1951-01-01</fromDateEnd>
                    <toDateBegin>1951-12-31</toDateBegin>
                    <toDateEnd>1951-12-31</toDateEnd>
                    <isApproximate>false</isApproximate>
                    <isDateRange>false</isDateRange>
                </valueDate>
                <isPrimary>true</isPrimary>
            </feature>
            <feature id="4699">
                <type featureTypeId="9">Place of Birth</type>
                <versionId>1613</versionId>
                <value>Umm Tuba</value>
                <isPrimary>true</isPrimary>
            </feature>
        </features>`

	var features ENHANCED_XML.EntityFeatures
	if err := xml.Unmarshal([]byte(featuresXML), &features); err != nil {
		t.Fatalf("failed to unmarshal features XML: %v", err)
	}

	person := &search.Person{}
	mapPersonFeatures(&features, person)

	// Test birth date
	if person.BirthDate == nil {
		t.Fatal("BirthDate is nil, want non-nil")
	}
	wantYear := 1951
	if person.BirthDate.Year() != wantYear {
		t.Errorf("BirthDate.Year() = %d, want %d", person.BirthDate.Year(), wantYear)
	}

	// // Test place of birth
	// wantPlace := "Umm Tuba"
	// if person.PlaceOfBirth != wantPlace {
	// 	t.Errorf("PlaceOfBirth = %s, want %s", person.PlaceOfBirth, wantPlace)
	// }
}

func TestMapPersonFeaturesEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		featuresXML string
		wantDate    *time.Time
		wantPlace   string
		wantGender  search.Gender
	}{
		{
			name: "male gender",
			featuresXML: `
                <features>
                    <feature>
                        <type featureTypeId="91526">Gender</type>
                        <value>Male</value>
                        <isPrimary>true</isPrimary>
                    </feature>
                </features>`,
			wantGender: search.GenderMale,
		},
		{
			name: "female gender",
			featuresXML: `
                <features>
                    <feature>
                        <type featureTypeId="91527">Gender</type>
                        <value>Female</value>
                        <isPrimary>true</isPrimary>
                    </feature>
                </features>`,
			wantGender: search.GenderFemale,
		},
		{
			name: "non-primary features should be ignored",
			featuresXML: `
                <features>
                    <feature>
                        <type featureTypeId="8">Birthdate</type>
                        <valueDate>
                            <fromDateBegin>1951-01-01</fromDateBegin>
                        </valueDate>
                        <isPrimary>false</isPrimary>
                    </feature>
                    <feature>
                        <type featureTypeId="9">Place of Birth</type>
                        <value>Test Place</value>
                        <isPrimary>false</isPrimary>
                    </feature>
                </features>`,
		},
		{
			name: "invalid date should be handled",
			featuresXML: `
                <features>
                    <feature>
                        <type featureTypeId="8">Birthdate</type>
                        <valueDate>
                            <fromDateBegin>invalid-date</fromDateBegin>
                        </valueDate>
                        <isPrimary>true</isPrimary>
                    </feature>
                </features>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var features ENHANCED_XML.EntityFeatures
			if err := xml.Unmarshal([]byte(tt.featuresXML), &features); err != nil {
				t.Fatalf("failed to unmarshal features XML: %v", err)
			}

			person := &search.Person{}
			mapPersonFeatures(&features, person)

			// Check date
			if tt.wantDate == nil {
				if person.BirthDate != nil {
					t.Errorf("BirthDate = %v, want nil", person.BirthDate)
				}
			} else if !person.BirthDate.Equal(*tt.wantDate) {
				t.Errorf("BirthDate = %v, want %v", person.BirthDate, tt.wantDate)
			}

			// // Check place
			// if person.PlaceOfBirth != tt.wantPlace {
			// 	t.Errorf("PlaceOfBirth = %v, want %v", person.PlaceOfBirth, tt.wantPlace)
			// }

			// Check gender
			if person.Gender != tt.wantGender {
				t.Errorf("Gender = %v, want %v", person.Gender, tt.wantGender)
			}
		})
	}
}
