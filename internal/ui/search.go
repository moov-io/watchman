package ui

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/ast"
	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func SearchContainer(ctx context.Context, env Environment) fyne.CanvasObject {
	wrapper := fyne.NewContainerWithLayout(layout.NewVBoxLayout())

	warning := container.NewVBox()
	warning.Hide()

	results := container.NewVBox()
	results.Hide()

	form := searchForm(ctx, env, warning, results)

	// Create a scroll container with fixed size
	scrollContainer := container.NewScroll(form)
	scrollContainer.SetMinSize(fyne.NewSize(0, env.Height))

	wrapper.Add(scrollContainer)
	wrapper.Add(warning)
	wrapper.Add(results)

	return wrapper
}

func searchForm(ctx context.Context, env Environment, warning *fyne.Container, results *fyne.Container) *widget.Form {
	warning.Hide()

	// Create entity type section contents
	peopleContent := createPeopleContent()
	businessContent := createBusinessContent()
	organizationContent := createOrganizationContent()
	aircraftContent := createAircraftContent()
	vesselContent := createVesselContent()
	contactContent := createContactContent()
	addressContent := createAddressContent()

	// Create expandable sections
	peopleSection := NewExpandableSection("People", peopleContent)
	businessSection := NewExpandableSection("Business", businessContent)
	organizationSection := NewExpandableSection("Organization", organizationContent)
	aircraftSection := NewExpandableSection("Aircraft", aircraftContent)
	vesselSection := NewExpandableSection("Vessel", vesselContent)
	contactSection := NewExpandableSection("Contact Info", contactContent)
	addressSection := NewExpandableSection("Addresses", addressContent)

	// Create entity type select with callback to show/hide appropriate sections
	entityTypeSelect := newSelect("EntityType")
	entityTypeSelect.OnChanged = func(selected string) {
		// Hide all entity-specific sections first
		peopleSection.Expanded = false
		businessSection.Expanded = false
		organizationSection.Expanded = false
		aircraftSection.Expanded = false
		vesselSection.Expanded = false

		// Show relevant section based on selection
		switch selected {
		case string(search.EntityPerson):
			peopleSection.Expanded = true
		case string(search.EntityBusiness):
			businessSection.Expanded = true
		case string(search.EntityOrganization):
			organizationSection.Expanded = true
		case string(search.EntityAircraft):
			aircraftSection.Expanded = true
		case string(search.EntityVessel):
			vesselSection.Expanded = true
		}

		// Refresh all sections
		peopleSection.Refresh()
		businessSection.Refresh()
		organizationSection.Refresh()
		aircraftSection.Refresh()
		vesselSection.Refresh()
	}

	items := []*widget.FormItem{
		{Text: "Name", Widget: newInput()},
		{Text: "EntityType", Widget: entityTypeSelect},
		{
			Text:     "SourceList",
			HintText: "Original list the entity appeared on",
			Widget:   newSelect("SourceList"),
		},

		// Entity type sections
		{Text: "", Widget: peopleSection},
		{Text: "", Widget: businessSection},
		{Text: "", Widget: organizationSection},
		{Text: "", Widget: aircraftSection},
		{Text: "", Widget: vesselSection},

		// Common fields for all entity types
		{Text: "", Widget: contactSection},
		{Text: "", Widget: addressSection},

		// Other Fields section
		{Text: "Other Fields", Widget: widget.NewLabel(" ")},
		{Text: "CryptoAddresses", Widget: newMultilineInput(2)},
	}

	form := &widget.Form{
		Items: items,
		OnSubmit: func() {
			warning.Hide()
			results.Hide()

			populatedItems := collectPopulatedItems(items)
			env.Logger.Info().Logf("searching with %d fields", len(populatedItems))

			query := buildQueryEntity(populatedItems)
			searchOpts := search.SearchOpts{
				Limit: 5,
			}
			resp, err := env.Client.SearchByEntity(ctx, query, searchOpts)
			if err != nil {
				env.Logger.Error().LogErrorf("ERROR performing search: %v", err)
				showWarning(env, warning, err)
				return
			}

			err = showResults(env, results, resp.Entities)
			if err != nil {
				env.Logger.Error().LogErrorf("ERROR showing results: %v", err)
				showWarning(env, warning, err)
				return
			}
		},
	}

	return form
}

