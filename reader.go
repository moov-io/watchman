package ofac

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/moov-io/base"
)

const (
	// csvFile is an OFAC CSVfile
	csvFile = ".csv"

	// addressFile is an OFAC Specially Designated National (SDN) address File
	addressFile = "add.csv"
	// alternateIDFile is an OFAC Specially Designated National (SDN) alternate ID File
	alternateIDFile = "alt.csv"
	// speciallyDesignatedNationalFile is an OFAC Specially Designated National (SDN) File
	speciallyDesignatedNationalFile = "sdn.csv"
	// speciallyDesignatedNationalCommentsFile is an OFAC Specially Designated National (SDN) Comments File
	speciallyDesignatedNationalCommentsFile = "sdn_comments.csv"
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

// Reader reads records from a CSV file.
type Reader struct {
	// FileName is the name of the file
	FileName string `json:"fileName"`
	// AddressArray returns an array of OFAC Specially Designated National Addresses
	AddressArray []Address `json:"addressArray"`
	// AddressArray returns an array of OFAC Specially Designated National Alternate Identity
	AlternateIdentityArray []AlternateIdentity `json:"alternateIdentityArray"`
	// SDNArray returns an array of OFAC Specially Designated Nationals
	SDNArray []SDN `json:"sdnArray"`
	// SDNCommentsArray returns an array of OFAC Specially Designated National Comments
	SDNCommentsArray []SDNComments `json:"sdnCommentsArray"`
	// errors holds each error encountered when attempting to parse the file
	errors base.ErrorList
}

func (r *Reader) Read() error {
	ext := filepath.Ext(r.FileName)

	switch ext {
	case csvFile:
		if err := r.csvFile(); err != nil {
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
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	// Loop through lines & turn into object
	for _, csvLine := range lines {
		//csvLine := replaceNull(csvLine)

		a := Address{
			EntityID:                    csvLine[0],
			AddressID:                   csvLine[1],
			Address:                     csvLine[2],
			CityStateProvincePostalCode: csvLine[3],
			Country:                     csvLine[4],
			AddressRemarks:              csvLine[5],
		}
		/*		fmt.Println(a.EntityID + " " + a.AddressID + " " + a.Address + " " + a.CityStateProvincePostalCode + " " +
				a.Country + " " + a.AddressRemarks)*/

		r.AddressArray = append(r.AddressArray, a)
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
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	// Loop through lines & turn into object
	for _, csvLine := range lines {
		//csvLine := replaceNull(csvLine)

		alt := AlternateIdentity{
			EntityID:         csvLine[0],
			AlternateID:      csvLine[1],
			AlternateType:    csvLine[2],
			AlternateName:    csvLine[3],
			AlternateRemarks: csvLine[4],
		}
		//fmt.Println(alt.EntityID + " " + alt.AlternateID + " " + alt.AlternateType + " " + alt.AlternateName + " " + alt.AlternateRemarks)

		r.AlternateIdentityArray = append(r.AlternateIdentityArray, alt)

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
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	// Loop through lines & turn into object
	for _, csvLine := range lines {
		//csvLine := replaceNull(csvLine)

		sdn := SDN{
			EntityID:               csvLine[0],
			SDNName:                csvLine[1],
			SDNType:                csvLine[2],
			Program:                csvLine[3],
			Title:                  csvLine[4],
			CallSign:               csvLine[5],
			VesselType:             csvLine[6],
			Tonnage:                csvLine[7],
			GrossRegisteredTonnage: csvLine[8],
			VesselFlag:             csvLine[9],
			VesselOwner:            csvLine[10],
			Remarks:                csvLine[11],
		}
		/*		fmt.Println(sdn.EntityID + " " + sdn.SDNName + " " + sdn.SDNType + " " + sdn.Program + " " + sdn.Title + " " +
				sdn.CallSign + " " + sdn.VesselType + " " + sdn.Tonnage + " " + sdn.GrossRegisteredTonnage + " " +
				sdn.VesselFlag + " " + sdn.VesselOwner + " " + sdn.Remarks)*/

		r.SDNArray = append(r.SDNArray, sdn)
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

		sdnComments := SDNComments{
			EntityID:        csvLine[0],
			RemarksExtended: csvLine[1],
		}
		//fmt.Println(sdnComments.EntityID + " " + sdnComments.RemarksExtended)

		r.SDNCommentsArray = append(r.SDNCommentsArray, sdnComments)
	}
	return nil
}

// replaceNull replaces a CSV field that contain -0- with "".  Null values for all four formats consist of "-0-"
// (ASCII characters 45, 48, 45).
func replaceNull(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "-0-") {
			s[i] = ""
		}
	}
	return s
}
