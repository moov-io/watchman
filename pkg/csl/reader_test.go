package csl

import (
	"compress/gzip"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	csl, err := ReadFile(filepath.Join("..", "..", "test", "testdata", "csl.csv"))
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

func TestRead__Large(t *testing.T) {
	fd, err := os.Open(filepath.Join("testdata", "consolidated.csv.gz"))
	require.NoError(t, err)

	reader, err := gzip.NewReader(fd)
	require.NoError(t, err)

	report, err := Parse(reader)
	require.NoError(t, err)
	require.NotNil(t, report)

	// Ensure we read each row as expected
	require.Len(t, report.ELs, 2001)
	require.Len(t, report.MEUs, 71)
	require.Len(t, report.SSIs, 290)
}

func TestRead_missingRow(t *testing.T) {
	fd, err := os.CreateTemp("", "csl-missing.csv")
	require.NoError(t, err)
	t.Cleanup(func() { os.Remove(fd.Name()) })

	_, err = fd.WriteString(`  \n invalid  \n  \n`)
	require.NoError(t, err)

	resp, err := ReadFile(fd.Name())
	require.NoError(t, err)

	require.Len(t, resp.ELs, 0)
	require.Len(t, resp.MEUs, 0)
	require.Len(t, resp.SSIs, 0)
}

func TestRead_invalidRow(t *testing.T) {
	csl, err := ReadFile(filepath.Join("..", "..", "test", "testdata", "invalidFiles", "csl.csv"))
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

func Test_unmarshalEL(t *testing.T) {
	record := []string{"Entity List (EL) - Bureau of Industry and Security", "", "", "", "GBNTT", "", "No. 34 Mansour Street, Tehran, IR", "73 FR 54506", "2008-09-22", "", "",
		"For all items subject to the EAR (See §744.11 of the EAR)", "Presumption of denial", "", "", "", "", "", "", "", "http://bit.ly/1L47xrV", "", "", "", "", "", "http://bit.ly/1L47xrV", ""}
	expectedEL := &EL{
		Name:               "GBNTT",
		AlternateNames:     nil,
		Addresses:          []string{"No. 34 Mansour Street, Tehran, IR"},
		StartDate:          "2008-09-22",
		LicenseRequirement: "For all items subject to the EAR (See §744.11 of the EAR)",
		LicensePolicy:      "Presumption of denial",
		FRNotice:           "73 FR 54506",
		SourceListURL:      "http://bit.ly/1L47xrV",
		SourceInfoURL:      "http://bit.ly/1L47xrV",
	}

	actualEL := unmarshalEL(record, 0)

	if !reflect.DeepEqual(expectedEL, actualEL) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedEL, actualEL)
	}
}