// Helper functions to create content containers for each section
func createPeopleContent() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Name:"),
		newInput(),
		widget.NewLabel("Alt Names:"),
		newMultilineInput(2),
		widget.NewLabel("Gender:"),
		newSelect("Gender"),
		widget.NewLabel("Birth Date:"),
		newInput(), // Date picker would be better here
		widget.NewLabel("Death Date:"),
		newInput(), // Date picker would be better here
		widget.NewLabel("Titles:"),
		newMultilineInput(2),
		widget.NewLabel("Government IDs:"),
		newMultilineInput(2),
	)
}

func createBusinessContent() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Name:"),
		newInput(),
		widget.NewLabel("Alt Names:"),
		newMultilineInput(2),
		widget.NewLabel("Created Date:"),
		newInput(), // Date picker would be better
		widget.NewLabel("Dissolved Date:"),
		newInput(), // Date picker would be better
		widget.NewLabel("Government IDs:"),
		newMultilineInput(2),
	)
}

func createOrganizationContent() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Name:"),
		newInput(),
		widget.NewLabel("Alt Names:"),
		newMultilineInput(2),
		widget.NewLabel("Created Date:"),
		newInput(), // Date picker would be better
		widget.NewLabel("Dissolved Date:"),
		newInput(), // Date picker would be better
		widget.NewLabel("Government IDs:"),
		newMultilineInput(2),
	)
}

func createAircraftContent() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Name:"),
		newInput(),
		widget.NewLabel("Alt Names:"),
		newMultilineInput(2),
		widget.NewLabel("Type:"),
		newSelect("AircraftType"),
		widget.NewLabel("Flag:"),
		newInput(), // Maybe a country select
		widget.NewLabel("Built Date:"),
		newInput(), // Date picker would be better
		widget.NewLabel("ICAO Code:"),
		newInput(),
		widget.NewLabel("Model:"),
		newInput(),
		widget.NewLabel("Serial Number:"),
		newInput(),
	)
}

func createVesselContent() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Name:"),
		newInput(),
		widget.NewLabel("Alt Names:"),
		newMultilineInput(2),
		widget.NewLabel("IMO Number:"),
		newInput(),
		widget.NewLabel("Type:"),
		newSelect("VesselType"),
		widget.NewLabel("Flag:"),
		newInput(), // Maybe a country select
		widget.NewLabel("Built Date:"),
		newInput(), // Date picker would be better
		widget.NewLabel("Model:"),
		newInput(),
		widget.NewLabel("Tonnage:"),
		newInput(),
		widget.NewLabel("MMSI:"),
		newInput(),
		widget.NewLabel("Call Sign:"),
		newInput(),
		widget.NewLabel("Gross Registered Tonnage:"),
		newInput(),
		widget.NewLabel("Owner:"),
		newInput(),
	)
}

func createContactContent() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Email Addresses:"),
		newMultilineInput(2),
		widget.NewLabel("Phone Numbers:"),
		newMultilineInput(2),
		widget.NewLabel("Fax Numbers:"),
		newMultilineInput(2),
		widget.NewLabel("Websites:"),
		newMultilineInput(2),
	)
}

func createAddressContent() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Line 1:"),
		newInput(),
		widget.NewLabel("Line 2:"),
		newInput(),
		widget.NewLabel("City:"),
		newInput(),
		widget.NewLabel("Postal Code:"),
		newInput(),
		widget.NewLabel("State:"),
		newInput(),
		widget.NewLabel("Country:"),
		newInput(), // Could be a country select
		widget.NewLabel("Latitude:"),
		newInput(),
		widget.NewLabel("Longitude:"),
		newInput(),
	)
}

