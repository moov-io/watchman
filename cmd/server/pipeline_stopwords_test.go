// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/watchman/pkg/csl"
	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/abadojack/whatlanggo"
)

func TestStopwordsEnv(t *testing.T) {
	if keepStopwords {
		t.Errorf("KEEP_STOPWORDS is set")
	}
}

func TestStopwords__detect(t *testing.T) {
	addrs := func(country string) []*ofac.Address {
		return []*ofac.Address{
			{
				Country: country,
			},
		}
	}

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

	for i := range cases {
		if lang := detectLanguage(cases[i].in, addrs(cases[i].country)); lang != cases[i].expected {
			t.Errorf("#%d in=%q country=%s lang=%v", i, cases[i].in, cases[i].country, lang)
		}
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
		if result != cases[i].expected {
			t.Errorf("\n#%d in=%q  lang=%v\ngot=%q", i, cases[i].in, cases[i].lang, result)
		}
	}
}

func TestStopwords__apply(t *testing.T) {
	cases := []struct {
		testName string
		in       *Name
		expected string
	}{
		{
			testName: "type missing",
			in:       &Name{Processed: "Trees and Trucks"},
			expected: "Trees and Trucks",
		},
		{
			testName: "alt name",
			in:       &Name{Processed: "Trees and Trucks", alt: &ofac.AlternateIdentity{}},
			expected: "trees trucks",
		},
		{
			testName: "sdn individual",
			in:       &Name{Processed: "Trees and Trucks", sdn: &ofac.SDN{SDNType: "individual"}},
			expected: "Trees and Trucks",
		},
		{
			testName: "sdn business",
			in:       &Name{Processed: "Trees and Trucks", sdn: &ofac.SDN{SDNType: "business"}},
			expected: "trees trucks",
		},
		{
			testName: "ssi individual",
			in:       &Name{Processed: "Trees and Trucks", ssi: &csl.SSI{Type: "individual"}},
			expected: "Trees and Trucks",
		},
		{
			testName: "ssi business",
			in:       &Name{Processed: "Trees and Trucks", ssi: &csl.SSI{Type: "business"}},
			expected: "trees trucks",
		},
	}

	for _, test := range cases {
		t.Run(test.testName, func(t *testing.T) {
			stopwords := stopwordsStep{}
			err := stopwords.apply(test.in)
			if err != nil {
				t.Errorf("\n#%v in=%v err=%v", test.testName, test.in, err)
			}

			if test.in.Processed != test.expected {
				t.Errorf("\n#%v expected=%v got=%v", test.testName, test.expected, test.in.Processed)
			}
		})
	}
}