func Test_unmarshalMEU(t *testing.T) {
	input := strings.NewReader(strings.TrimSpace(`
26744194bd9b5cbec49db6ee29a4b53c697c7420,Military End User (MEU) List - Bureau of Industry and Security,,,,AECC Aviation Power Co. Ltd.,,"Xiujia Bay, Weiyong Dt, Xian, 710021, CN",85 FR 83799,2020-12-23,,,For any item subject to the EAR listed in supplement no. 2 to part 744.,The license application procedure and license review policy for entities specified in supplement no. 2 to part 744 is specified in §744.21(d) and (e).,,,,,,,,https://bit.ly/2XaGPYw,"",,,,,https://bit.ly/2XaGPYw,
baba9becd5dd994a2f9748dd051aeb144dc5a35e,Military End User (MEU) List - Bureau of Industry and Security,,,,AECC Beijing Institute of Aeronautical. Materials,,"No. 8 Hangcai Avenue, Haidian District, Beijing, CN",85 FR 83799,2020-12-23,,,For any item subject to the EAR listed in supplement no. 2 to part 744.,The license application procedure and license review policy for entities specified in supplement no. 2 to part 744 is specified in §744.21(d) and (e).,,,,,,,,https://bit.ly/2XaGPYw,"",,,,,https://bit.ly/2XaGPYw,
d54346ef81802673c1b1daeb2ca8bd5d13755abd,Military End User (MEU) List - Bureau of Industry and Security,,,,AECC China Gas Turbine Establishment,,"No. 1 Hangkong Road, Mianyang, Sichuan, CN",85 FR 83799,2020-12-23,,,For any item subject to the EAR listed in supplement no. 2 to part 744.,The license application procedure and license review policy for entities specified in supplement no. 2 to part 744 is specified in §744.21(d) and (e).,,,,,,,,https://bit.ly/2XaGPYw,"",,,,,https://bit.ly/2XaGPYw,`))

	report, err := Parse(input)
	require.NoError(t, err)
	require.Len(t, report.MEUs, 3)

	require.Equal(t, &MEU{
		EntityID:  "26744194bd9b5cbec49db6ee29a4b53c697c7420",
		Name:      "AECC Aviation Power Co. Ltd.",
		Addresses: "Xiujia Bay, Weiyong Dt, Xian, 710021, CN",
		FRNotice:  "85 FR 83799",
		StartDate: "2020-12-23",
		EndDate:   "",
	}, report.MEUs[0])

	require.Equal(t, &MEU{
		EntityID:  "d54346ef81802673c1b1daeb2ca8bd5d13755abd",
		Name:      "AECC China Gas Turbine Establishment",
		Addresses: "No. 1 Hangkong Road, Mianyang, Sichuan, CN",
		FRNotice:  "85 FR 83799",
		StartDate: "2020-12-23",
		EndDate:   "",
	}, report.MEUs[2])
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

	actualSSI := unmarshalSSI(record, 0)

	if !reflect.DeepEqual(expectedSSI, actualSSI) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedSSI, actualSSI)
	}
}

func Test_unmarshalUVL(t *testing.T) {
	// Record from official CSV
	record := []string{"f15fa805ff4ac5e09026f5e78011a1bb6b26dec2", "Unverified List (UVL) - Bureau of Industry and Security", "", "", "",
		"Atlas Sanatgaran", "", "Komitas 26/114, Yerevan, Armenia, AM", "", "", "", "", "", "", "", "", "", "", "", "", "", "http://bit.ly/1iwwTSJ", "", "", "", "", "",
		"http://bit.ly/1Qi4R7Z", "",
	}
	expectedUVL := &UVL{
		EntityID:      "f15fa805ff4ac5e09026f5e78011a1bb6b26dec2",
		Name:          "Atlas Sanatgaran",
		Addresses:     []string{"Komitas 26/114, Yerevan, Armenia, AM"},
		SourceListURL: "http://bit.ly/1iwwTSJ",
		SourceInfoURL: "http://bit.ly/1Qi4R7Z",
	}

	actualUVL := unmarshalUVL(record, 1)

	if !reflect.DeepEqual(expectedUVL, actualUVL) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedUVL, actualUVL)
	}
}

func Test_unmarshalISN(t *testing.T) {
	// Record from official CSV
	record := []string{"2d2db09c686e4829d0ef1b0b04145eec3d42cd88", "Nonproliferation Sanctions (ISN) - State Department", "", "",
		"E.O. 13382; Export-Import Bank Act; Nuclear Proliferation Prevention Act", "Abdul Qadeer Khan", "", "", "Vol. 74, No. 11, 01/16/09", "2009-01-09",
		"", "", "", "", "", "", "", "", "", "", "Associated with the A.Q. Khan Network", "http://bit.ly/1NuVFxV", "ZAMAN; Haydar", "", "", "", "", "http://bit.ly/1NuVFxV", "",
	}
	expectedISN := &ISN{
		EntityID:              "2d2db09c686e4829d0ef1b0b04145eec3d42cd88",
		Programs:              []string{"E.O. 13382", "Export-Import Bank Act", "Nuclear Proliferation Prevention Act"},
		Name:                  "Abdul Qadeer Khan",
		FederalRegisterNotice: "Vol. 74, No. 11, 01/16/09",
		StartDate:             "2009-01-09",
		Remarks:               []string{"Associated with the A.Q. Khan Network"},
		SourceListURL:         "http://bit.ly/1NuVFxV",
		AlternateNames:        []string{"ZAMAN", "Haydar"},
		SourceInfoURL:         "http://bit.ly/1NuVFxV",
	}

	actualISN := unmarshalISN(record, 1)

	if !reflect.DeepEqual(expectedISN, actualISN) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedISN, actualISN)
	}
}

