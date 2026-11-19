// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"strings"
	"testing"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestComputeDelta(t *testing.T) {
	a := testPerson("1", "Alice")
	b := testPerson("2", "Bob")
	c := testPerson("3", "Carol")

	hashA, err := hashEntity(a)
	require.NoError(t, err)
	hashB, err := hashEntity(b)
	require.NoError(t, err)

	t.Run("first snapshot all added", func(t *testing.T) {
		delta, err := computeDelta(nil, []search.Entity[search.Value]{a, b})
		require.NoError(t, err)
		require.Len(t, delta.Added, 2)
		require.Empty(t, delta.Removed)
		require.Empty(t, delta.Changed)
	})

	t.Run("add remove change", func(t *testing.T) {
		old := map[string]string{"1": hashA, "2": hashB}
		changedA := testPerson("1", "Alice Ann")
		delta, err := computeDelta(old, []search.Entity[search.Value]{changedA, c})
		require.NoError(t, err)
		require.Len(t, delta.Added, 1)
		require.Equal(t, "3", delta.Added[0].SourceID)
		require.Equal(t, []string{"2"}, delta.Removed)
		require.Len(t, delta.Changed, 1)
		require.Equal(t, "1", delta.Changed[0].SourceID)
		require.Equal(t, "Alice Ann", delta.Changed[0].Name)
	})

	t.Run("last wins duplicate source id", func(t *testing.T) {
		first := testPerson("1", "Old")
		last := testPerson("1", "New")
		delta, err := computeDelta(nil, []search.Entity[search.Value]{first, last})
		require.NoError(t, err)
		require.Len(t, delta.Added, 1)
		require.Equal(t, "New", delta.Added[0].Name)
	})

	t.Run("skip empty and oversized source id", func(t *testing.T) {
		empty := testPerson("", "Nobody")
		longID := testPerson(strings.Repeat("x", 256), "TooLong")
		delta, err := computeDelta(nil, []search.Entity[search.Value]{empty, longID, a})
		require.NoError(t, err)
		require.Len(t, delta.Added, 1)
		require.Equal(t, "1", delta.Added[0].SourceID)
	})

	t.Run("unchanged skipped", func(t *testing.T) {
		old := map[string]string{"1": hashA}
		delta, err := computeDelta(old, []search.Entity[search.Value]{a})
		require.NoError(t, err)
		require.Empty(t, delta.Added)
		require.Empty(t, delta.Removed)
		require.Empty(t, delta.Changed)
	})
}

func TestHashEntityStable(t *testing.T) {
	a := testPerson("1", "Alice")
	h1, err := hashEntity(a)
	require.NoError(t, err)
	h2, err := hashEntity(a)
	require.NoError(t, err)
	require.Equal(t, h1, h2)
	require.Len(t, h1, 64)

	b := testPerson("1", "Alice")
	b.Person.AltNames = []string{"Ally"}
	h3, err := hashEntity(b)
	require.NoError(t, err)
	require.NotEqual(t, h1, h3)
}

func TestLikeContains(t *testing.T) {
	got := likeContains(`100%_fun!`)
	require.Equal(t, `%100!%!_fun!!%`, got)
}

func TestTruncateName(t *testing.T) {
	require.Equal(t, "short", truncateName("short"))
	long := strings.Repeat("ä", maxNameRunes+10)
	got := truncateName(long)
	require.Equal(t, maxNameRunes, len([]rune(got)))
}

func TestValidRefAndListName(t *testing.T) {
	require.True(t, validListName("us_ofac"))
	require.True(t, validListName("eu-csl"))
	require.False(t, validListName("US_OFAC"))
	require.False(t, validListName(""))
	require.True(t, validRefToken(""))
	require.True(t, validRefToken("latest"))
	require.True(t, validRefToken("abc123"))
	require.False(t, validRefToken("has space"))
	require.False(t, validRefToken(strings.Repeat("a", 65)))
}
