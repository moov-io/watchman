// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStripLeadingPartyLabels(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"SHIPPER: TRANSPORTES REFRIGERADOS DEL SURESTE SA DE CV", "TRANSPORTES REFRIGERADOS DEL SURESTE SA DE CV"},
		{"BENEFICIARY: Sarıkaya Hububat Ltd. Şti.", "Sarıkaya Hububat Ltd. Şti."},
		{"ORDERING CUSTOMER: KALLIOSAARI TRADING OY", "KALLIOSAARI TRADING OY"},
		{"MESSRS. BALTIC TIMBERLINE OU", "BALTIC TIMBERLINE OU"},
		{"FIELD 59: BLACKWOOD CATTLE CO PTY LTD", "BLACKWOOD CATTLE CO PTY LTD"},
		{"Transportes Refrigerados del Sureste", "Transportes Refrigerados del Sureste"},
		{"", ""},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, StripLeadingPartyLabels(tc.in))
		})
	}
}
