//go:build libpostal

package address

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	postal "github.com/openvenues/gopostal/parser"
	"github.com/stretchr/testify/require"
)

func TestParseAddress(t *testing.T) {
	cases := []struct {
		input    string
		expected search.Address
	}{
		{
			input: "101 Maple Street Apt 202 Bigcity, New York 11222",
			expected: search.Address{
				Line1:      "101 maple street",
				Line2:      "apt 202",
				City:       "bigcity",
				PostalCode: "11222",
				State:      "new york",
			},
		},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("%#v", tc.expected)

		t.Run(name, func(t *testing.T) {
			got := ParseAddress(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}

func TestOrganizeLibpostalComponents(t *testing.T) {
	cases := []struct {
		parts    []postal.ParsedComponent
		expected search.Address
	}{
		{
			parts: []postal.ParsedComponent{
				{Label: "house_number", Value: "101"},
				{Label: "road", Value: "Main Street"},
				{Label: "city", Value: "Springfield"},
				{Label: "state", Value: "Illinois"},
				{Label: "postcode", Value: "62704"},
				{Label: "country", Value: "United States"},
			},
			expected: search.Address{
				Line1:      "101 Main Street",
				City:       "Springfield",
				PostalCode: "62704",
				State:      "Illinois",
				Country:    "United States",
			},
		},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("%#v", tc.expected)

		t.Run(name, func(t *testing.T) {
			got := organizeLibpostalComponents(tc.parts)
			require.Equal(t, tc.expected, got)
		})
	}
}

func Benchmark_ParseAddress(b *testing.B) {
	inputs := []string{
		"Flat 7B, Tower 2, Ocean Financial Centre, 12 Marina Boulevard, Singapore 018982",
		"Room 1403, West Wing, Trading Complex No. 5, 47 Al Souq Street, Dubai, United Arab Emirates",
		"Office 892, Floor 8, Edificio Comercial Torres, Avenida Balboa y Calle 42, Panama City, Panama",
		"Unit 15, 3rd Floor, 123 Pyongyang Industrial Zone, Rangnang District, Pyongyang, DPRK",
		"Suite 405, Business Center Red Square, 17 Tverskaya Street, Moscow 125009, Russian Federation",
		"Warehouse 23, Port Zone B, Terminal 4, Latakia Port Complex, Latakia, Syria",
		"Office 78, Tehran Trade Tower, Block 2, Valiasr Street, Tehran 19395-4791, Iran",
		"Villa 15, Street 7, Block 4, Diplomatic Quarter, Caracas 1010, Venezuela",
		"Room 2201, Finance Plaza Building, 333 Lujiazui Ring Road, Shanghai 200120, China",
		"Suite 17, Victoria Business Park, 45 Harare Drive, Harare, Zimbabwe",
		"Office Complex Delta, Building C, Floor 5, 89 Minsk Boulevard, Minsk 220114, Belarus",
		"Unit 908, Golden Trade Center, 78 Yangon Port Road, Yangon 11181, Myanmar",
		"Floor 3, Al-Zawra Tower, Block 215, Baghdad Commercial District, Baghdad, Iraq",
		"Building 45, Industrial Zone 3, Damascus International Airport Road, Damascus, Syria",
		"Suite 301, Havana Trade Building, 67 Malecon Avenue, Havana 10400, Cuba",
		"Office 12, Floor 4, Conakry Commerce Center, Route du Niger, Conakry, Guinea",
		"Unit 55, Khartoum Business Complex, Al Gamhoria Avenue, Khartoum, Sudan",
		"Room 789, Floor 7, Trade Tower 3, Kim Il Sung Square, Pyongyang, DPRK",
		"Building 23, Floor 2, Sevastopol Maritime Complex, 45 Port Street, Sevastopol 99011",
		"Office 445, Tripoli Trade Center, Omar Al-Mukhtar Street, Tripoli, Libya",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var index atomic.Int32
		for pb.Next() {
			// Get next address in a thread-safe way
			ParseAddress(inputs[int(index.Add(1))%len(inputs)])
		}
	})
}

func Benchmark_ParseSingleAddress(b *testing.B) {
	address := "Flat 7B, Tower 2, Ocean Financial Centre, 12 Marina Boulevard, Singapore 018982"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseAddress(address)
	}
}

func Benchmark_ParseMultipleAddressesSequential(b *testing.B) {
	inputs := []string{
		"Flat 7B, Tower 2, Ocean Financial Centre, 12 Marina Boulevard, Singapore 018982",
		"Room 1403, West Wing, Trading Complex No. 5, 47 Al Souq Street, Dubai, United Arab Emirates",
		"Office 892, Floor 8, Edificio Comercial Torres, Avenida Balboa y Calle 42, Panama City, Panama",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseAddress(inputs[i%len(inputs)])
	}
}
