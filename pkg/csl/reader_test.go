package csl

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	csl, err := Read(filepath.Join("..", "..", "test", "testdata", "csl.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if csl == nil {
		t.Fatal("failed to parse csl.csv")
	}
	if len(csl.SSIs) != 26 { // test CSL csv file has 26 SSI entries
		t.Errorf("len(SSIs)=%d", len(csl.SSIs))
	}
	if len(csl.ELs) != 22 {
		t.Errorf("len(ELs)=%d", len(csl.ELs))
	}
}

func TestRead_invalidRow(t *testing.T) {
	csl, err := Read(filepath.Join("..", "..", "test", "testdata", "invalidFiles", "csl.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if csl == nil {
		t.Fatal("failed to parse csl.csv")
	}
	if len(csl.SSIs) != 1 {
		t.Errorf("len(SSIs)=%d", len(csl.SSIs))
	}
	if len(csl.ELs) != 1 {
		t.Errorf("len(ELs)=%d", len(csl.ELs))
	}
}

func Test_unmarshalSSI(t *testing.T) {
	// row from the live CSL
	record := []string{"Sectoral Sanctions Identifications List (SSI) - Treasury Department", "17254", "Entity", "UKRAINE-EO13662]; SYRIA", "AK TRANSNEFT OAO",
		"", "57 B. Polyanka ul., Moscow, 119180, RU; 57 Bolshaya. Polyanka, Moscow, 119180, RU", "", "", "", "", "", "", "", "", "", "", "", "",
		"For more information on directives, please visit the following link: http://www.treasury.gov/resource-center/sanctions/Programs/Pages/ukraine.aspx#directives.", "http://bit.ly/1QWTIfE",
		"OAO AK TRANSNEFT; AKTSIONERNAYA KOMPANIYA PO TRANSPORTUNEFTI TRANSNEFT OAO; OIL TRANSPORTING JOINT-STOCK COMPANY TRANSNEFT; TRANSNEFT, JSC; TRANSNEFT OJSC; TRANSNEFT",
		"", "", "", "", "http://bit.ly/1MLgou0", "1027700049486, Registration ID; 00044463, Government Gazette Number; 7706061801, Tax ID No.; transneft@ak.transneft.ru, Email Address; www.transneft.ru, Website; Subject to Directive 2, Executive Order 13662 Directive Determination -",
	}
	expectedSSI := &SSI{
		EntityID:       "17254",
		Type:           "Entity",
		Programs:       []string{"UKRAINE-EO13662", "SYRIA"},
		Name:           "AK TRANSNEFT OAO",
		Addresses:      []string{"57 B. Polyanka ul., Moscow, 119180, RU", "57 Bolshaya. Polyanka, Moscow, 119180, RU"},
		Remarks:        []string{"For more information on directives, please visit the following link: http://www.treasury.gov/resource-center/sanctions/Programs/Pages/ukraine.aspx#directives."},
		AlternateNames: []string{"OAO AK TRANSNEFT", "AKTSIONERNAYA KOMPANIYA PO TRANSPORTUNEFTI TRANSNEFT OAO", "OIL TRANSPORTING JOINT-STOCK COMPANY TRANSNEFT", "TRANSNEFT, JSC", "TRANSNEFT OJSC", "TRANSNEFT"},
		IDsOnRecord:    []string{"1027700049486, Registration ID", "00044463, Government Gazette Number", "7706061801, Tax ID No.", "transneft@ak.transneft.ru, Email Address", "www.transneft.ru, Website", "Subject to Directive 2, Executive Order 13662 Directive Determination -"},
		SourceListURL:  "http://bit.ly/1QWTIfE",
		SourceInfoURL:  "http://bit.ly/1MLgou0",
	}

	actualSSI := unmarshalSSI(record)

	if !reflect.DeepEqual(expectedSSI, actualSSI) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedSSI, actualSSI)
	}
}

func Test_unmarshalEL(t *testing.T) {
	record := []string{"Entity List (EL) - Bureau of Industry and Security", "", "", "", "GBNTT", "", "No. 34 Mansour Street, Tehran, IR", "73 FR 54506", "2008-09-22", "", "",
		"For all items subject to the EAR (See §744.11 of the EAR)", "Presumption of denial", "", "", "", "", "", "", "", "http://bit.ly/1L47xrV", "", "", "", "", "", "http://bit.ly/1L47xrV", ""}
	expectedEL := &EL{
		Name:               "GBNTT",
		AlternateNames:     []string{""},
		Addresses:          []string{"No. 34 Mansour Street, Tehran, IR"},
		StartDate:          "2008-09-22",
		LicenseRequirement: "For all items subject to the EAR (See §744.11 of the EAR)",
		LicensePolicy:      "Presumption of denial",
		FRNotice:           "73 FR 54506",
		SourceListURL:      "http://bit.ly/1L47xrV",
		SourceInfoURL:      "http://bit.ly/1L47xrV",
	}

	actualEL := unmarshalEL(record)

	if !reflect.DeepEqual(expectedEL, actualEL) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedEL, actualEL)
	}
}

func Test_expandField(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"", []string{""}},
		{"1021100731190, Registration ID     ; 00159025, Government Gazette Number; 1102024468, Tax ID No.; ukhta-tr.gazprom.ru, Website; azaharov@sgp.gazprom.ru, Email Address; Subject to Directive 4, Executive Order 13662 Directive Determination -",
			[]string{"1021100731190, Registration ID", "00159025, Government Gazette Number", "1102024468, Tax ID No.", "ukhta-tr.gazprom.ru, Website",
				"azaharov@sgp.gazprom.ru, Email Address", "Subject to Directive 4, Executive Order 13662 Directive Determination -"}},
		{"Yakimanka B. Street, Building 39, Moscow, 119049, RU; 27-29/1, building 6, Smolenskaya-Sennaya st., Moscow, 119121, RU",
			[]string{"Yakimanka B. Street, Building 39, Moscow, 119049, RU", "27-29/1, building 6, Smolenskaya-Sennaya st., Moscow, 119121, RU"}},
	}
	for _, test := range tests {
		if got := expandField(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("expandField() = %v, want %v", got, test.want)
		}
	}
}

func Test_expandProgramsList(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"", []string{""}},
		{"IFSR] [SDGT", []string{"IFSR", "SDGT"}}, // Sometimes the CSL "programs" column contains data like "IFSR] [SDGT" instead of "IFSR; SDGT"
		{"IFSR; SDGT", []string{"IFSR", "SDGT"}},
		{"CYBER2; CAATSA - RUSSIA", []string{"CYBER2", "CAATSA - RUSSIA"}},
	}
	for _, test := range tests {
		if got := expandProgramsList(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("expandField() = %v, want %v", got, test.want)
		}
	}
}
