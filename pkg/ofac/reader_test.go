// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	testdata := func(fn string) map[string]io.ReadCloser {
		fd, err := os.Open(filepath.Join("..", "..", "test", "testdata", fn))
		if err != nil {
			t.Error(err)
		}
		return map[string]io.ReadCloser{fn: fd}
	}
	res, err := Read(testdata("add.csv"))
	require.NoError(t, err)

	require.Len(t, res.Addresses, 7379)
	require.Len(t, res.AlternateIdentities, 0)
	require.Len(t, res.SDNs, 0)
	require.Len(t, res.SDNComments, 0)

	res, err = Read(testdata("alt.csv"))
	require.NoError(t, err)
	require.Len(t, res.Addresses, 0)
	require.Len(t, res.AlternateIdentities, 3196)
	require.Len(t, res.SDNs, 0)
	require.Len(t, res.SDNComments, 0)

	res, err = Read(testdata("sdn.csv"))
	require.NoError(t, err)
	require.Len(t, res.Addresses, 0)
	require.Len(t, res.AlternateIdentities, 0)
	require.Len(t, res.SDNs, 7379)
	require.Len(t, res.SDNComments, 0)

	res, err = Read(testdata("sdn_comments.csv"))
	require.NoError(t, err)
	require.Len(t, res.Addresses, 0)
	require.Len(t, res.AlternateIdentities, 0)
	require.Len(t, res.SDNs, 0)
	require.Len(t, res.SDNComments, 13)
}

func TestReplaceNull(t *testing.T) {
	ans := replaceNull(nil)
	if ans != nil {
		t.Errorf("Got %v", ans)
	}
	ans = replaceNull([]string{" -0-"})
	if len(ans) != 1 || ans[0] != "" {
		t.Errorf("Got %v", ans)
	}
	ans = replaceNull([]string{"foo", " -0-"})
	if len(ans) != 2 || ans[0] != "foo" || ans[1] != "" {
		t.Errorf("Got %v", ans)
	}
}

func TestCleanPrgmsList(t *testing.T) {
	tests := []struct {
		prgms    string
		expected string
	}{
		{"SDGT] ", "SDGT"},
		{" SDGT] [IFSR", "SDGT; IFSR"},
		{"SDNTK] [FTO] [SDGT", "SDNTK; FTO; SDGT"},
		{"SDNTK] [FTO] [SDGT; IFSR]", "SDNTK; FTO; SDGT; IFSR"},
		{"[IFSR] [SDNTK] [FTO] [SDGT", "IFSR; SDNTK; FTO; SDGT"},
	}

	for _, test := range tests {
		if actual := cleanPrgmsList(test.prgms); actual != test.expected {
			t.Errorf("Wanted %q, got %q", test.expected, actual)
		}
	}
}

func TestSplitPrograms(t *testing.T) {
	if items := splitPrograms("SGDT"); !reflect.DeepEqual(items, []string{"SGDT"}) {
		t.Errorf("items=%v", items)
	}
	if items := splitPrograms("IRAN; SGDT"); !reflect.DeepEqual(items, []string{"IRAN", "SGDT"}) {
		t.Errorf("items=%v", items)
	}
	if items := splitPrograms("IFSR; SDNTK; FTO; SDGT"); !reflect.DeepEqual(items, []string{"IFSR", "SDNTK", "FTO", "SDGT"}) {
		t.Errorf("items=%v", items)
	}
}

func TestSDNComments(t *testing.T) {
	fd, err := os.CreateTemp("", "sdn-csv")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := fd.WriteString(`28264,"hone Number 8613314257947; alt. Phone Number 8618004121000; Identification Number 210302198701102136 (China); a.k.a. "blackjack1987"; a.k.a. "khaleesi"; Linked To: LAZARUS GROUP."`); err != nil {
		t.Fatal(err)
	}

	fd.Seek(0, 0)

	// read with lazy quotes enabled
	res, err := csvSDNCommentsFile(fd)
	require.NoError(t, err)
	require.Len(t, res.SDNComments, 1)

	comments, found := res.SDNComments["28264"]
	require.True(t, found)
	require.Len(t, comments, 1)

	comment := comments[0]
	require.NotEmpty(t, comment.RemarksExtended)
}

