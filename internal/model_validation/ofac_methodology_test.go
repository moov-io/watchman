package model_validation

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/tabwriter"

	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/internal/stringscore"

	"github.com/stretchr/testify/require"
)

// The official OFAC portal ( https://sanctionssearch.ofac.treas.gov/ ) returns very high scores for poor matches.
// Watchman performs many comparisons and breaks apart the query and indexed terms for a more precise comparison.
//
// In the OFAC portal searches for "Khamis", "Khamis Al", or "Khamis Ali" will return 100% matches when "Khamis",
// "Al", or "Ali" match. This leads to many "100%" matches that are not perfect matches.
//

func TestOFACMethodology_Name(t *testing.T) {
	ctx := context.Background()
	ofacClient := ofactest.NewClient()

	queryName := "Khamis Al"
	ofacResults, err := ofacClient.Search(ctx, ofactest.SearchParams{
		Name: queryName,
	})
	require.NoError(t, err)

	var ofacPerfectMatches []ofactest.SearchResult
	for _, result := range ofacResults {
		if result.NameScore == 100 {
			ofacPerfectMatches = append(ofacPerfectMatches, result)
		}
	}

	var buf bytes.Buffer
	buf.WriteString("## Comparison with OFAC portal\n")
	buf.WriteString(fmt.Sprintf("The [OFAC portal](https://sanctionssearch.ofac.treas.gov/) returned %d 100%% matches for %q\n", len(ofacPerfectMatches), queryName))

	// Scoring these names against our original query
	buf.WriteString(fmt.Sprintf("\nComparing OFAC Names against %q using Watchman's name similarity algorithm\n\n", queryName))

	var nameTable bytes.Buffer
	nameTable.WriteString("| Query Term | OFAC Name | Similarity Score |\n")
	nameTable.WriteString("|----|----|----|\n")
	w := tabwriter.NewWriter(&nameTable, 0, 0, 1, ' ', 0)
	for _, result := range ofacPerfectMatches {
		score := stringscore.BestPairsJaroWinkler(strings.Fields(queryName), strings.Fields(result.Name))
		fmt.Fprintf(&nameTable, "| %s |\t%s\t| %.3f |\n", queryName, strings.TrimSpace(result.Name), score)
	}
	w.Flush()
	buf.Write(nameTable.Bytes())

	// Update our methodology docs
	where := filepath.Join("..", "..", "docs", "methodology", "pages", "ofac-name-comparison.md")
	err = os.WriteFile(where, buf.Bytes(), 0600)
	require.NoError(t, err)

	fmt.Println(buf.String())
}