func Test_unmarshalFSE(t *testing.T) {
	// Record from official CSV
	record := []string{"17526", "Foreign Sanctions Evaders (FSE) - Treasury Department", "17526", "Individual", "SYRIA; FSE-SY",
		"BEKTAS, Halis", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "https://bit.ly/1QWTIfE", "", "CH", "1966-02-13", "", "",
		"http://bit.ly/1N1docf", "CH, X0906223, Passport",
	}
	expectedFSE := &FSE{
		EntityID:      "17526",
		EntityNumber:  "17526",
		Type:          "Individual",
		Programs:      []string{"SYRIA", "FSE-SY"},
		Name:          "BEKTAS, Halis",
		Addresses:     nil,
		SourceListURL: "https://bit.ly/1QWTIfE",
		Citizenships:  "CH",
		DatesOfBirth:  "1966-02-13",
		SourceInfoURL: "http://bit.ly/1N1docf",
		IDs:           []string{"CH, X0906223, Passport"},
	}

	actualFSE := unmarshalFSE(record, 1)

	if !reflect.DeepEqual(expectedFSE, actualFSE) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedFSE, actualFSE)
	}
}

func Test_unmarshalPLC(t *testing.T) {
	// Record derived from official CSV
	record := []string{"9702", "Palestinian Legislative Council List (PLC) - Treasury Department", "9702", "Individual", "NS-PLC;Office of Misinformation", "SALAMEH, Salem",
		"", "123 Dunbar Street, Testerville, TX, Palestine", "", "", "", "", "", "", "", "", "", "", "", "", "HAMAS - Der al-Balah", "https://bit.ly/1QWTIfE", "SALAMEH, Salem Ahmad Abdel Hadi", "", "1951",
		"", "", "http://bit.ly/2tjOLpx", "",
	}

	expectedPLC := &PLC{
		EntityID:       "9702",
		EntityNumber:   "9702",
		Type:           "Individual",
		Programs:       []string{"NS-PLC", "Office of Misinformation"},
		Name:           "SALAMEH, Salem",
		Addresses:      []string{"123 Dunbar Street, Testerville, TX, Palestine"},
		Remarks:        "HAMAS - Der al-Balah",
		SourceListURL:  "https://bit.ly/1QWTIfE",
		AlternateNames: []string{"SALAMEH, Salem Ahmad Abdel Hadi"},
		DatesOfBirth:   "1951",
		PlacesOfBirth:  "",
		SourceInfoURL:  "http://bit.ly/2tjOLpx",
	}

	actualPLC := unmarshalPLC(record, 1)

	if !reflect.DeepEqual(expectedPLC, actualPLC) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedPLC, actualPLC)
	}
}

