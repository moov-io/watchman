// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/moov-io/base"
)

const (
	// csvFile is an OFAC CSVfile
	csvFile = ".csv"

	// txtFile is a DPL tab-delimited file
	// it should be a ".tsv" file, but BIS named it ".txt"...
	txtFile  = ".txt"
	txtDelim = '\t'

	// addressFile is an OFAC Specially Designated National (SDN) address File
	addressFile = "add.csv"
	// alternateIDFile is an OFAC Specially Designated National (SDN) alternate ID File
	alternateIDFile = "alt.csv"
	// speciallyDesignatedNationalFile is an OFAC Specially Designated National (SDN) File
	speciallyDesignatedNationalFile = "sdn.csv"
	// speciallyDesignatedNationalCommentsFile is an OFAC Specially Designated National (SDN) Comments File
	speciallyDesignatedNationalCommentsFile = "sdn_comments.csv"

	// deniedPersonsListFile is the Denied Persons List provided by the Bureau of Industry and Security, US Department of Commerce
	deniedPersonsListFile = "dpl.txt"
)

// Error strings specific to parsing/reading an OFAC file
var (
	msgFileExtension = "%v is an invalid file type"
	msgFileName      = "%v is an invalid file name"
)

// error returns a new ParseError based on err
func (r *Reader) parseError(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*base.ParseError); ok {
		return err
	}
	return &base.ParseError{
		Err: err,
	}
}

// Reader reads OFAC records from a CSV file and populates the associated arrays.
//
// For more details on the raw OFAC files see https://docs.moov.io/ofac/file-structure/
type Reader struct {
	// FileName is the name of the file
	FileName string `json:"fileName"`
	// Addresses returns an array of OFAC Specially Designated National Addresses
	Addresses []*Address `json:"address"`
	// AlternateIdentities returns an array of OFAC Specially Designated National Alternate Identity
	AlternateIdentities []*AlternateIdentity `json:"alternateIdentity"`
	// SDNs returns an array of OFAC Specially Designated Nationals
	SDNs []*SDN `json:"sdn"`
	// SDNComments returns an array of OFAC Specially Designated National Comments
	SDNComments []*SDNComments `json:"sdnComments"`
	// DPL returns an array of BIS Denied Persons
	DeniedPersons []*DPL
	// errors holds each error encountered when attempting to parse the file
	errors base.ErrorList
}

// Read will consume the file at r.FileName and attempt to parse it was a CSV OFAC file.
func (r *Reader) Read() error {
	ext := filepath.Ext(r.FileName)

	switch ext {
	case csvFile:
		if err := r.csvFile(); err != nil {
			return err
		}
	case txtFile:
		if err := r.txtFile(); err != nil {
			return err
		}
	default:
		msg := fmt.Sprintf(msgFileExtension, ext)
		r.errors.Add(r.parseError(errors.New(msg)))
		return r.errors
	}
	return nil
}

// csvFile parses an SDN, Address, AlternateID, SDN_Comments OFAC CSV file
func (r *Reader) csvFile() error {
	if strings.Contains(r.FileName, addressFile) {
		if err := r.csvAddressFile(); err != nil {
			return err
		}
	} else if strings.Contains(r.FileName, alternateIDFile) {
		if err := r.csvAlternateIdentityFile(); err != nil {
			return err
		}
	} else if strings.Contains(r.FileName, speciallyDesignatedNationalFile) {
		if err := r.csvSDNFile(); err != nil {
			return err
		}
	} else if strings.Contains(r.FileName, speciallyDesignatedNationalCommentsFile) {
		if err := r.csvSDNCommentsFile(); err != nil {
			return err
		}
	} else {
		msg := fmt.Sprintf(msgFileName, r.FileName)
		r.errors.Add(r.parseError(errors.New(msg)))
		return r.errors
	}
	return nil
}

