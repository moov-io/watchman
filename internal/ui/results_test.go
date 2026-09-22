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