func Test_unmarshalCAP(t *testing.T) {
	// Record derived from official CSV
	record := []string{"20002", "Capta List (CAP) - Treasury Department", "20002", "Entity", "UKRAINE-EO13662; RUSSIA-EO14024", "BM BANK PUBLIC JOINT STOCK COMPANY",
		"", "Bld 3 8/15, Rozhdestvenka St., Moscow, 107996, RU", "", "", "", "", "", "", "", "", "", "", "", "",
		"All offices worldwide; for more information on directives, please visit the following link: https://www.treasury.gov/resource-center/sanctions/Programs/Pages/ukraine.aspx#directives; (Linked To: VTB BANK PUBLIC JOINT STOCK COMPANY)",
		"", "BM BANK JSC; BM BANK AO; AKTSIONERNOE OBSHCHESTVO BM BANK; PAO BM BANK; BANK MOSKVY PAO; BANK OF MOSCOW; AKTSIONERNY KOMMERCHESKI BANK BANK MOSKVY OTKRYTOE AKTSIONERNOE OBSCHCHESTVO; JOINT STOCK COMMERCIAL BANK - BANK OF MOSCOW OPEN JOINT STOCK COMPANY",
		"", "", "", "", "http://bit.ly/2PqohAD", "RU, 1027700159497, Registration Number; RU, 29292940, Government Gazette Number; MOSWRUMM, SWIFT/BIC; www.bm.ru, Website; Subject to Directive 1, Executive Order 13662 Directive Determination -; 044525219, BIK (RU); Financial Institution, Target Type",
	}

	expectedCAP := &CAP{
		EntityID:      "20002",
		EntityNumber:  "20002",
		Type:          "Entity",
		Programs:      []string{"UKRAINE-EO13662", "RUSSIA-EO14024"},
		Name:          "BM BANK PUBLIC JOINT STOCK COMPANY",
		Addresses:     []string{"Bld 3 8/15, Rozhdestvenka St., Moscow, 107996, RU"},
		Remarks:       []string{"All offices worldwide", "for more information on directives, please visit the following link: https://www.treasury.gov/resource-center/sanctions/Programs/Pages/ukraine.aspx#directives", "(Linked To: VTB BANK PUBLIC JOINT STOCK COMPANY)"},
		SourceListURL: "",
		AlternateNames: []string{"BM BANK JSC", "BM BANK AO", "AKTSIONERNOE OBSHCHESTVO BM BANK",
			"PAO BM BANK", "BANK MOSKVY PAO", "BANK OF MOSCOW",
			"AKTSIONERNY KOMMERCHESKI BANK BANK MOSKVY OTKRYTOE AKTSIONERNOE OBSCHCHESTVO",
			"JOINT STOCK COMMERCIAL BANK - BANK OF MOSCOW OPEN JOINT STOCK COMPANY"},
		SourceInfoURL: "http://bit.ly/2PqohAD",
		IDs: []string{"RU, 1027700159497, Registration Number",
			"RU, 29292940, Government Gazette Number",
			"MOSWRUMM, SWIFT/BIC",
			"www.bm.ru, Website",
			"Subject to Directive 1, Executive Order 13662 Directive Determination -",
			"044525219, BIK (RU)",
			"Financial Institution, Target Type"},
	}

	actualCAP := unmarshalCAP(record, 1)

	if !reflect.DeepEqual(expectedCAP, actualCAP) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedCAP, actualCAP)
	}
}

func Test_unmarshalDTC(t *testing.T) {
	// Record derived from official CSV
	record := []string{"d44d88d0265d93927b9ff1c13bbbb7c7db64142c", "ITAR Debarred (DTC) - State Department", "", "", "",
		"Yasmin Ahmed", "", "", "69 FR 17468", "", "", "", "", "", "", "", "", "", "", "", "",
		"http://bit.ly/307FuRQ", "Yasmin Tariq; Fatimah Mohammad", "", "", "", "", "http://bit.ly/307FuRQ",
	}

	expectedDTC := &DTC{
		EntityID:              "d44d88d0265d93927b9ff1c13bbbb7c7db64142c",
		Name:                  "Yasmin Ahmed",
		FederalRegisterNotice: "69 FR 17468",
		SourceListURL:         "http://bit.ly/307FuRQ",
		AlternateNames:        []string{"Yasmin Tariq", "Fatimah Mohammad"},
		SourceInfoURL:         "http://bit.ly/307FuRQ",
	}

	actualDTC := unmarshalDTC(record, 1)

	if !reflect.DeepEqual(expectedDTC, actualDTC) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedDTC, actualDTC)
	}
}

