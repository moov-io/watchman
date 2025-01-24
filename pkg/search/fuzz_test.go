package search

import (
	"encoding/json"
	"testing"
)

func FuzzSimilarity(f *testing.F) {
	// Setup the fuzz corpus
	corpusAdd(f, Entity[Value]{})
	corpusAdd(f, Entity[Value]{
		Person: &Person{},
	})
	corpusAdd(f, Entity[Value]{
		Business: &Business{},
	})
	corpusAdd(f, Entity[Value]{
		Organization: &Organization{},
	})
	corpusAdd(f, Entity[Value]{
		Aircraft: &Aircraft{},
	})
	corpusAdd(f, Entity[Value]{
		Vessel: &Vessel{},
	})
	corpusAdd(f, Entity[Value]{
		SanctionsInfo: &SanctionsInfo{},
	})
	corpusAdd(f, Entity[Value]{
		Addresses:       []Address{{}},
		CryptoAddresses: []CryptoAddress{{}},
		Affiliations:    []Affiliation{{}},
		HistoricalInfo:  []HistoricalInfo{{}},
	})

	// Run the fuzz loop
	f.Fuzz(func(t *testing.T, queryData []byte, indexData []byte) {
		var query Entity[Value]
		err := json.Unmarshal(queryData, &query)
		if err != nil {
			return
		}

		var index Entity[Value]
		err = json.Unmarshal(indexData, &index)
		if err != nil {
			return
		}

		Similarity(query, index)
		Similarity(index, query)
	})
}

func corpusAdd(f *testing.F, entity Entity[Value]) {
	f.Helper()

	bs, err := json.Marshal(entity)
	if err != nil {
		f.Fatal(err)
	}
	f.Add(bs, bs)
}
