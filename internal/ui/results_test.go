package ui

import (
	"strings"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"

	"github.com/stretchr/testify/require"
)

func TestEntityDetailTextIsSelectable(t *testing.T) {
	test.NewApp()

	born := time.Date(1962, 11, 23, 0, 0, 0, 0, time.UTC)
	root := entityDetails(search.SearchedEntity[search.Value]{
		Entity: search.Entity[search.Value]{
			Name:     "Nicolas MADURO MOROS",
			Type:     search.EntityPerson,
			Source:   search.SourceUSOFAC,
			SourceID: "22790",
			Person: &search.Person{
				Name:      "Nicolas MADURO MOROS",
				AltNames:  []string{"Nicolas MADURO"},
				Gender:    search.GenderMale,
				BirthDate: &born,
			},
			Addresses: []search.Address{{
				Line1: "Miraflores Palace", City: "Caracas", Country: "VE",
			}},
		},
		Match: 0.96,
	})

	var texts []string
	walkObjects(root, func(object fyne.CanvasObject) {
		label, ok := object.(*widget.Label)
		if !ok || label.Text == "" {
			return
		}
		texts = append(texts, label.Text)
		require.True(t, label.Selectable, label.Text)
	})
	joined := strings.Join(texts, "\n")
	require.Contains(t, joined, "Nicolas MADURO MOROS")
	require.Contains(t, joined, "22790")
	require.Contains(t, joined, "Miraflores Palace")
}

func TestScoreBreakdownRendersNameDebug(t *testing.T) {
	test.NewApp()

	root := entityDetails(search.SearchedEntity[search.Value]{
		Entity: search.Entity[search.Value]{
			Name:     "Mohamed SALEM",
			Type:     search.EntityPerson,
			Source:   search.SourceUSOFAC,
			SourceID: "12889",
		},
		Match: 0.69,
		Details: search.SimilarityScore{
			FinalScore: 0.69,
			Name: &search.NameDebug{
				Score:     0.91,
				QueryName: "primary",
				IndexName: "alt",
				Query: search.NameSide{
					Original: "Mohamed Salem",
					Prepared: "mohamed salem",
					Tokens:   []string{"mohamed", "salem"},
				},
				Index: search.NameSide{
					Original: "'Abd al-Rahman Ould Muhammad al-Husayn Ould Muhammad SALIM",
					Compared: "Mohamed SALEM",
					Prepared: "mohamed salem",
					Tokens:   []string{"mohamed", "salem"},
				},
				Pairs: []search.TokenPair{
					{Query: "mohamed", Index: "mohamed", Score: 1},
					{Query: "salem", Index: "salem", Score: 1},
				},
				Steps: []search.ScoreStep{
					{Stage: "token-pairs", Score: 0.91},
					{Stage: "close-enough", Score: 0.7735, Delta: -0.1365, Reason: "primary name term counts 2 vs 12"},
				},
			},
			Steps: []search.ScoreStep{
				{Stage: "base", Score: 0.7735},
				{Stage: "name-only", Score: 0.7348, Delta: -0.0387, Reason: "name-only query"},
			},
			Pieces: []search.ScorePiece{
				{PieceType: "name", Score: 0.7735, Weight: 35, Matched: true, Required: true, FieldsCompared: 1},
			},
		},
	})

	var texts []string
	walkObjects(root, func(object fyne.CanvasObject) {
		label, ok := object.(*widget.Label)
		if !ok || label.Text == "" {
			return
		}
		texts = append(texts, label.Text)
	})
	joined := strings.Join(texts, "\n")
	require.Contains(t, joined, "mohamed")
	require.Contains(t, joined, "token-pairs")
	require.Contains(t, joined, "name-only")
	require.Contains(t, joined, "Query: Mohamed Salem")
	require.Contains(t, joined, "compared: Mohamed SALEM")
}

func TestScoreBreakdownIdentifierOnlyOmitsDebugDump(t *testing.T) {
	test.NewApp()

	details := search.SimilarityScore{
		FinalScore: 1.0,
		Steps: []search.ScoreStep{
			{Stage: "base", Score: 1.0},
			{Stage: "exact-override", Score: 1.0, Reason: "unique identity key matched exactly"},
		},
		Pieces: []search.ScorePiece{
			{PieceType: "crypto-exact", Score: 1, Weight: 50, Matched: true, Exact: true, FieldsCompared: 1, UniqueIdentity: true},
		},
	}
	root := entityDetails(search.SearchedEntity[search.Value]{
		Entity: search.Entity[search.Value]{
			Name:     "SUEX OTC, S.R.O.",
			Type:     search.EntityBusiness,
			Source:   search.SourceUSOFAC,
			SourceID: "36216",
			CryptoAddresses: []search.CryptoAddress{
				{Currency: "XBT", Address: "12VrYZgS1nmf9KHHped24xBb1aLLRpV2cT"},
			},
		},
		Match:   1.0,
		Debug:   details.String(),
		Details: details,
	})

	var texts []string
	walkObjects(root, func(object fyne.CanvasObject) {
		label, ok := object.(*widget.Label)
		if !ok || label.Text == "" {
			return
		}
		texts = append(texts, label.Text)
	})
	joined := strings.Join(texts, "\n")
	require.Contains(t, joined, "exact-override")
	require.Contains(t, joined, "crypto-exact")
	require.NotContains(t, joined, "score steps")
	require.NotContains(t, joined, "pieces\n")
}

func TestEntityDetailIncludesSwiftBIC(t *testing.T) {
	test.NewApp()

	root := entityDetails(search.SearchedEntity[search.Value]{
		Entity: search.Entity[search.Value]{
			Name:      "ONERBANK ZAO",
			Type:      search.EntityBusiness,
			Source:    search.SourceUSOFAC,
			SourceID:  "12028",
			Business:  &search.Business{Name: "ONERBANK ZAO"},
			SwiftBICs: []string{"HNRBBY2X"},
		},
		Match: 1.0,
	})

	var texts []string
	walkObjects(root, func(object fyne.CanvasObject) {
		label, ok := object.(*widget.Label)
		if !ok || label.Text == "" {
			return
		}
		texts = append(texts, label.Text)
	})
	joined := strings.Join(texts, "\n")
	require.Contains(t, joined, "ONERBANK ZAO")
	require.Contains(t, joined, "12028")
	require.Contains(t, joined, "HNRBBY2X")
}

func walkObjects(object fyne.CanvasObject, fn func(fyne.CanvasObject)) {
	if object == nil {
		return
	}
	fn(object)
	switch typed := object.(type) {
	case *fyne.Container:
		for _, child := range typed.Objects {
			walkObjects(child, fn)
		}
	case *container.AppTabs:
		for _, item := range typed.Items {
			walkObjects(item.Content, fn)
		}
	case *container.Split:
		walkObjects(typed.Leading, fn)
		walkObjects(typed.Trailing, fn)
	case *container.Scroll:
		walkObjects(typed.Content, fn)
	case *widget.Form:
		for _, item := range typed.Items {
			walkObjects(item.Widget, fn)
		}
	case *widget.Card:
		walkObjects(typed.Content, fn)
	case *widget.Accordion:
		for _, item := range typed.Items {
			walkObjects(item.Detail, fn)
		}
	}
}
