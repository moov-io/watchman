// Copyright The Moov Authors
// SPDX-License-Identifier: Apache-2.0

package fincen_311

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClassifyEntityType(t *testing.T) {
	tests := []struct {
		name     string
		expected SMType
	}{
		{"ABLV Bank", SMTypeFinancialInstitution},
		{"FBME Bank Ltd.", SMTypeFinancialInstitution},
		{"Banca Privada d'Andorra", SMTypeFinancialInstitution},
		{"Bitzlato Limited", SMTypeFinancialInstitution},
		{"Liberty Reserve S.A.", SMTypeFinancialInstitution},
		{"Casa de Cambio Puebla", SMTypeFinancialInstitution},

		{"Islamic Republic of Iran", SMTypeJurisdiction},
		{"Democratic People's Republic of Korea", SMTypeJurisdiction},
		{"Burma", SMTypeJurisdiction},
		{"Nauru", SMTypeJurisdiction},

		{"Convertible Virtual Currency Mixing", SMTypeTransactionClass},
		{"Class of Transactions", SMTypeTransactionClass},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := classifyEntityType(tc.name)
			require.Equal(t, tc.expected, result, "entity: %s", tc.name)
		})
	}
}

func TestCleanText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  ABLV Bank  ", "ABLV Bank"},
		{"ABLV\n\nBank", "ABLV Bank"},
		{"ABLV\t\tBank", "ABLV Bank"},
		{"  Multiple   Spaces  ", "Multiple Spaces"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result := cleanText(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestParseHTMLContents(t *testing.T) {
	sampleHTML := `
<!DOCTYPE html>
<html>
<body>
<table>
	<thead>
		<tr>
			<th>Entity Name</th>
			<th>Finding</th>
			<th>NPRM</th>
			<th>Final Rule</th>
			<th>Rescinded</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>ABLV Bank</td>
			<td><a href="/sites/default/files/finding.pdf">02/13/2018</a></td>
			<td><a href="/sites/default/files/nprm.pdf">02/16/2018</a></td>
			<td><a href="/sites/default/files/final.pdf">06/25/2018</a></td>
			<td>---</td>
		</tr>
		<tr>
			<td>Islamic Republic of Iran</td>
			<td><a href="/iran/finding.pdf">11/21/2011</a></td>
			<td>---</td>
			<td><a href="/iran/final.pdf">01/01/2012</a></td>
			<td>---</td>
		</tr>
		<tr>
			<td>Bitzlato Limited</td>
			<td><a href="/bitzlato/finding.pdf">01/18/2023</a></td>
			<td>---</td>
			<td><a href="/bitzlato/final.pdf">05/01/2023</a></td>
			<td><a href="/bitzlato/rescind.pdf">12/01/2024</a></td>
		</tr>
	</tbody>
</table>
</body>
</html>`

	reader := &testReadCloser{Reader: strings.NewReader(sampleHTML)}
	data, err := parseHTMLContents(reader)
	require.NoError(t, err)
	require.NotNil(t, data)
	require.Len(t, data.SpecialMeasures, 3)
	require.NotEmpty(t, data.ListHash)

	// Check first entity (ABLV Bank)
	ablv := data.SpecialMeasures[0]
	require.Equal(t, "ABLV Bank", ablv.EntityName)
	require.Equal(t, SMTypeFinancialInstitution, ablv.EntityType)
	require.Contains(t, ablv.FindingURL, "finding.pdf")
	require.Equal(t, "02/13/2018", ablv.FindingDate)
	require.Contains(t, ablv.FinalRuleURL, "final.pdf")
	require.False(t, ablv.IsRescinded)

	// Check second entity (Iran - jurisdiction)
	iran := data.SpecialMeasures[1]
	require.Equal(t, "Islamic Republic of Iran", iran.EntityName)
	require.Equal(t, SMTypeJurisdiction, iran.EntityType)
	require.False(t, iran.IsRescinded)

	// Check third entity (Bitzlato - rescinded)
	bitzlato := data.SpecialMeasures[2]
	require.Equal(t, "Bitzlato Limited", bitzlato.EntityName)
	require.Equal(t, SMTypeFinancialInstitution, bitzlato.EntityType)
	require.True(t, bitzlato.IsRescinded)
	require.Contains(t, bitzlato.RescindedURL, "rescind.pdf")
}

func TestParseHTMLContents_NoTable(t *testing.T) {
	sampleHTML := `
<!DOCTYPE html>
<html>
<body>
<p>No table here</p>
</body>
</html>`

	reader := &testReadCloser{Reader: strings.NewReader(sampleHTML)}
	_, err := parseHTMLContents(reader)
	require.Error(t, err)
	require.Contains(t, err.Error(), "no table rows found")
}

func TestParseHTMLContents_EmptyTable(t *testing.T) {
	sampleHTML := `
<!DOCTYPE html>
<html>
<body>
<table>
	<thead>
		<tr>
			<th>Entity Name</th>
		</tr>
	</thead>
	<tbody>
	</tbody>
</table>
</body>
</html>`

	reader := &testReadCloser{Reader: strings.NewReader(sampleHTML)}
	_, err := parseHTMLContents(reader)
	require.Error(t, err)
}

// testReadCloser wraps a Reader to implement ReadCloser
type testReadCloser struct {
	io.Reader
}

func (t *testReadCloser) Close() error {
	return nil
}
