package model_validation

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
	"text/tabwriter"

	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

const (
	debugSimilarity = false
)

func TestOFAC_Comparsion_Person(t *testing.T) {
	// 29702,"LIFSHITS, Artem Mikhaylovich","individual","CYBER2] [ELECTION-EO13848",-0- ,-0- ,-0- ,-0- ,-0- ,-0- ,-0- ,"DOB 26 Dec 1992; nationality Russia; Email Address mycryptodeals@yandex.ru; alt. Email Address artemlv@hotmail.com; Gender Male; Digital Currency Address - XBT 12udabs2TkX7NXCSj6KpqXfakjE52ZPLhz; alt. Digital Currency Address - XBT 1DT3tenf14cxz9WFNxmYrXFbB6TFiVWA9U; Digital Currency Address - ETH 0x901bb9583b24d97e995513c6778dc6888ab6870e; alt. Digital Currency Address - ETH 0xa7e5d5a720f06526557c513402f2e6b5fa20b008; Secondary sanctions risk: Ukraine-/Russia-Related Sanctions Regulations, 31 CFR 589.201; Phone Number 79110354982; Digital Currency Address - LTC Leo3j36nn1JcsUQruytQhFUdCdCH5YHMR3; Digital Currency Address - DASH Xs3vzQmNvAxRa3Xo8XzQqUb3BMgb9EogF4; Passport 719032284."
	indexEntry := ofactest.FindEntity(t, "29702")

	ofacClient := ofactest.NewClient()

	t.Run("Name", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Name: indexEntry.Name,
			Type: search.EntityPerson,
		}
		var sanctionsListServiceParams ofactest.SearchParams
		sanctionsListServiceParams.Name = indexEntry.Name

		var buf bytes.Buffer
		watchmanSimilarity := search.DebugSimilarity(&buf, query, indexEntry)
		sanctionsServiceResult := topSanctionsListSearchResult(t, ofacClient, sanctionsListServiceParams)

		fmt.Printf("Name: %v\n", indexEntry.Name)
		printResults(buf, watchmanSimilarity, sanctionsServiceResult)
	})

	t.Run("Passport", func(t *testing.T) {
		if len(indexEntry.Person.GovernmentIDs) == 0 {
			t.Skip("no passport on index entry")
		}
		if indexEntry.Person.GovernmentIDs[0].Type != search.GovernmentIDPassport {
			t.Skipf("unexpected index government ID: %#v", indexEntry.Person.GovernmentIDs[0])
		}

		query := search.Entity[search.Value]{
			Name: indexEntry.Name,
			Type: search.EntityPerson,
			Person: &search.Person{
				GovernmentIDs: indexEntry.Person.GovernmentIDs,
			},
		}

		var sanctionsListServiceParams ofactest.SearchParams
		sanctionsListServiceParams.Name = indexEntry.Name
		sanctionsListServiceParams.IDNumber = indexEntry.Person.GovernmentIDs[0].Identifier

		var buf bytes.Buffer
		watchmanSimilarity := search.DebugSimilarity(&buf, query, indexEntry)
		sanctionsServiceResult := topSanctionsListSearchResult(t, ofacClient, sanctionsListServiceParams)

		fmt.Printf("Passport: %#v\n", query.Person.GovernmentIDs[0])
		printResults(buf, watchmanSimilarity, sanctionsServiceResult)

		// Wipe
		query.Person = &search.Person{}
		sanctionsListServiceParams.IDNumber = ""
	})

	t.Run("Crypto Address", func(t *testing.T) {
		// Add "XBT 1DT3tenf14cxz9WFNxmYrXFbB6TFiVWA9U"
		query := search.Entity[search.Value]{
			Name:            indexEntry.Name,
			Type:            search.EntityPerson,
			CryptoAddresses: indexEntry.CryptoAddresses[1:2],
		}
		var sanctionsListServiceParams ofactest.SearchParams
		sanctionsListServiceParams.Name = indexEntry.Name
		sanctionsListServiceParams.IDNumber = indexEntry.CryptoAddresses[1].Address

		var buf bytes.Buffer
		watchmanSimilarity := search.DebugSimilarity(&buf, query, indexEntry)
		sanctionsServiceResult := topSanctionsListSearchResult(t, ofacClient, sanctionsListServiceParams)

		fmt.Printf("Crypto Address: %#v\n", query.CryptoAddresses[0])
		printResults(buf, watchmanSimilarity, sanctionsServiceResult)

		// Wipe
		query.CryptoAddresses = nil
		sanctionsListServiceParams.IDNumber = ""
	})

	t.Run("Address", func(t *testing.T) {
		if len(indexEntry.Addresses) == 0 {
			t.Skip("no addresses on index entry")
		}

		query := search.Entity[search.Value]{
			Name:      indexEntry.Name,
			Type:      search.EntityPerson,
			Addresses: indexEntry.Addresses[0:1],
		}

		var sanctionsListServiceParams ofactest.SearchParams
		sanctionsListServiceParams.Name = indexEntry.Name
		sanctionsListServiceParams.Address = indexEntry.Addresses[0].Line1
		sanctionsListServiceParams.City = indexEntry.Addresses[0].City
		sanctionsListServiceParams.StateProvince = indexEntry.Addresses[0].State

		var buf bytes.Buffer
		watchmanSimilarity := search.DebugSimilarity(&buf, query, indexEntry)
		sanctionsServiceResult := topSanctionsListSearchResult(t, ofacClient, sanctionsListServiceParams)

		fmt.Printf("Address: %v\n", indexEntry.Addresses[0].Format())
		printResults(buf, watchmanSimilarity, sanctionsServiceResult)
	})
}

func topSanctionsListSearchResult(tb testing.TB, ofacClient ofactest.Client, params ofactest.SearchParams) *ofactest.SearchResult {
	tb.Helper()

	results, err := ofacClient.Search(context.Background(), params)
	require.NoError(tb, err)

	if len(results) > 0 {
		return &results[0]
	}
	return nil
}

func printResults(buf bytes.Buffer, watchmanSimilarity float64, sanctionsServiceResult *ofactest.SearchResult) {
	if debugSimilarity {
		fmt.Println(buf.String())
	}
	fmt.Println("")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	fmt.Fprintf(w, "Watchman\t%.2f%%\n", watchmanSimilarity*100.0)
	if sanctionsServiceResult != nil {
		fmt.Fprintf(w, "OFAC Sanctions List Service\t%.2f%%\n", float64(sanctionsServiceResult.NameScore))
	} else {
		fmt.Fprintln(w, "OFAC Sanctions List Service\tno results")
	}
}
