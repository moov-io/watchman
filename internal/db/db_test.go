package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	ForEachDatabase(t, func(db DB) {
		err := db.Ping()
		require.NoError(t, err)
	})
}
