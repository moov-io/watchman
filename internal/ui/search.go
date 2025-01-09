package ui

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/moov-io/watchman/internal/ast"
	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func SearchContainer(env Environment) fyne.CanvasObject {
	wrapper := fyne.NewContainerWithLayout(layout.NewVBoxLayout())

	warning := container.NewVBox()
	warning.Hide()

	results := container.NewVBox()
	results.Hide()

	form := searchForm(env, warning, results)

	wrapper.Add(form)
	wrapper.Add(warning)
	wrapper.Add(results)

	return wrapper
}

func searchForm(env Environment, warning *fyne.Container, results *fyne.Container) *widget.Form {
	warning.Hide()

	blankSpace := widget.NewLabel(" ")

	items := []*widget.FormItem{
		{Text: searchName, Widget: newInput()},
		{Text: "EntityType", Widget: newSelect("EntityType")},
		{
			Text:     "SourceList",
			HintText: "Original list the entity appeared on",
			Widget:   newSelect("SourceList"),
		},

		// Person
		{Text: "People", Widget: blankSpace},

		// Business

		// Organization

		// Aircraft

		// Vessel

		// Other Fields
		{Text: "Other Fields", Widget: blankSpace},
		{Text: "CryptoAddresses", Widget: newMultilineInput(2)},
	}

	form := &widget.Form{
		Items: items,
		OnSubmit: func() {
			warning.Hide()
			results.Hide()

			populatedItems := collectPopulatedItems(items)
			fmt.Printf("searching with %d fields\n", len(populatedItems))

			ctx := context.Background() // TODO(adam):
			query := buildQueryEntity(populatedItems)
			resp, err := env.Client.SearchByEntity(ctx, query)
			if err != nil {
				fmt.Printf("ERROR performing search: %v\n", err)
				showWarning(env, warning, err)
				return
			}

			err = showResults(env, results, resp.Entities)
			if err != nil {
				fmt.Printf("ERROR showing results: %v\n", err)
				showWarning(env, warning, err)
				return
			}
		},
	}

	return form
}

func showResults(env Environment, results *fyne.Container, entities []search.SearchedEntity[search.Value]) error {
	results.RemoveAll()
	defer results.Show()

	header := widget.NewLabelWithStyle("Results", fyne.TextAlignLeading, fyne.TextStyle{
		Bold: true,
	})
	results.Add(header)

	var data = [][]string{
		[]string{"top left", "top right"},
		[]string{"bottom left", "bottom right"},
	}

	for _, row := range data {
		elm := container.NewHBox()

		for _, cell := range row {
			c := widget.NewLabel(cell)
			elm.Add(c)
		}

		results.Add(elm)
	}

	results.Refresh()

	return nil
}

func showWarning(env Environment, warning *fyne.Container, err error) {
	warning.RemoveAll()
	defer warning.Show()

	header := widget.NewLabelWithStyle("Problem", fyne.TextAlignLeading, fyne.TextStyle{
		Bold: true,
	})
	warning.Add(header)

	msg := widget.NewLabel(err.Error())
	warning.Add(msg)

	warning.Refresh()
}

func newInput() *widget.Entry {
	e := widget.NewEntry()
	e.Validator = func(input string) error {
		return nil
	}
	return e
}

func newMultilineInput(visibleRows int) *widget.Entry {
	e := widget.NewMultiLineEntry()
	e.SetMinRowsVisible(visibleRows)
	return e
}

var (
	modelsPath = filepath.Join("pkg", "search", "models.go") // TODO(adam): relative to Watchman root dir
)

func newSelect(modelName string) *widget.Select {
	values, err := ast.ExtractVariablesOfType(modelsPath, modelName)
	if err != nil {
		panic(fmt.Sprintf("reading %s values: %v", modelName, err)) //nolint:forbidigo
	}

	selectWidget := widget.NewSelect(values, func(_ string) {})

	return selectWidget
}

type item struct {
	name, value string
}

func collectPopulatedItems(formItems []*widget.FormItem) []item {
	var out []item
	for i := range formItems {
		switch w := formItems[i].Widget.(type) {
		case *widget.Entry:
			if w.Text != "" {
				out = append(out, item{name: formItems[i].Text, value: w.Text})
			}
		case *widget.Select:
			if w.Selected != "" {
				out = append(out, item{name: formItems[i].Text, value: w.Selected})
			}

		case *widget.Label:
			// ignore

		default:
			fmt.Printf("%T\n", w)
		}
	}
	return out
}
