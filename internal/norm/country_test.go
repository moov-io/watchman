package norm_test

import (
	"testing"

	"github.com/moov-io/watchman/internal/norm"

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
			require.Equal(t, tc.expected, norm.Country(tc.input))
		})
	}
}

func TestNormalizeOverrides(t *testing.T) {
	require.Equal(t, "Czech Republic", norm.Country("CZ"))
	require.Equal(t, "United Kingdom", norm.Country("GB"))
	require.Equal(t, "Iran", norm.Country("IR"))
	require.Equal(t, "North Korea", norm.Country("KP"))
	require.Equal(t, "South Korea", norm.Country("KR"))
	require.Equal(t, "Moldova", norm.Country("MD"))
	require.Equal(t, "Saint Martin", norm.Country("MF"))
	require.Equal(t, "Russia", norm.Country("RU"))
	require.Equal(t, "Saint Martin", norm.Country("SX"))
	require.Equal(t, "Syria", norm.Country("SY"))
	require.Equal(t, "Turkey", norm.Country("TR"))
	require.Equal(t, "Taiwan", norm.Country("TW"))
	require.Equal(t, "United Kingdom", norm.Country("UK"))
	require.Equal(t, "United States", norm.Country("US"))
	require.Equal(t, "Venezuela", norm.Country("VE"))
	require.Equal(t, "Vietnam", norm.Country("VN"))
	require.Equal(t, "Virgin Islands", norm.Country("VG"))
	require.Equal(t, "Virgin Islands", norm.Country("VI"))
}
