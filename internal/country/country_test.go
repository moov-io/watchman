package country_test

import (
	"testing"

	"github.com/moov-io/watchman/internal/country"

	"github.com/stretchr/testify/require"
)

func TestNormalizeCountry(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{input: "us", expected: "United States"},
		{input: "US", expected: "United States"},
		{input: "united states", expected: "United States"},
		{input: "UNITED STATES", expected: "United States"},
		{input: "United States of America", expected: "United States"},
		{input: "usa", expected: "United States"},
		{input: "USA", expected: "United States"},
		{input: "Spain", expected: "Spain"},
		{input: "russia", expected: "Russia"},
		{input: "RUSSIAN FEDERATION", expected: "Russia"},
		{input: "IRAN", expected: "Iran"},
		{input: "uk", expected: "United Kingdom"},
		{input: "UK", expected: "United Kingdom"},
		{input: "ENGLAND", expected: "United Kingdom"},
		{input: "United Kingdom", expected: "United Kingdom"},
		{input: "china", expected: "China"},
		{input: "north korea", expected: "North Korea"},
		{input: "South KOREA", expected: "South Korea"},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			require.Equal(t, tc.expected, country.Normalize(tc.input))
		})
	}
}

func TestNormalizeOverrides(t *testing.T) {
	require.Equal(t, "Czech Republic", country.Normalize("CZ"))
	require.Equal(t, "United Kingdom", country.Normalize("GB"))
	require.Equal(t, "Iran", country.Normalize("IR"))
	require.Equal(t, "North Korea", country.Normalize("KP"))
	require.Equal(t, "South Korea", country.Normalize("KR"))
	require.Equal(t, "Moldova", country.Normalize("MD"))
	require.Equal(t, "Saint Martin", country.Normalize("MF"))
	require.Equal(t, "Russia", country.Normalize("RU"))
	require.Equal(t, "Saint Martin", country.Normalize("SX"))
	require.Equal(t, "Syria", country.Normalize("SY"))
	require.Equal(t, "Turkey", country.Normalize("TR"))
	require.Equal(t, "Taiwan", country.Normalize("TW"))
	require.Equal(t, "United Kingdom", country.Normalize("UK"))
	require.Equal(t, "United States", country.Normalize("US"))
	require.Equal(t, "Venezuela", country.Normalize("VE"))
	require.Equal(t, "Vietnam", country.Normalize("VN"))
	require.Equal(t, "Virgin Islands", country.Normalize("VG"))
	require.Equal(t, "Virgin Islands", country.Normalize("VI"))
}
