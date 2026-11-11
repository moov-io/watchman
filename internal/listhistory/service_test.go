// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"context"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestService_NotConfigured(t *testing.T) {
	svc := NewService(log.NewTestLogger(), nil)
	require.False(t, svc.Enabled())

	_, err := svc.ListLatest(context.Background())
	require.ErrorIs(t, err, ErrNotConfigured)

	svc.StoreRefreshAsync(context.Background(), testStats("aaa", testTime(1), testPerson("1", "Alice")))
	require.NoError(t, svc.Wait(context.Background()))
}

func TestService_StoreAndDiff(t *testing.T) {
	repo := NewMemoryRepository()
	svc := NewService(log.NewTestLogger(), repo)
	ctx := context.Background()

	alice := testPerson("1", "Alice")
	bob := testPerson("2", "Bob")
	carol := testPerson("3", "Carol")

	svc.StoreRefreshAsync(ctx, testStats("aaa", testTime(1), alice, bob))
	require.NoError(t, svc.Wait(ctx))

	svc.StoreRefreshAsync(ctx, testStats("aaa", testTime(2), alice, bob))
	require.NoError(t, svc.Wait(ctx))

	refs, err := svc.ListRefs(ctx, "us_ofac")
	require.NoError(t, err)
	require.Len(t, refs, 1)
	require.Equal(t, "aaa", refs[0].Ref)
	require.Equal(t, 2, refs[0].AddedCount)

	alice2 := testPerson("1", "Alice Ann")
	svc.StoreRefreshAsync(ctx, testStats("bbb", testTime(3), alice2, carol))
	require.NoError(t, svc.Wait(ctx))

	refs, err = svc.ListRefs(ctx, "us_ofac")
	require.NoError(t, err)
	require.Len(t, refs, 2)
	require.Equal(t, "bbb", refs[0].Ref)
	require.Equal(t, 1, refs[0].AddedCount)
	require.Equal(t, 1, refs[0].RemovedCount)
	require.Equal(t, 1, refs[0].ChangedCount)
	require.Equal(t, 2, refs[0].EntityCount)

	listed, err := svc.ListEntities(ctx, "us_ofac", RefLatest, EntityFilter{Limit: 50})
	require.NoError(t, err)
	require.Equal(t, "bbb", listed.Ref.Ref)
	require.Len(t, listed.Entities, 2)

	got, err := svc.GetEntity(ctx, "us_ofac", "aaa", "2")
	require.NoError(t, err)
	require.Equal(t, "Bob", got.Name)

	diff, err := svc.Diff(ctx, "us_ofac", "", "bbb", "", DetailSummary, 100)
	require.NoError(t, err)
	require.Equal(t, "aaa", diff.From.Ref)
	require.Equal(t, "bbb", diff.To.Ref)
	require.Equal(t, 1, diff.Counts.Added)
	require.Equal(t, 1, diff.Counts.Removed)
	require.Equal(t, 1, diff.Counts.Changed)
	require.Equal(t, "3", diff.Added[0].SourceID)
	require.Equal(t, "2", diff.Removed[0].SourceID)
	require.Equal(t, "1", diff.Changed[0].SourceID)
	require.Nil(t, diff.Added[0].After)

	full, err := svc.Diff(ctx, "us_ofac", "aaa", "bbb", "", DetailFull, 100)
	require.NoError(t, err)
	require.NotNil(t, full.Added[0].After)
	require.Equal(t, "Carol", full.Added[0].After.Name)
	require.NotNil(t, full.Removed[0].Before)
	require.Equal(t, "Bob", full.Removed[0].Before.Name)
	require.NotNil(t, full.Changed[0].Before)
	require.NotNil(t, full.Changed[0].After)
	require.Equal(t, "Alice", full.Changed[0].Before.Name)
	require.Equal(t, "Alice Ann", full.Changed[0].After.Name)

	page, err := svc.Diff(ctx, "us_ofac", "aaa", "bbb", "", DetailSummary, 1)
	require.NoError(t, err)
	require.Equal(t, 1, page.Counts.Added)
	require.Len(t, page.Added, 1)
	require.Empty(t, page.Changed)
	require.Empty(t, page.Removed)
	require.Equal(t, "added:3", page.After)

	page2, err := svc.Diff(ctx, "us_ofac", "aaa", "bbb", page.After, DetailSummary, 10)
	require.NoError(t, err)
	require.Empty(t, page2.Added)
	require.Len(t, page2.Changed, 1)
	require.Len(t, page2.Removed, 1)
}

func TestService_KnownHashSkip(t *testing.T) {
	repo := NewMemoryRepository()
	svc := NewService(log.NewTestLogger(), repo)
	ctx := context.Background()

	alice := testPerson("1", "Alice")
	bob := testPerson("2", "Bob")
	svc.StoreRefreshAsync(ctx, testStats("aaa", testTime(1), alice, bob))
	require.NoError(t, svc.Wait(ctx))
	svc.StoreRefreshAsync(ctx, testStats("bbb", testTime(2), alice))
	require.NoError(t, svc.Wait(ctx))
	svc.StoreRefreshAsync(ctx, testStats("aaa", testTime(3), alice, bob))
	require.NoError(t, svc.Wait(ctx))

	refs, err := svc.ListRefs(ctx, "us_ofac")
	require.NoError(t, err)
	require.Len(t, refs, 2)
	require.Equal(t, "bbb", refs[0].Ref)
	require.Equal(t, 1, refs[0].EntityCount)
}

func TestService_StoreOrder(t *testing.T) {
	repo := NewMemoryRepository()
	svc := NewService(log.NewTestLogger(), repo)
	ctx := context.Background()

	svc.StoreRefreshAsync(ctx, testStats("aaa", testTime(1), testPerson("1", "Alice")))
	svc.StoreRefreshAsync(ctx, testStats("bbb", testTime(2), testPerson("1", "Alice"), testPerson("2", "Bob")))
	require.NoError(t, svc.Wait(ctx))

	refs, err := svc.ListRefs(ctx, "us_ofac")
	require.NoError(t, err)
	require.Len(t, refs, 2)
	require.Equal(t, "bbb", refs[0].Ref)
	require.Equal(t, "aaa", refs[1].Ref)
	require.Equal(t, 1, refs[0].AddedCount)
}

func TestPageChanges(t *testing.T) {
	added := []Change{{ChangeType: ChangeAdded, SourceID: "a"}}
	changed := []Change{{ChangeType: ChangeChanged, SourceID: "c"}}
	removed := []Change{{ChangeType: ChangeRemoved, SourceID: "r"}}

	outA, outR, outC, next := pageChanges(added, removed, changed, "", 2)
	require.Len(t, outA, 1)
	require.Len(t, outC, 1)
	require.Empty(t, outR)
	require.Equal(t, "changed:c", next)

	outA, outR, outC, next = pageChanges(added, removed, changed, next, 2)
	require.Empty(t, outA)
	require.Empty(t, outC)
	require.Len(t, outR, 1)
	require.Empty(t, next)
}