func TestSDN__remarks(t *testing.T) {
	// individual
	remarks := splitRemarks("DOB 12 Oct 1972; POB Corozal, Belize; Passport 0291622 (Belize); Linked To: D'S SUPERMARKET COMPANY LTD.")
	expected := []string{"12 Oct 1972"}
	require.ElementsMatch(t, expected, findRemarkValues(remarks, "DOB"))
	expected = []string{"0291622 (Belize)"}
	require.ElementsMatch(t, expected, findRemarkValues(remarks, "Passport"))

	// Contact info
	remarks = splitRemarks("Website www.nitc.co.ir; Email Address info@nitc.co.ir; alt. Email Address administrator@nitc.co.ir; IFCA Determination - Involved in the Shipping Sector; Additional Sanctions Information - Subject to Secondary Sanctions; Telephone (98)(21)(66153220); Telephone (98)(21)(23803202); Telephone (98)(21)(23803303); Telephone (98)(21)(66153224); Telephone (98)(21)(23802230); Telephone (98)(9121115315);  Telephone (98)(9128091642); Telephone (98)(9127389031); Fax (98)(21)(22224537); Fax (98)(21)(23803318); Fax (98)(21)(22013392); Fax (98)(21)(22058763).")
	expected = []string{"www.nitc.co.ir"}
	require.ElementsMatch(t, expected, findRemarkValues(remarks, "Website"))
	expected = []string{"info@nitc.co.ir", "administrator@nitc.co.ir"}
	require.ElementsMatch(t, expected, findRemarkValues(remarks, "Email Address"))
	expected = []string{"(98)(21)(66153220)", "(98)(21)(23803202)", "(98)(21)(23803303)", "(98)(21)(66153224)", "(98)(21)(23802230)", "(98)(9121115315)", "(98)(9128091642)", "(98)(9127389031)"}
	require.ElementsMatch(t, expected, findRemarkValues(remarks, "Telephone"))
	expected = []string{"(98)(21)(22224537)", "(98)(21)(23803318)", "(98)(21)(22013392)", "(98)(21)(22058763)"}
	require.ElementsMatch(t, expected, findRemarkValues(remarks, "Fax"))

	// Vessel
	remarks = splitRemarks("Former Vessel Flag Malta; alt. Former Vessel Flag Tuvalu; alt. Former Vessel Flag None Identified; alt. Former Vessel Flag Tanzania; Additional Sanctions Information - Subject to Secondary Sanctions; Vessel Registration Identification IMO 9187629; MMSI 572469210; Linked To: NATIONAL IRANIAN TANKER COMPANY.")
	expected = []string{"9187629"}
	require.ElementsMatch(t, expected, findRemarkValues(remarks, "Vessel Registration Identification IMO"))
	expected = []string{"572469210"}
	require.ElementsMatch(t, expected, findRemarkValues(remarks, "MMSI"))

	// Aircraft
	remarks = splitRemarks("Aircraft Construction Number (also called L/N or S/N or F/N) 8401; Aircraft Manufacture Date 1992; Aircraft Model IL76-TD; Aircraft Operator YAS AIR; Aircraft Manufacturer's Serial Number (MSN) 1023409321; Linked To: POUYA AIR.")
	expected = []string{"1992"}
	require.ElementsMatch(t, expected, findRemarkValues(remarks, "Manufacture Date"))
	expected = []string{"IL76-TD"}
	require.ElementsMatch(t, expected, findRemarkValues(remarks, "Model"))
	expected = []string{"(MSN) 1023409321"}
	require.ElementsMatch(t, expected, findRemarkValues(remarks, "Serial Number"))

	t.Run("error conditions", func(t *testing.T) {
		remarks = splitRemarks("")
		require.Len(t, findRemarkValues(remarks, ""), 0)
		require.Len(t, findRemarkValues(remarks, "DOB"), 0)

		remarks = splitRemarks("  ;  ;;;;;  ;   ;")
		require.Len(t, findRemarkValues(remarks, "DOB"), 0)
	})
}

