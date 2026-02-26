// Copyright Bloomfielddev Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_un

import (
	"encoding/xml"
)

/*type CLSDATA struct {
	XMLName     xml.Name   `xml:"CONSOLIDATED_LIST"`
	Individuals []UNPerson `xml:"INDIVIDUALS>INDIVIDUAL"`
	Entities    []UNEntity `xml:"ENTITIES>ENTITY"`
}
*/

// UNConsolidatedList represents the root of the UN XML file
type CLSDATA struct {
	XMLName       xml.Name       `xml:"CONSOLIDATED_LIST"`
	DateGenerated string         `xml:"dateGenerated,attr"`
	Individuals   []UNIndividual `xml:"INDIVIDUALS>INDIVIDUAL"`
	Entities      []UNEntity     `xml:"ENTITIES>ENTITY"`
}

// UNIndividual captures specific data for persons
type UNIndividual struct {
	DataID          string       `xml:"DATAID"`
	FirstName       string       `xml:"FIRST_NAME"`
	SecondName      string       `xml:"SECOND_NAME"`
	ThirdName       string       `xml:"THIRD_NAME"`
	FourthName      string       `xml:"FOURTH_NAME"`
	UNListType      string       `xml:"UN_LIST_TYPE"`
	ReferenceNumber string       `xml:"REFERENCE_NUMBER"`
	ListedOn        string       `xml:"LISTED_ON"`
	Gender          string       `xml:"GENDER"`
	Comments        string       `xml:"COMMENTS1"`
	Nationalities   []Value      `xml:"NATIONALITY"`
	Aliases         []Alias      `xml:"INDIVIDUAL_ALIAS"`
	Addresses       []Address    `xml:"INDIVIDUAL_ADDRESS"`
	BirthDates      []BirthDate  `xml:"INDIVIDUAL_DATE_OF_BIRTH"`
	BirthPlaces     []BirthPlace `xml:"INDIVIDUAL_PLACE_OF_BIRTH"`
}

// UNEntity captures data for organizations/groups
type UNEntity struct {
	DataID          string    `xml:"DATAID"`
	FirstName       string    `xml:"FIRST_NAME"` // Entities use FIRST_NAME for the organization name
	UNListType      string    `xml:"UN_LIST_TYPE"`
	ReferenceNumber string    `xml:"REFERENCE_NUMBER"`
	ListedOn        string    `xml:"LISTED_ON"`
	Comments        string    `xml:"COMMENTS1"`
	Aliases         []Alias   `xml:"ENTITY_ALIAS"`
	Addresses       []Address `xml:"ENTITY_ADDRESS"`
}

// Supporting nested structs
type Value struct {
	Text string `xml:"VALUE"`
}

type Alias struct {
	Quality string `xml:"QUALITY"`
	Name    string `xml:"ALIAS_NAME"`
}

type Address struct {
	Street        string `xml:"STREET"`
	City          string `xml:"CITY"`
	StateProvince string `xml:"STATE_PROVINCE"`
	Country       string `xml:"COUNTRY"`
	Note          string `xml:"NOTE"`
}

type BirthDate struct {
	Type string `xml:"TYPE_OF_DATE"`
	Date string `xml:"DATE"`
	Year string `xml:"YEAR"`
}

type BirthPlace struct {
	City    string `xml:"CITY"`
	State   string `xml:"STATE_PROVINCE"`
	Country string `xml:"COUNTRY"`
}