func Test_unmarshalCMIC(t *testing.T) {
	// Record derived from official CSV
	record := []string{"32091", "Non-SDN Chinese Military-Industrial Complex Companies List (CMIC) - Treasury Department", "32091", "Entity",
		"CMIC-EO13959", "PROVEN HONOUR CAPITAL LIMITED", "", "C/O Vistra Corporate Services Centre, Wickhams Cay II, Road Town, VG1110, VG",
		"", "", "", "", "", "", "", "", "", "", "", "", "(Linked To: HUAWEI INVESTMENT & HOLDING CO., LTD.)", "https://bit.ly/1QWTIfE",
		"PROVEN HONOUR CAPITAL LTD; PROVEN HONOUR", "", "", "", "", "https://bit.ly/3zsMQ4n",
		"Proven Honour Capital Ltd, Issuer Name; Proven Honour Capital Limited, Issuer Name; XS1233275194, ISIN; HK0000216777, ISIN; Private Company, Target Type; XS1401816761, ISIN; HK0000111952, ISIN; 03 Jun 2021, Listing Date (CMIC); 02 Aug 2021, Effective Date (CMIC); 03 Jun 2022, Purchase/Sales For Divestment Date (CMIC)",
	}

	expectedCMIC := &CMIC{
		EntityID:       "32091",
		EntityNumber:   "32091",
		Type:           "Entity",
		Programs:       []string{"CMIC-EO13959"},
		Name:           "PROVEN HONOUR CAPITAL LIMITED",
		Addresses:      []string{"C/O Vistra Corporate Services Centre, Wickhams Cay II, Road Town, VG1110, VG"},
		Remarks:        []string{"(Linked To: HUAWEI INVESTMENT & HOLDING CO., LTD.)"},
		SourceListURL:  "https://bit.ly/1QWTIfE",
		AlternateNames: []string{"PROVEN HONOUR CAPITAL LTD", "PROVEN HONOUR"},
		SourceInfoURL:  "https://bit.ly/3zsMQ4n",
		IDs: []string{"Proven Honour Capital Ltd, Issuer Name", "Proven Honour Capital Limited, Issuer Name", "XS1233275194, ISIN",
			"HK0000216777, ISIN", "Private Company, Target Type", "XS1401816761, ISIN", "HK0000111952, ISIN", "03 Jun 2021, Listing Date (CMIC)",
			"02 Aug 2021, Effective Date (CMIC)", "03 Jun 2022, Purchase/Sales For Divestment Date (CMIC)"},
	}

	actualCMIC := unmarshalCMIC(record, 1)

	if !reflect.DeepEqual(expectedCMIC, actualCMIC) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedCMIC, actualCMIC)
	}
}

func Test_unmarshalNS_MBS(t *testing.T) {
	// Record derived from official CSV
	record := []string{"17016", "Non-SDN Menu-Based Sanctions List (NS-MBS List) - Treasury Department", "17016", "Entity", "UKRAINE-EO13662; MBS",
		"GAZPROMBANK JOINT STOCK COMPANY", "", "16 Nametkina Street, Bldg. 1, Moscow, 117420, RU", "", "", "", "", "", "", "", "", "", "", "", "",
		"For more information on directives, please visit the following link: http://www.treasury.gov/resource-center/sanctions/Programs/Pages/ukraine.aspx#directives.",
		"", "GAZPROMBANK OPEN JOINT STOCK COMPANY; BANK GPB JSC; GAZPROMBANK AO; JOINT STOCK BANK OF THE GAS INDUSTRY GAZPROMBANK", "", "", "", "", "https://bit.ly/2MbsybU",
		"RU, 1027700167110, Registration Number; RU, 09807684, Government Gazette Number; RU, 7744001497, Tax ID No.; www.gazprombank.ru, Website; GAZPRUMM, SWIFT/BIC; Subject to Directive 1, Executive Order 13662 Directive Determination -; Subject to Directive 3 - All transactions in, provision of financing for, and other dealings in new debt of longer than 14 days maturity or new equity where such new debt or new equity is issued on or after the 'Effective Date (EO 14024 Directive)' associated with this name are prohibited., Executive Order 14024 Directive Information; 31 Jul 1990, Organization Established Date; 24 Feb 2022, Listing Date (EO 14024 Directive 3):; 26 Mar 2022, Effective Date (EO 14024 Directive 3):; For more information on directives, please visit the following link: https://home.treasury.gov/policy-issues/financial-sanctions/sanctions-programs-and-country-information/russian-harmful-foreign-activities-sanctions#directives, Executive Order 14024 Directive Information -",
	}

	expectedNS_MBS := &NS_MBS{
		EntityID:       "17016",
		EntityNumber:   "17016",
		Type:           "Entity",
		Programs:       []string{"UKRAINE-EO13662", "MBS"},
		Name:           "GAZPROMBANK JOINT STOCK COMPANY",
		Addresses:      []string{"16 Nametkina Street, Bldg. 1, Moscow, 117420, RU"},
		Remarks:        []string{"For more information on directives, please visit the following link: http://www.treasury.gov/resource-center/sanctions/Programs/Pages/ukraine.aspx#directives."},
		AlternateNames: []string{"GAZPROMBANK OPEN JOINT STOCK COMPANY", "BANK GPB JSC", "GAZPROMBANK AO", "JOINT STOCK BANK OF THE GAS INDUSTRY GAZPROMBANK"},
		SourceInfoURL:  "https://bit.ly/2MbsybU",
		IDs: []string{"RU, 1027700167110, Registration Number", "RU, 09807684, Government Gazette Number", "RU, 7744001497, Tax ID No.",
			"www.gazprombank.ru, Website", "GAZPRUMM, SWIFT/BIC", "Subject to Directive 1, Executive Order 13662 Directive Determination -",
			"Subject to Directive 3 - All transactions in, provision of financing for, and other dealings in new debt of longer than 14 days maturity or new equity where such new debt or new equity is issued on or after the 'Effective Date (EO 14024 Directive)' associated with this name are prohibited., Executive Order 14024 Directive Information",
			"31 Jul 1990, Organization Established Date", "24 Feb 2022, Listing Date (EO 14024 Directive 3):", "26 Mar 2022, Effective Date (EO 14024 Directive 3):",
			"For more information on directives, please visit the following link: https://home.treasury.gov/policy-issues/financial-sanctions/sanctions-programs-and-country-information/russian-harmful-foreign-activities-sanctions#directives, Executive Order 14024 Directive Information -"},
	}

	actualNB_MBS := unmarshalNS_MBS(record, 1)

	if !reflect.DeepEqual(expectedNS_MBS, actualNB_MBS) {
		t.Errorf("Expected: %#v\nFound: %#v\n", expectedNS_MBS, actualNB_MBS)
	}
}

