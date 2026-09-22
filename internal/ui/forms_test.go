package ui

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestParseMinMatch(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in      string
		want    float64
		wantErr bool
	}{
		{in: "", want: 0},
		{in: "0.8", want: 0.8},
		{in: "1", want: 1},
		{in: "80", want: 0.8},
		{in: "80%", want: 0.8},
		{in: "1%", want: 0.01},
		{in: "100", want: 1},
		{in: "100%", want: 1},
		{in: "0", want: 0},
		{in: "101", wantErr: true},
		{in: "101%", wantErr: true},
		{in: "-1", wantErr: true},
		{in: "abc", wantErr: true},
		{in: "%", wantErr: true},
	}

	for _, tc := range cases {
		got, err := parseMinMatch(tc.in)
		if tc.wantErr {
			require.Error(t, err, tc.in)
			continue
		}
		require.NoError(t, err, tc.in)
		require.InDelta(t, tc.want, got, 0.0001, tc.in)
	}
}

func TestParseLimit(t *testing.T) {
	t.Parallel()

	require.Equal(t, 5, parseLimit(""))
	require.Equal(t, 5, parseLimit("0"))
	require.Equal(t, 5, parseLimit("nope"))
	require.Equal(t, 10, parseLimit("10"))
}

func TestParseGovernmentIDs(t *testing.T) {
	t.Parallel()

	require.Nil(t, parseGovernmentIDs(""))
	require.Nil(t, parseGovernmentIDs("\n  \n"))

	got := parseGovernmentIDs("AB123\npassport:XY9:US\n\ntax-id:12-345")
	require.Equal(t, []search.GovernmentID{
		{Identifier: "AB123"},
		{Type: "passport", Identifier: "XY9", Country: "US"},
		{Type: "tax-id", Identifier: "12-345"},
	}, got)
}

func TestLinesAndDates(t *testing.T) {
	t.Parallel()

	require.Nil(t, linesOf(" \n\t"))
	require.Equal(t, []string{"Ada", "Lovelace"}, linesOf(" Ada \n\nLovelace\n"))

	require.Nil(t, parseDate(""))
	require.Nil(t, parseDate("03/14/2020"))
	got := parseDate("2020-03-14")
	require.NotNil(t, got)
	require.True(t, got.Equal(time.Date(2020, 3, 14, 0, 0, 0, 0, time.UTC)))

	require.Equal(t, 0, parsePositiveInt(""))
	require.Equal(t, 0, parsePositiveInt("0"))
	require.Equal(t, 0, parsePositiveInt("-4"))
	require.Equal(t, 85000, parsePositiveInt(" 85000 "))
}

func TestFormatCountAndElapsed(t *testing.T) {
	t.Parallel()

	require.Equal(t, "0", formatCount(0))
	require.Equal(t, "12", formatCount(12))
	require.Equal(t, "1,234", formatCount(1234))
	require.Equal(t, "-1,234", formatCount(-1234))
	require.Equal(t, "1,000,000", formatCount(1000000))

	require.Equal(t, "500µs", formatElapsed(500*time.Microsecond))
	require.Equal(t, "2ms", formatElapsed(2*time.Millisecond+400*time.Microsecond))
	require.Equal(t, "1.5s", formatElapsed(1500*time.Millisecond))
}

func TestEntityTypeFromLabel(t *testing.T) {
	t.Parallel()

	require.Equal(t, "Person", entityTypeFromLabel("person"))
	require.Equal(t, "Vessel", entityTypeFromLabel(" VESSEL "))
	require.Equal(t, "", entityTypeFromLabel("widget"))
}

func TestDecodeDebug(t *testing.T) {
	t.Parallel()

	encoded := base64.StdEncoding.EncodeToString([]byte(" name 0.9 \n"))
	require.Equal(t, "name 0.9", decodeDebug(encoded))
	require.Equal(t, "", decodeDebug(""))
	require.Equal(t, "not-base64", decodeDebug("not-base64"))
}