func showResults(env Environment, results *fyne.Container, entities []search.SearchedEntity[search.Value]) error {
	results.RemoveAll()
	defer results.Show()

	header := widget.NewLabelWithStyle("Results", fyne.TextAlignLeading, fyne.TextStyle{
		Bold: true,
	})
	results.Add(header)

	var data = [][]string{
		{"top left", "top right"},
		{"bottom left", "bottom right"},
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
	modelsPath = filepath.Join("pkg", "search", "models.go") // needs to match what's in ../../package.go
)

func newSelect(modelName string) *widget.Select {
	values, err := ast.ExtractVariablesOfType(watchman.ModelsFilesystem, modelsPath, modelName)
	if err != nil {
		panic(fmt.Sprintf("reading %s values: %v", modelName, err)) //nolint:forbidigo
	}

	selectWidget := widget.NewSelect(values, func(_ string) {})

	return selectWidget
}

type item struct {
	name, value string
}

// Updated to collect values from expandable sections as well
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
		case *ExpandableSection:
			// For expandable sections, we need to examine their content
			if content, ok := w.Content.(*fyne.Container); ok {
				// Process each item in the container
				for j := 0; j < len(content.Objects); j += 2 {
					// Even indices should be labels, odd indices should be widgets
					if j+1 < len(content.Objects) {
						if label, okLabel := content.Objects[j].(*widget.Label); okLabel {
							fieldName := label.Text
							// Remove the colon if present
							if len(fieldName) > 0 && fieldName[len(fieldName)-1] == ':' {
								fieldName = fieldName[:len(fieldName)-1]
							}

							// Check the widget type
							switch inputWidget := content.Objects[j+1].(type) {
							case *widget.Entry:
								if inputWidget.Text != "" {
									out = append(out, item{name: fieldName, value: inputWidget.Text})
								}
							case *widget.Select:
								if inputWidget.Selected != "" {
									out = append(out, item{name: fieldName, value: inputWidget.Selected})
								}
							}
						}
					}
				}
			}
		case *widget.Label:
			// ignore
		}
	}
	return out
}

// Update buildQueryEntity to handle fields from the expandable sections
func buildQueryEntity(populatedItems []item) search.Entity[search.Value] {
	var out search.Entity[search.Value]

	// Initialize the nested structs
	out.Person = &search.Person{}
	out.Business = &search.Business{}
	out.Organization = &search.Organization{}
	out.Aircraft = &search.Aircraft{}
	out.Vessel = &search.Vessel{}

	for _, qry := range populatedItems {
		switch qry.name {
		// Basic properties
		case "Name":
			// Handle based on selected entity type
			if out.Type == search.EntityPerson {
				out.Person.Name = qry.value
			} else if out.Type == search.EntityBusiness {
				out.Business.Name = qry.value
			} else if out.Type == search.EntityOrganization {
				out.Organization.Name = qry.value
			} else if out.Type == search.EntityAircraft {
				out.Aircraft.Name = qry.value
			} else if out.Type == search.EntityVessel {
				out.Vessel.Name = qry.value
			}
			out.Name = qry.value
		case "EntityType":
			out.Type = search.EntityType(qry.value)
		case "SourceList":
			out.Source = search.SourceList(qry.value)

		// Person properties
		case "Gender":
			out.Person.Gender = search.Gender(qry.value)
		case "Birth Date":
			// Would need parsing logic for date
		case "Death Date":
			// Would need parsing logic for date
		case "Titles":
			// Would need parsing for multi-line

		// Business/Organization properties
		case "Created Date":
			// Would need parsing logic for date
		case "Dissolved Date":
			// Would need parsing logic for date

		// Aircraft properties
		case "ICAO Code":
			out.Aircraft.ICAOCode = qry.value
		case "Type":
			if out.Type == search.EntityAircraft {
				out.Aircraft.Type = search.AircraftType(qry.value)
			} else if out.Type == search.EntityVessel {
				out.Vessel.Type = search.VesselType(qry.value)
			}
		case "Serial Number":
			out.Aircraft.SerialNumber = qry.value

		// Vessel properties
		case "IMO Number":
			out.Vessel.IMONumber = qry.value
		case "MMSI":
			out.Vessel.MMSI = qry.value
		case "Call Sign":
			out.Vessel.CallSign = qry.value
		case "Tonnage":
			// Would need int parsing
		case "Gross Registered Tonnage":
			// Would need int parsing
		case "Owner":
			out.Vessel.Owner = qry.value

		// Contact info
		case "Email Addresses":
			// Would need parsing for multi-line
		case "Phone Numbers":
			// Would need parsing for multi-line
		case "Fax Numbers":
			// Would need parsing for multi-line
		case "Websites":
			// Would need parsing for multi-line

		// Address info - would need to handle as a slice
		case "Line 1":
			// Would need to create an address and append
		case "Line 2":
			// ...
		case "City":
			// ...

		// Crypto addresses
		case "CryptoAddresses":
			// Would need parsing logic for crypto addresses
		}
	}

	// Clean up any nil struct pointers if they weren't populated
	if out.Person.Name == "" {
		out.Person = nil
	}
	if out.Business.Name == "" {
		out.Business = nil
	}
	if out.Organization.Name == "" {
		out.Organization = nil
	}
	if out.Aircraft.Name == "" {
		out.Aircraft = nil
	}
	if out.Vessel.Name == "" {
		out.Vessel = nil
	}

	return out
}
