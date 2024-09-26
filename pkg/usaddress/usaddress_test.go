package usaddress

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStandardizeAddress_Example(t *testing.T) {
	rawAddress := "123 N Main St Apt 4B\nAnytown, California 12345-6789"
	standardizedAddress := StandardizeAddress(rawAddress)

	require.Equal(t, "123", standardizedAddress.PrimaryNumber)
	require.Equal(t, "", standardizedAddress.StreetPredir)
	require.Equal(t, "N MAIN", standardizedAddress.StreetName)
	require.Equal(t, "ST", standardizedAddress.StreetSuffix)
	require.Equal(t, "", standardizedAddress.StreetPostdir)
	require.Equal(t, "APT 4B", standardizedAddress.SecondaryUnit)
	require.Equal(t, "ANYTOWN", standardizedAddress.City)
	require.Equal(t, "CA", standardizedAddress.State)
	require.Equal(t, "12345", standardizedAddress.ZIPCode)
	require.Equal(t, "6789", standardizedAddress.Plus4)
	require.Equal(t, "", standardizedAddress.POBox)
}

func TestStandardizeAddress(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected Address

		validationError string
	}{
		{
			name:  "Basic Address",
			input: "123 Main Street\nAnytown, California 12345",
			expected: Address{
				PrimaryNumber: "123",
				StreetName:    "MAIN",
				StreetSuffix:  "ST",
				City:          "ANYTOWN",
				State:         "CA",
				ZIPCode:       "12345",
			},
		},
		{
			name:  "Address with Directional Prefix",
			input: "456 North Elm Road\nOthertown, Texas 78910",
			expected: Address{
				PrimaryNumber: "456",
				StreetPredir:  "N",
				StreetName:    "ELM",
				StreetSuffix:  "RD",
				City:          "OTHERTOWN",
				State:         "TX",
				ZIPCode:       "78910",
			},
		},
		{
			name:  "Address with Directional Suffix",
			input: "789 Pine Avenue Southwest\nSmalltown, Ohio 45678",
			expected: Address{
				PrimaryNumber: "789",
				StreetName:    "PINE",
				StreetSuffix:  "AVE",
				StreetPostdir: "SW",
				City:          "SMALLTOWN",
				State:         "OH",
				ZIPCode:       "45678",
			},
		},
		{
			name:  "Address with Secondary Unit",
			input: "101 Maple Street Apt 202\nBigcity, New York 11222",
			expected: Address{
				PrimaryNumber: "101",
				StreetName:    "MAPLE",
				StreetSuffix:  "ST",
				SecondaryUnit: "APT 202",
				City:          "BIGCITY",
				State:         "NY",
				ZIPCode:       "11222",
			},
		},
		{
			name:  "PO Box Address",
			input: "PO Box 12345\nCapitol City, Virginia 20100",
			expected: Address{
				POBox:   "PO BOX 12345",
				City:    "CAPITOL CITY",
				State:   "VA",
				ZIPCode: "20100",
			},
		},
		{
			name:  "Address with Rural Route",
			input: "RR 2 Box 152\nFarmtown, Kansas 67890",
			expected: Address{
				RuralRoute: "RR 2 BOX 152",
				City:       "FARMTOWN",
				State:      "KS",
				ZIPCode:    "67890",
			},
		},
		{
			name:  "Highway Contract Route",
			input: "HC 1 Box 23\nCountryside, Nebraska 68000",
			expected: Address{
				HighwayContract: "HC 1 BOX 23",
				City:            "COUNTRYSIDE",
				State:           "NE",
				ZIPCode:         "68000",
			},
		},
		{
			name:  "Address with Fractional Street Number",
			input: "12 1/2 Center Street\nMiddletown, Colorado 80452",
			expected: Address{
				PrimaryNumber: "12 1/2",
				StreetName:    "CENTER",
				StreetSuffix:  "ST",
				City:          "MIDDLETOWN",
				State:         "CO",
				ZIPCode:       "80452",
			},
		},
		{
			name:  "Address with Building and Floor",
			input: "500 Fifth Avenue Bldg 1 Fl 3\nMetropolis, Illinois 60601",
			expected: Address{
				PrimaryNumber: "500",
				StreetName:    "FIFTH",
				StreetSuffix:  "AVE",
				SecondaryUnit: "BLDG 1 FL 3",
				City:          "METROPOLIS",
				State:         "IL",
				ZIPCode:       "60601",
			},
		},
		{
			name:  "Address with Intersection (Not typically standard)",
			input: "Broadway & Main Street\nCityville, Pennsylvania 15000",
			expected: Address{
				StreetName:   "", // might be "BROADWAY & MAIN"
				StreetSuffix: "", // might be ST
				City:         "CITYVILLE",
				State:        "PA",
				ZIPCode:      "15000",
			},
			validationError: "address must contain either a PO Box, rural route, highway contract, or a valid street address with a primary number and street name",
		},
		{
			name:  "Address with Company Name (Should be ignored)",
			input: "Acme Corp\n123 Corporate Blvd\nBusiness City, Georgia 30303",
			expected: Address{
				PrimaryNumber: "123",
				StreetName:    "CORPORATE",
				StreetSuffix:  "BLVD",
				City:          "BUSINESS CITY",
				State:         "GA",
				ZIPCode:       "30303",
			},
		},
		{
			name:  "Address with Box Number Symbol (#)",
			input: "789 Market Street #123\nCommerce City, California 94105",
			expected: Address{
				PrimaryNumber: "789",
				StreetName:    "MARKET",
				StreetSuffix:  "ST",
				SecondaryUnit: "# 123",
				City:          "COMMERCE CITY",
				State:         "CA",
				ZIPCode:       "94105",
			},
		},
		{
			name:  "Military Address",
			input: "PSC 1234 Box 5678\nAPO AE 09012",
			expected: Address{
				POBox:   "PSC 1234 BOX 5678",
				City:    "APO",
				State:   "AE",
				ZIPCode: "09012",
			},
			validationError: "State 'AE' is not a valid US state abbreviation",
		},
		{
			name:  "Unique ZIP Code (No city/state)",
			input: "1600 Pennsylvania Avenue NW\n20500",
			expected: Address{
				PrimaryNumber: "1600",
				StreetName:    "PENNSYLVANIA",
				StreetSuffix:  "AVE",
				StreetPostdir: "NW",
				ZIPCode:       "20500",
			},
			validationError: "State is required; City is required",
		},
		{
			name:  "Address with Hyphenated Primary Number",
			input: "42-40 Bell Boulevard\nBayside, New York 11361",
			expected: Address{
				PrimaryNumber: "42-40",
				StreetName:    "BELL",
				StreetSuffix:  "BLVD",
				City:          "BAYSIDE",
				State:         "NY",
				ZIPCode:       "11361",
			},
		},
		{
			name:  "Address with Building Name",
			input: "Empire State Building\n350 Fifth Avenue\nNew York, NY 10118",
			expected: Address{
				PrimaryNumber: "350",
				StreetName:    "FIFTH",
				StreetSuffix:  "AVE",
				City:          "NEW YORK",
				State:         "NY",
				ZIPCode:       "10118",
			},
		},
		{
			name:  "Address with Numeric Street Name",
			input: "123 5TH AVE\nNEW YORK, NY 10001",
			expected: Address{
				PrimaryNumber: "123",
				StreetName:    "5TH",
				StreetSuffix:  "AVE",
				City:          "NEW YORK",
				State:         "NY",
				ZIPCode:       "10001",
			},
		},
		{
			name:  "Address with Care Of (C/O)",
			input: "C/O JOHN DOE\n456 MAIN ST\nANYTOWN, CA 90210",
			expected: Address{
				PrimaryNumber: "456",
				StreetName:    "MAIN",
				StreetSuffix:  "ST",
				City:          "ANYTOWN",
				State:         "CA",
				ZIPCode:       "90210",
			},
		},
		{
			name:  "Address with Building Name in Secondary Unit",
			input: "350 FIFTH AVE EMPIRE STATE BUILDING\nNEW YORK, NY 10118",
			expected: Address{
				PrimaryNumber: "350",
				StreetName:    "FIFTH",
				StreetSuffix:  "AVE",
				SecondaryUnit: "EMPIRE STATE BUILDING",
				City:          "NEW YORK",
				State:         "NY",
				ZIPCode:       "10118",
			},
		},
		{
			name:  "Address with Multiple Secondary Units",
			input: "123 MAIN ST APT 4B FL 2\nBIGCITY, TX 75001",
			expected: Address{
				PrimaryNumber: "123",
				StreetName:    "MAIN",
				StreetSuffix:  "ST",
				SecondaryUnit: "APT 4B FL 2",
				City:          "BIGCITY",
				State:         "TX",
				ZIPCode:       "75001",
			},
		},
		{
			name:  "Address with Extended ZIP+4 Code",
			input: "789 ELM ST\nSMALLTOWN, OH 45678-1234",
			expected: Address{
				PrimaryNumber: "789",
				StreetName:    "ELM",
				StreetSuffix:  "ST",
				City:          "SMALLTOWN",
				State:         "OH",
				ZIPCode:       "45678",
				Plus4:         "1234",
			},
		},
		{
			name:  "Address with Uncommon Street Suffix",
			input: "321 OAK PASS\nWOODLAND, CA 95695",
			expected: Address{
				PrimaryNumber: "321",
				StreetName:    "OAK",
				StreetSuffix:  "PASS",
				City:          "WOODLAND",
				State:         "CA",
				ZIPCode:       "95695",
			},
		},
		{
			name:  "Address with Directional Word in Street Name",
			input: "456 EASTWOOD DR\nLAKESIDE, MI 49116",
			expected: Address{
				PrimaryNumber: "456",
				StreetName:    "EASTWOOD",
				StreetSuffix:  "DR",
				City:          "LAKESIDE",
				State:         "MI",
				ZIPCode:       "49116",
			},
		},
		{
			name:  "Address with Special Characters in Street Name",
			input: "789 O'CONNOR ST\nDUBLIN, CA 94568",
			expected: Address{
				PrimaryNumber: "789",
				StreetName:    "O'CONNOR",
				StreetSuffix:  "ST",
				City:          "DUBLIN",
				State:         "CA",
				ZIPCode:       "94568",
			},
		},
		{
			name:  "Address with PO Box and Street Address",
			input: "PO BOX 123\n456 MAIN ST\nANYTOWN, IA 12345",
			expected: Address{
				PrimaryNumber: "456",
				StreetName:    "MAIN",
				StreetSuffix:  "ST",
				City:          "ANYTOWN",
				State:         "IA",
				ZIPCode:       "12345",
				POBox:         "PO BOX 123",
			},
		},
		{
			name:  "Address with In Care Of Symbol (℅)",
			input: "℅ JOHN SMITH\n123 MAPLE ST\nSPRINGFIELD, IL 62704",
			expected: Address{
				PrimaryNumber: "123",
				StreetName:    "MAPLE",
				StreetSuffix:  "ST",
				City:          "SPRINGFIELD",
				State:         "IL",
				ZIPCode:       "62704",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := StandardizeAddress(tc.input)
			require.Equal(t, tc.expected, result)

			err := result.Validate()
			msg := fmt.Sprintf("address: %v", result.String())
			if tc.validationError == "" {
				require.NoError(t, err, msg)
			} else {
				require.ErrorContains(t, err, tc.validationError, msg)
			}
		})
	}
}

