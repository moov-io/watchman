// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package opensanctions

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/watchman/pkg/download"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	testdataPath := filepath.Join("..", "..", "..", "test", "testdata", "opensanctions", "peps_senzing.json")

	fd, err := os.Open(testdataPath)
	require.NoError(t, err)
	defer fd.Close()

	files := make(download.Files)
	files["peps_senzing.json"] = io.NopCloser(fd)

	results, err := Read(files)
	require.NoError(t, err)
	require.NotNil(t, results)

	// Check that we have the expected number of entities
	require.Len(t, results.Entities, 3)

	// Check that the list hash was computed
	require.NotEmpty(t, results.ListHash)

	// Verify the first entity (person with first/last name)
	entity1 := results.Entities[0]
	require.Equal(t, search.SourceOpenSanctionsPEP, entity1.Source)
	require.Equal(t, "Q123456", entity1.SourceID)
	require.Equal(t, search.EntityPerson, entity1.Type)
	require.Equal(t, "John Smith", entity1.Name)
	require.NotNil(t, entity1.Person)

	// Verify the second entity (person with full name)
	entity2 := results.Entities[1]
	require.Equal(t, "Q789012", entity2.SourceID)
	require.Equal(t, search.EntityPerson, entity2.Type)
	require.Equal(t, "Maria Garcia Lopez", entity2.Name)

	// Verify the third entity (organization)
	entity3 := results.Entities[2]
	require.Equal(t, "Q345678", entity3.SourceID)
	require.Equal(t, search.EntityBusiness, entity3.Type)
	require.Equal(t, "Ministry of Finance", entity3.Name)
}

func TestRead_EmptyFiles(t *testing.T) {
	files := make(download.Files)

	results, err := Read(files)
	require.Error(t, err)
	require.Nil(t, results)
	require.Contains(t, err.Error(), "no files provided")
}

func TestRead_UnknownFile(t *testing.T) {
	files := make(download.Files)
	files["unknown.json"] = io.NopCloser(nil)

	results, err := Read(files)
	require.Error(t, err)
	require.Nil(t, results)
	require.Contains(t, err.Error(), "unknown file")
}