func TestSDNComments_CryptoCurrencies(t *testing.T) {
	fd, err := os.CreateTemp("", "sdn-comments")
	require.NoError(t, err)

	_, err = fd.WriteString(`42496," alt. Digital Currency Address - XBT 12jVCWW1ZhTLA5yVnroEJswqKwsfiZKsax; alt. Digital Currency Address - XBT 1J378PbmTKn2sEw6NBrSWVfjZLBZW3DZem; alt. Digital Currency Address - XBT 18aqbRhHupgvC9K8qEqD78phmTQQWs7B5d; alt. Digital Currency Address - XBT 16ti2EXaae5izfkUZ1Zc59HMcsdnHpP5QJ; Secondary sanctions risk: North Korea Sanctions Regulations, sections 510.201 and 510.210; Transactions Prohibited For Persons Owned or Controlled By U.S. Financial Institutions: North Korea Sanctions Regulations section 510.214; Passport E59165201 (China) expires 01 Sep 2025; Identification Number 371326198812157611 (China); a.k.a. 'WAKEMEUPUPUP'; a.k.a. 'FAST4RELEASE'; Linked To: LAZARUS GROUP."`)
	require.NoError(t, err)
	fd.Seek(0, 0)

	sdn, err := csvSDNCommentsFile(fd)
	require.NoError(t, err)
	require.Len(t, sdn.SDNComments, 1)

	comments, found := sdn.SDNComments["42496"]
	require.True(t, found)
	require.Len(t, comments, 1)

	addresses := comments[0].DigitalCurrencyAddresses
	require.Len(t, addresses, 4)

	expected := []DigitalCurrencyAddress{
		{Currency: "XBT", Address: "12jVCWW1ZhTLA5yVnroEJswqKwsfiZKsax"},
		{Currency: "XBT", Address: "1J378PbmTKn2sEw6NBrSWVfjZLBZW3DZem"},
		{Currency: "XBT", Address: "18aqbRhHupgvC9K8qEqD78phmTQQWs7B5d"},
		{Currency: "XBT", Address: "16ti2EXaae5izfkUZ1Zc59HMcsdnHpP5QJ"},
	}
	require.ElementsMatch(t, expected, addresses)
}

func TestDownload__OFAC_Spillover(t *testing.T) {
	ctx := context.Background()
	logger := log.NewTestLogger()
	initialDir := filepath.Join("..", "..", "test", "testdata", "static")

	files, err := Download(ctx, logger, initialDir)
	require.NoError(t, err)

	res, err := Read(files)
	require.NoError(t, err)

	var found *SDN
	for _, sdn := range res.SDNs {
		if sdn.EntityID == "12300" {
			found = &sdn
			break
		}
	}
	require.NotNil(t, found)

	expected := `DOB 13 May 1965; alt. DOB 13 Apr 1968; alt. DOB 07 Jul 1964; POB Medellin, Colombia; alt. POB Marinilla, Antioquia, Colombia; alt. POB Ciudad Victoria, Tamaulipas, Mexico; Cedula No. 7548733 (Colombia); alt. Cedula No. 70163752 (Colombia); alt. Cedula No. 172489729-1 (Ecuador); Passport AL720622 (Colombia); R.F.C. CIVJ650513LJA (Mexico); alt. R.F.C. OUSV-640707 (Mexico); C.U.R.P. CIVJ650513HNEFLR06 (Mexico); alt. C.U.R.P. OUVS640707HTSSLR07 (Mexico); Matricula Mercantil No 181301-1 Cali (Colombia); alt. Matricula Mercantil No 405885 Bogota (Colombia); Linked To: BIO FORESTAL S.A.S.; Linked To: CUBI CAFE CLICK CUBE MEXICO, S.A. DE C.V.; Linked To: DOLPHIN DIVE SCHOOL S.A.; Linked To: GANADERIA LA SORGUITA S.A.S.; Linked To: GESTORES DEL ECUADOR GESTORUM S.A.; Linked To: INVERPUNTO DEL VALLE S.A.; Linked To: INVERSIONES CIFUENTES Y CIA. S. EN C.; Linked To: LE CLAUDE, S.A. DE C.V.; Linked To: OPERADORA NUEVA GRANADA, S.A. DE C.V.; Linked To: PARQUES TEMATICOS S.A.S.; Linked To: PROMO RAIZ S.A.S.; Linked To: RED MUNDIAL INMOBILIARIA, S.A. DE C.V.; Linked To: FUNDACION PARA EL BIENESTAR Y EL PORVENIR; Linked To: C.I. METALURGIA EXTRACTIVA DE COLOMBIA S.A.S.; Linked To: GRUPO MUNDO MARINO, S.A.; Linked To: C.I. DISERCOM S.A.S.; Linked To: C.I. OKCOFFEE COLOMBIA S.A.S.; Linked To: C.I. OKCOFFEE INTERNATIONAL S.A.S.; Linked To: FUNDACION OKCOFFEE COLOMBIA; Linked To: CUBICAFE S.A.S.; Linked To: HOTELES Y BIENES S.A.; Linked To: FUNDACION SALVA LA SELVA; Linked To: LINEA AEREA PUEBLOS AMAZONICOS S.A.S.; Linked To: DESARROLLO MINERO RESPONSABLE C.I. S.A.S.; Linked To: R D I S.A.`
	require.Equal(t, expected, found.Remarks)
}