func TestStandardizeAddress_InvalidCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Address
	}{
		{
			name:     "Empty String",
			input:    "",
			expected: Address{},
		},
		{
			name:     "Whitespace Only",
			input:    "   \n  \t",
			expected: Address{},
		},
		{
			name:  "Non-address Text",
			input: "Hello World!",
			expected: Address{
				StreetName: "HELLO WORLD",
			},
		},
		{
			name:  "Random Characters",
			input: "!@#$%^&*()_+",
			expected: Address{
				StreetName: "#&*()_+",
			},
		},
		{
			name:  "Incomplete Address",
			input: "Unknown Place",
			expected: Address{
				StreetName:   "UNKNOWN",
				StreetSuffix: "PL",
			},
		},
		{
			name:  "Only City and State",
			input: "Springfield, IL",
			expected: Address{
				StreetName: "SPRINGFIELD",
				State:      "IL",
			},
		},
		{
			name:  "Only ZIP Code",
			input: "12345",
			expected: Address{
				PrimaryNumber: "12345",
			},
		},
		{
			name:  "Invalid ZIP Code",
			input: "ABCDE",
			expected: Address{
				StreetName: "ABCDE",
			},
		},
		{
			name:  "Address with Special Symbols Only",
			input: "### $$$ %%% ^^^",
			expected: Address{
				StreetName: "###",
			},
		},
		{
			name:  "Address with Missing Street Name",
			input: "123\nAnytown, NY 10001",
			expected: Address{
				PrimaryNumber: "123",
				City:          "ANYTOWN",
				State:         "NY",
				ZIPCode:       "10001",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addr := StandardizeAddress(tt.input)
			require.Equal(t, tt.expected, addr)
		})
	}
}
