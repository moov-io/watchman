package norm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwiftBIC(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in, want string
	}{
		{"HNRBBY2X", "HNRBBY2X"},
		{"hnrbby2x", "HNRBBY2X"},
		{" HNRB BY2X ", "HNRBBY2X"},
		{"HNRB-BY2X", "HNRBBY2X"},
		{"HNRBBY2XXXX", "HNRBBY2X"},
		{"SOMRRUM1KST", "SOMRRUM1KST"},
		{"CKLBCNBJ", "CKLBCNBJ"},
		{"GAZPRUMM", "GAZPRUMM"},
		{"RUDVRUMM", "RUDVRUMM"},
		{"not-a-bic", ""},
		{"HNRBBY", ""},
		{"12345678", ""},
		{"HN12BY2X", ""},
		{"PASSPORT", ""},
		{"SWIFTBIC", ""},
		{"SWIFT/BIC", ""},
		{"", ""},
	}
	for _, tc := range cases {
		require.Equal(t, tc.want, SwiftBIC(tc.in), tc.in)
		require.Equal(t, tc.want != "", IsSwiftBIC(tc.in), tc.in)
	}
}

func TestSwiftBIC8(t *testing.T) {
	t.Parallel()

	require.Equal(t, "HNRBBY2X", SwiftBIC8("HNRBBY2XXXX"))
	require.Equal(t, "SOMRRUM1", SwiftBIC8("SOMRRUM1KST"))
	require.Equal(t, "HNRBBY2X", SwiftBIC8("HNRBBY2X"))
	require.Empty(t, SwiftBIC8("nope"))
}

func TestEqualSwiftBIC(t *testing.T) {
	t.Parallel()

	require.True(t, EqualSwiftBIC("HNRBBY2X", "hnrbby2x"))
	require.True(t, EqualSwiftBIC("HNRBBY2X", "HNRBBY2XXXX"))
	require.True(t, EqualSwiftBIC("SOMRRUM1KST", "SOMRRUM1"))
	require.False(t, EqualSwiftBIC("HNRBBY2X", "RUDVRUMM"))
	require.False(t, EqualSwiftBIC("HNRBBY2X", ""))
	require.False(t, EqualSwiftBIC("", ""))
}

func TestParseSwiftBICs(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		in   []string
		want []string
	}{
		{
			name: "OFAC remark",
			in:   []string{"SWIFT/BIC HNRBBY2X; Registration ID 807000227 (Belarus) issued 16 Oct 2009; all offices worldwide."},
			want: []string{"HNRBBY2X"},
		},
		{
			name: "OFAC alt remark",
			in:   []string{"alt. SWIFT/BIC VVBKRU2Y"},
			want: []string{"VVBKRU2Y"},
		},
		{
			name: "OFAC 11-character",
			in:   []string{"SWIFT/BIC SOMRRUM1KST"},
			want: []string{"SOMRRUM1KST"},
		},
		{
			name: "CSL type-first",
			in:   []string{"Public Registration Number, 1027739108649; SWIFT/BIC, RUDVRUMM; Website, mspbank.ru"},
			want: []string{"RUDVRUMM"},
		},
		{
			name: "CSL value-first",
			in:   []string{"CKLBCNBJ, SWIFT/BIC"},
			want: []string{"CKLBCNBJ"},
		},
		{
			name: "bare column",
			in:   []string{"GAZPRUMM"},
			want: []string{"GAZPRUMM"},
		},
		{
			name: "comma-separated bare codes",
			in:   []string{"HNRBBY2X, RUDVRUMM"},
			want: []string{"HNRBBY2X", "RUDVRUMM"},
		},
		{
			name: "multiple",
			in:   []string{"SWIFT/BIC AFABAFKA", "alt. SWIFT/BIC VVBKRU2Y"},
			want: []string{"AFABAFKA", "VVBKRU2Y"},
		},
		{
			name: "dedupe 8 and XXX",
			in:   []string{"HNRBBY2X", "HNRBBY2XXXX"},
			want: []string{"HNRBBY2X"},
		},
		{
			name: "ignores other IDs",
			in:   []string{"Passport, U00242309, TR; Gender, Male"},
		},
		{
			name: "SWIFT/BIC CODE",
			in:   []string{"SWIFT/BIC CODE HNRBBY2X; all offices worldwide."},
			want: []string{"HNRBBY2X"},
		},
		{
			name: "does not treat Registration as a BIC",
			in:   []string{"SWIFT/BIC; Registration ID 807000227 (Belarus) issued 16 Oct 2009."},
		},
		{
			name: "period after BIC",
			in:   []string{"SWIFT/BIC SYLCLBBE."},
			want: []string{"SYLCLBBE"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, ParseSwiftBICs(tc.in...))
		})
	}
}
