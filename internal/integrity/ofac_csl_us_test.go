package integrity

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/moov-io/watchman/internal/cslustest"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestIntegrity_OFAC_US_CSL(t *testing.T) {
	t.Run("28603", func(t *testing.T) {
		ofacEntity := ofactest.FindEntity(t, "28603")
		cslUSEntity := cslustest.FindEntity(t, "28603")

		// Basic Fields
		require.Equal(t, ofacEntity.Name, cslUSEntity.Name)
		require.Equal(t, ofacEntity.Type, cslUSEntity.Type)
		require.NotNil(t, ofacEntity.Business)
		require.NotNil(t, cslUSEntity.Business)

		// Business Fields
		require.Equal(t, ofacEntity.Business.Name, cslUSEntity.Business.Name)
		require.ElementsMatch(t, ofacEntity.Business.AltNames, cslUSEntity.Business.AltNames)
		require.Equal(t, ofacEntity.Business.Created, cslUSEntity.Business.Created)
		require.Equal(t, ofacEntity.Business.Dissolved, cslUSEntity.Business.Dissolved)
		require.ElementsMatch(t, ofacEntity.Business.GovernmentIDs, cslUSEntity.Business.GovernmentIDs)

		// Common Fields
		require.Equal(t, ofacEntity.Contact, cslUSEntity.Contact)
		require.ElementsMatch(t, ofacEntity.Addresses, cslUSEntity.Addresses)

		require.ElementsMatch(t, ofacEntity.CryptoAddresses, cslUSEntity.CryptoAddresses)

		// require.ElementsMatch(t, ofacEntity.Affiliations, cslUSEntity.Affiliations) // TODO(adam):
		// require.Equal(t, ofacEntity.SanctionsInfo, cslUSEntity.SanctionsInfo) // TODO(adam):
		require.ElementsMatch(t, ofacEntity.HistoricalInfo, cslUSEntity.HistoricalInfo)

		var buf bytes.Buffer
		score := search.DebugSimilarity(&buf, ofacEntity, cslUSEntity)
		require.InDelta(t, 1.00, score, 0.001)

		fmt.Println(buf.String())

	})
}