func (r *Reader) csvAddressFile() error {
	// Open CSV file
	f, err := os.Open(r.FileName)
	if err != nil {
		return err
	}
	defer f.Close()

	// Read File into a Variable
	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err != nil && err == csv.ErrFieldCount {
			continue
		}
		if err == io.EOF { // TODO(Adam): add max line count break here also
			break
		}
		if len(record) != 6 {
			continue
		}

		record = replaceNull(record)
		addr := &Address{
			EntityID:                    record[0],
			AddressID:                   record[1],
			Address:                     record[2],
			CityStateProvincePostalCode: record[3],
			Country:                     record[4],
			AddressRemarks:              record[5],
		}
		r.Addresses = append(r.Addresses, addr)
	}
	return nil
}

func (r *Reader) csvAlternateIdentityFile() error {
	// Open CSV file
	f, err := os.Open(r.FileName)
	if err != nil {
		return err
	}
	defer f.Close()

	// Read File into a Variable
	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err != nil && err == csv.ErrFieldCount {
			continue
		}
		if err == io.EOF { // TODO(adam)
			break
		}
		if len(record) != 5 {
			continue
		}
		record = replaceNull(record)
		alt := &AlternateIdentity{
			EntityID:         record[0],
			AlternateID:      record[1],
			AlternateType:    record[2],
			AlternateName:    record[3],
			AlternateRemarks: record[4],
		}
		r.AlternateIdentities = append(r.AlternateIdentities, alt)
	}
	return nil
}

func (r *Reader) csvSDNFile() error {
	// Open CSV file
	f, err := os.Open(r.FileName)
	if err != nil {
		return err
	}
	defer f.Close()

	// Read File into a Variable
	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err != nil && err == csv.ErrFieldCount {
			continue
		}
		if err == io.EOF { // TODO(Adam): add max line count break here also
			break
		}
		if len(record) != 12 {
			continue
		}
		record = replaceNull(record)
		sdn := &SDN{
			EntityID:               record[0],
			SDNName:                record[1],
			SDNType:                record[2],
			Program:                record[3],
			Title:                  record[4],
			CallSign:               record[5],
			VesselType:             record[6],
			Tonnage:                record[7],
			GrossRegisteredTonnage: record[8],
			VesselFlag:             record[9],
			VesselOwner:            record[10],
			Remarks:                record[11],
		}
		r.SDNs = append(r.SDNs, sdn)
	}
	return nil
}

func (r *Reader) csvSDNCommentsFile() error {
	// Open CSV file
	f, err := os.Open(r.FileName)
	if err != nil {
		return err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	// Loop through lines & turn into object
	for _, csvLine := range lines {
		csvLine := replaceNull(csvLine)
		sdnComments := &SDNComments{
			EntityID:        csvLine[0],
			RemarksExtended: csvLine[1],
		}
		r.SDNComments = append(r.SDNComments, sdnComments)
	}
	return nil
}

func (r *Reader) txtFile() error {
	if strings.Contains(r.FileName, deniedPersonsListFile) {
		if err := r.txtDeniedPersonsFile(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Reader) txtDeniedPersonsFile() error {
	// open txt file
	f, err := os.Open(r.FileName)
	if err != nil {
		return err
	}
	defer f.Close()

	// create a new csv.Reader and set the delim char to txtDelim char
	reader := csv.NewReader(f)
	reader.Comma = txtDelim
	// Read File into a Variable
	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, txtLine := range lines {
		deniedPerson := &DPL{
			Name:           txtLine[0],
			StreetAddress:  txtLine[1],
			City:           txtLine[2],
			State:          txtLine[3],
			Country:        txtLine[4],
			PostalCode:     txtLine[5],
			EffectiveDate:  txtLine[6],
			ExpirationDate: txtLine[7],
			StandardOrder:  txtLine[8],
			LastUpdate:     txtLine[9],
			Action:         txtLine[10],
			FRCitation:     txtLine[11],
		}
		r.DeniedPersons = append(r.DeniedPersons, deniedPerson)
	}
	return nil
}

// replaceNull replaces a CSV field that contain -0- with "".  Null values for all four formats consist of "-0-"
// (ASCII characters 45, 48, 45).
func replaceNull(s []string) []string {
	for i := 0; i < len(s); i++ {
		s[i] = strings.TrimSpace(strings.Replace(s[i], "-0-", "", -1))
	}
	return s
}
