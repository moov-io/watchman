// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/abadojack/whatlanggo"
)

func TestStopwordsEnv(t *testing.T) {
	if keepStopwords {
		t.Errorf("KEEP_STOPWORDS is set")
	}
}

func TestStopwords__detect(t *testing.T) {
	cases := []struct {
		in       string
		country  string
		expected whatlanggo.Lang
	}{
		{"COLOMBIANA DE CERDOS LTDA.", "Colombia", whatlanggo.Spa},
		{"INVERSIONES LA QUINTA Y CIA. LTDA.", "Colombia", whatlanggo.Spa},
		{"COMITE' DE BIENFAISANCE ET DE SECOURS AUX PALESTINIENS", "France", whatlanggo.Fra}, //nolint:misspell
		{"WELFARE AND DEVELOPMENT ORGANIZATION OF JAMAAT-UD-DAWAH FOR QUR'AN AND SUNNAH", "Pakistan", whatlanggo.Eng},
		{"WAQFIYA RI'AYA AL-USRA AL-FILISTINYA WA AL-LUBNANYA", "Lebanon", whatlanggo.Eng},
		{"PREDUZECE ZA TRGOVINU NA VELIKO I MALO PARTIZAN TECH DOO BEOGRAD-SAVSKI VENAC", "Serbia", whatlanggo.Srp},
		{"OTKRYTOE AKTSIONERNOE OBSHCHESTVO VNESHNEEKONOMICHESKOE OBEDINENIE TEKHNOPROMEKSPORT", "Russia", whatlanggo.Rus},
		{"KONSTRUKTORSKOE BYURO PRIBOROSTROENIYA OTKRYTOE AKTSIONERNOE OBSHCHESTVO", "Russia", whatlanggo.Rus},
		{"INTERCONTINENTAL BAUMASCHINEN UND NUTZFAHRZEUGE HANDELS GMBH", "Germany", whatlanggo.Deu},
	}

	for _, tc := range cases {
		got := detectLanguage(tc.in, tc.country)
		require.Equal(t, tc.expected, got)
	}
}

func TestStopwords__clean(t *testing.T) {
	cases := []struct {
		in       string
		lang     whatlanggo.Lang
		expected string
	}{
		{"Trees and Trucks", whatlanggo.Eng, "trees trucks"},
		{"COLOMBIANA DE CERDOS LTDA.", whatlanggo.Spa, "colombiana cerdos ltda"},
		{"INVERSIONES LA QUINTA Y CIA. LTDA.", whatlanggo.Spa, "inversiones quinta y cia ltda"},
		{"COMITE' DE BIENFAISANCE ET DE SECOURS AUX PALESTINIENS", whatlanggo.Fra, "comite' bienfaisance secours palestiniens"}, //nolint:misspell
		{"WELFARE AND DEVELOPMENT ORGANIZATION OF JAMAAT-UD-DAWAH FOR QUR'AN AND SUNNAH", whatlanggo.Eng, "welfare development organization jamaat-ud-dawah qur'an sunnah"},
		{"WAQFIYA RI'AYA AL-USRA AL-FILISTINYA WA AL-LUBNANYA", whatlanggo.Eng, "waqfiya ri'aya al-usra al-filistinya wa al-lubnanya"},
		{"PREDUZECE ZA TRGOVINU NA VELIKO I MALO PARTIZAN TECH DOO BEOGRAD-SAVSKI VENAC", whatlanggo.Srp, "preduzece za trgovinu na veliko i malo partizan tech doo beograd-savski venac"},
		{"OTKRYTOE AKTSIONERNOE OBSHCHESTVO VNESHNEEKONOMICHESKOE OBEDINENIE TEKHNOPROMEKSPORT", whatlanggo.Rus, "otkrytoe aktsionernoe obshchestvo vneshneekonomicheskoe obedinenie tekhnopromeksport"},
		{"KONSTRUKTORSKOE BYURO PRIBOROSTROENIYA OTKRYTOE AKTSIONERNOE OBSHCHESTVO", whatlanggo.Rus, "konstruktorskoe byuro priborostroeniya otkrytoe aktsionernoe obshchestvo"},
		{"INTERCONTINENTAL BAUMASCHINEN UND NUTZFAHRZEUGE HANDELS GMBH", whatlanggo.Deu, "intercontinental baumaschinen nutzfahrzeuge handels gmbh"},
	}

	for i := range cases {
		result := removeStopwords(cases[i].in, cases[i].lang)
		require.Equal(t, cases[i].expected, result)
	}
}

func TestStopwords__apply(t *testing.T) {
	cases := []struct {
		in       string
		expected string
	}{
		{
			in:       "Trees and Trucks",
			expected: "trees trucks",
		},
		{
			in:       "11420 CORP.", // Issue 483 #1
			expected: "11420 corp",
		},
		{
			in:       "11,420.2-1 CORP.", // Issue 483 #2
			expected: "11,420.2-1 corp",
		},
		{
			in:       "11AA420 CORP.", // Issue 483 #3
			expected: "11aa420 corp",
		},
	}
	for _, test := range cases {
		got := RemoveStopwords(test.in, "")
		require.Equal(t, test.expected, got)
	}
}