func Test__Issue326EL(t *testing.T) {
	fd, err := os.CreateTemp("", "csl*.csv")
	if err != nil {
		t.Fatal(err)
	}
	fd.WriteString(`764ecc9bd00a36930e6bfba2e65ffe3f8e96a123,Entity List (EL) - Bureau of Industry and Security,,,,Huawei Cloud Beijing,,"Beijing, CN",85 FR 51603,2020-08-20,,,"For all items subject to the EAR, see §§ 736.2(b)(3)(vi), and 744.11 of the EAR, EXCEPT for technology subject to the EAR that is designated as EAR99, or controlled on the Commerce Control List for anti-terrorism reasons only, when released to members of a standards organization (see §772.1) for the purpose of contributing to the revision or development of a standard (see §772.1).",Presumption of denial.,,,,,,,,http://bit.ly/1L47xrV,"",,,,,http://bit.ly/1L47xrV,`)
	if err := fd.Sync(); err != nil {
		t.Fatal(err)
	}

	// read the line back
	csl, err := ReadFile(fd.Name())
	if err != nil {
		t.Fatal(err)
	}
	if len(csl.ELs) != 1 {
		t.Fatalf("unexpected: %#v", csl.ELs)
	}

	// Compare the EntityList item
	el := csl.ELs[0]
	assert.Equal(t, el.ID, "764ecc9bd00a36930e6bfba2e65ffe3f8e96a123")
	assert.Equal(t, el.Name, "Huawei Cloud Beijing")
	assert.Empty(t, el.AlternateNames)
	assert.Equal(t, el.Addresses, []string{"Beijing, CN"})
	assert.Equal(t, el.StartDate, "2020-08-20")
}

func Test_expandField(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"", nil},
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
		{"", nil},
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

func TestCSL__UniqueIDs(t *testing.T) {
	// CSL datafiles have added a unique identifier as the first column.
	// We need verify the old and new file formats can be parsed.

	records, err := ReadFile(filepath.Join("..", "..", "test", "testdata", "csl-unique-ids.csv"))
	if err != nil {
		t.Fatal(err)
	}

	if n := len(records.SSIs); n != 290 {
		t.Errorf("got %d SSI records", n)
	}
	if n := len(records.ELs); n != 1332 {
		t.Errorf("got %d EL records", n)
	}
}
