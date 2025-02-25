package ui

import (
	"context"
	"encoding/base64"
	"fmt"
	"image/color"
	"path/filepath"
	"strings"

	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/ast"
	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func SearchContainer(ctx context.Context, env Environment) fyne.CanvasObject {
	// Create main container
	mainContainer := container.New(layout.NewBorderLayout(nil, nil, nil, nil))

	// Create containers for different sections
	formContainer := container.NewVBox()
	resultsContainer := ResultsContainer(env)

	// Create a warning label
	warningLabel := widget.NewLabel("")
	warningLabel.Hide()

	// Create the search form and button
	searchForm, submitButton := searchFormWithButton(ctx, env, warningLabel, resultsContainer)

	// Create a container for the bottom panel with warning and button
	bottomPanel := container.NewHBox(warningLabel, layout.NewSpacer(), submitButton)

	// Create a scroll container for the form
	formScrollContainer := container.NewScroll(searchForm)
	formScrollContainer.SetMinSize(fyne.NewSize(0, env.Height*0.5))

	// Add components to the form container
	formContainer.Add(formScrollContainer)
	formContainer.Add(bottomPanel)

	// Create a split layout
	splitContent := container.NewVSplit(
		formContainer,
		resultsContainer,
	)
	splitContent.SetOffset(0.6)

	mainContainer.Add(splitContent)
	return mainContainer
}

func searchFormWithButton(ctx context.Context, env Environment, warningLabel *widget.Label, results *fyne.Container) (*widget.Form, *widget.Button) {
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
	entityTypeSelect.PlaceHolder = "(required)"

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
		{Text: "EntityType *", Widget: entityTypeSelect},
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
		Items:    items,
		OnSubmit: nil, // We'll handle submission with our custom button
	}

	submitButton := widget.NewButton("Search", func() {
		// Hide any existing warning
		warningLabel.Hide()

		// Check if EntityType is selected
		if entityTypeSelect.Selected == "" {
			warningLabel.SetText("Problem: EntityType is required")
			warningLabel.Show()
			return
		}

		// Rest of the existing code
		populatedItems := collectPopulatedItems(items)
		env.Logger.Info().Logf("searching with %d fields", len(populatedItems))

		query := buildQueryEntity(populatedItems)
		searchOpts := search.SearchOpts{
			Limit: 5,
			Debug: true,
		}
		resp, err := env.Client.SearchByEntity(ctx, query, searchOpts)
		if err != nil {
			env.Logger.Error().LogErrorf("ERROR performing search: %v", err)
			warningLabel.SetText("Problem: " + err.Error())
			warningLabel.Show()
			return
		}

		UpdateResults(ctx, env, results, resp.Entities)
	})
	submitButton.Importance = widget.HighImportance

	return form, submitButton
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
		case "EntityType", "EntityType *":
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

// ResultsContainer provides a fixed-size container for displaying search results
// that won't jump around when populated
func ResultsContainer(env Environment) *fyne.Container {
	// Create a fixed height container to prevent layout jumps
	resultsContainer := container.NewVBox()
	resultsContainer.Resize(fyne.NewSize(env.Width, env.Height*0.4))

	// Add a placeholder that will be replaced with actual results
	placeholder := widget.NewLabel("Search results will appear here")
	placeholder.Alignment = fyne.TextAlignCenter

	resultsContainer.Add(placeholder)

	// Wrap in a scroll container with larger fixed height to prevent jumping
	scrollContainer := container.NewScroll(resultsContainer)
	scrollContainer.SetMinSize(fyne.NewSize(0, env.Height*0.4))

	return container.NewVBox(
		widget.NewLabelWithStyle("Results", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		scrollContainer,
	)
}

// UpdateResults populates the results container with entity data
func UpdateResults(ctx context.Context, env Environment, resultsContainer *fyne.Container, entities []search.SearchedEntity[search.Value]) {
	// Get the scroll container (which is the second element in the results VBox)
	scrollContainer := resultsContainer.Objects[2].(*container.Scroll)

	// Get the actual content container
	contentContainer := scrollContainer.Content.(*fyne.Container)

	// Clear existing content
	contentContainer.RemoveAll()

	if len(entities) == 0 {
		noResults := widget.NewLabelWithStyle("No matching results found", fyne.TextAlignCenter, fyne.TextStyle{})
		contentContainer.Add(noResults)
		contentContainer.Refresh()
		return
	}

	// Create a tab container to organize results
	tabs := container.NewAppTabs()

	// Create a tab for each result entity
	for _, entity := range entities {
		// Create the content for this entity tab with integrated debug panel
		entityContent := createEntityDetailsView(env, entity)

		// Add as a tab
		tabItem := container.NewTabItem(
			fmt.Sprintf("%s (%s)", entity.Entity.Name, formatEntityType(entity.Entity.Type)),
			entityContent,
		)
		tabs.Append(tabItem)
	}

	// Add a summary tab at the beginning that shows all results in a table (no debug panel)
	summaryTab := container.NewTabItem("Summary", createSummaryView(entities))
	tabs.Items = append([]*container.TabItem{summaryTab}, tabs.Items...)
	tabs.SetTabLocation(container.TabLocationTop)

	// Add the tabs directly to the content container
	contentContainer.Add(tabs)
	contentContainer.Refresh()
}

func createSummaryView(entities []search.SearchedEntity[search.Value]) fyne.CanvasObject {
	// Create a table for the summary view
	table := widget.NewTable(
		// Function to determine number of rows/columns
		func() (int, int) {
			return len(entities) + 1, 5 // +1 for header row, 5 columns
		},
		// Function to create cell content
		func() fyne.CanvasObject {
			return widget.NewLabel("Wide Content To Set Column Width")
		},
		// Function to update cell content
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.Alignment = fyne.TextAlignLeading
			label.Wrapping = fyne.TextTruncate

			// Header row
			if id.Row == 0 {
				label.TextStyle = fyne.TextStyle{Bold: true}
				switch id.Col {
				case 0:
					label.SetText("Name")
				case 1:
					label.SetText("Type")
				case 2:
					label.SetText("Match Score")
				case 3:
					label.SetText("Source")
				case 4:
					label.SetText("ID")
				}
				return
			}

			// Data rows
			entity := entities[id.Row-1]
			switch id.Col {
			case 0:
				label.SetText(entity.Entity.Name)
			case 1:
				label.SetText(string(entity.Entity.Type))
			case 2:
				label.SetText(fmt.Sprintf("%.1f%%", entity.Match*100))
			case 3:
				label.SetText(string(entity.Entity.Source))
			case 4:
				label.SetText(entity.Entity.SourceID)
			}
		},
	)

	// Set column widths with more appropriate distribution
	// With the debug panel taking up space, we need to adjust column widths
	table.SetColumnWidth(0, 400) // Name - slightly narrower
	table.SetColumnWidth(1, 90)  // Type - narrower
	table.SetColumnWidth(2, 90)  // Match Score
	table.SetColumnWidth(3, 90)  // Source
	table.SetColumnWidth(4, 90)  // ID

	return container.NewPadded(table)
}

// createEntityDetailsView builds a detailed view for a single entity with integrated debug panel
func createEntityDetailsView(env Environment, entity search.SearchedEntity[search.Value]) fyne.CanvasObject {
	// Create a vertical container for the entire view
	mainContainer := container.NewVBox()

	// Create a consistent header with match score
	headerContainer := container.NewVBox()

	// Add match score badge to header in a consistent position
	matchScoreBadge := createMatchScoreBadge(entity.Match)
	headerContainer.Add(matchScoreBadge)
	headerContainer.Add(widget.NewSeparator())

	// Add the header to the main container
	mainContainer.Add(headerContainer)

	// Create a horizontal split container for entity details and debug info
	detailsAndDebugSplit := container.NewHSplit(nil, nil)

	// Create a container for entity details
	content := container.NewVBox()

	// Basic information section
	basicInfoSection := createBasicInfoSection(entity.Entity)
	content.Add(basicInfoSection)

	// Create entity-specific sections based on type
	switch entity.Entity.Type {
	case search.EntityPerson:
		if entity.Entity.Person != nil {
			content.Add(createPersonSection(*entity.Entity.Person))
		}
	case search.EntityBusiness:
		if entity.Entity.Business != nil {
			content.Add(createBusinessSection(*entity.Entity.Business))
		}
	case search.EntityOrganization:
		if entity.Entity.Organization != nil {
			content.Add(createOrganizationSection(*entity.Entity.Organization))
		}
	case search.EntityAircraft:
		if entity.Entity.Aircraft != nil {
			content.Add(createAircraftSection(*entity.Entity.Aircraft))
		}
	case search.EntityVessel:
		if entity.Entity.Vessel != nil {
			content.Add(createVesselSection(*entity.Entity.Vessel))
		}
	}

	// Add contact information section if available
	if hasContactInfo(entity.Entity.Contact) {
		content.Add(createContactSection(entity.Entity.Contact))
	}

	// Add addresses section if available
	if len(entity.Entity.Addresses) > 0 {
		content.Add(createAddressesSection(entity.Entity.Addresses))
	}

	// Add crypto addresses section if available
	if len(entity.Entity.CryptoAddresses) > 0 {
		content.Add(createCryptoAddressesSection(entity.Entity.CryptoAddresses))
	}

	// Wrap entity details in a scroll container
	detailScrollContainer := container.NewScroll(content)
	detailScrollContainer.SetMinSize(fyne.NewSize(0, env.Height*0.4))

	// Create debug panel for this entity
	var debugContent fyne.CanvasObject

	// Debug panel header
	debugHeader := widget.NewLabelWithStyle("Debug Information", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	if entity.Debug != "" {
		// Decode base64 debug information
		debugData, err := base64.StdEncoding.DecodeString(entity.Debug)

		if err != nil {
			debugContent = container.NewVBox(
				debugHeader,
				widget.NewSeparator(),
				widget.NewLabel(fmt.Sprintf("Error decoding debug data: %v", err)),
			)
		} else {
			// Create a multiline text display for the debug info with standard text color
			debugText := widget.NewLabel(string(debugData) + strings.Repeat("\n", 5))
			debugText.Wrapping = fyne.TextWrapWord

			// Use a VBox container to preserve the header formatting
			debugPanel := container.NewVBox(debugHeader, debugText)

			// Wrap in scroll container for long debug output
			debugContent = container.NewScroll(debugPanel)
		}
	} else {
		debugContent = container.NewVBox(
			debugHeader,
			widget.NewSeparator(),
			widget.NewLabel("No debug information available for this entity"),
		)
	}

	// Set the components of the split container
	detailsAndDebugSplit.Leading = detailScrollContainer // Left side: entity details
	detailsAndDebugSplit.Trailing = debugContent         // Right side: debug panel

	// Set the split offset (35% for entity details, 65% for debug)
	detailsAndDebugSplit.SetOffset(0.35)

	// Add the split container to the main container
	mainContainer.Add(detailsAndDebugSplit)

	return mainContainer
}

// createBasicInfoSection creates a section for basic entity information
func createBasicInfoSection(entity search.Entity[search.Value]) fyne.CanvasObject {
	grid := container.New(layout.NewFormLayout())

	// Add basic entity information
	grid.Add(widget.NewLabelWithStyle("Name:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
	grid.Add(widget.NewLabel(entity.Name))

	grid.Add(widget.NewLabelWithStyle("Type:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
	grid.Add(widget.NewLabel(string(entity.Type)))

	grid.Add(widget.NewLabelWithStyle("Source:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
	grid.Add(widget.NewLabel(string(entity.Source)))

	if entity.SourceID != "" {
		grid.Add(widget.NewLabelWithStyle("Source ID:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(entity.SourceID))
	}

	section := NewExpandableSection("Basic Information", grid)
	section.Expanded = true
	return section
}

// createPersonSection creates a section for person-specific information
func createPersonSection(person search.Person) fyne.CanvasObject {
	grid := container.New(layout.NewFormLayout())

	if len(person.AltNames) > 0 {
		grid.Add(widget.NewLabelWithStyle("Alt Names:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(strings.Join(person.AltNames, ", ")))
	}

	if person.Gender != "" {
		grid.Add(widget.NewLabelWithStyle("Gender:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(string(person.Gender)))
	}

	if person.BirthDate != nil {
		grid.Add(widget.NewLabelWithStyle("Birth Date:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(person.BirthDate.Format("2006-01-02")))
	}

	if person.DeathDate != nil {
		grid.Add(widget.NewLabelWithStyle("Death Date:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(person.DeathDate.Format("2006-01-02")))
	}

	if len(person.Titles) > 0 {
		grid.Add(widget.NewLabelWithStyle("Titles:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(strings.Join(person.Titles, ", ")))
	}

	if len(person.GovernmentIDs) > 0 {
		grid.Add(widget.NewLabelWithStyle("Gov't IDs:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		idStrings := make([]string, len(person.GovernmentIDs))
		for i, id := range person.GovernmentIDs {
			idStrings[i] = fmt.Sprintf("%s: %s (%s)", id.Type, id.Identifier, id.Country)
		}
		grid.Add(widget.NewLabel(strings.Join(idStrings, "\n")))
	}

	section := NewExpandableSection("Person Details", grid)
	section.Expanded = true
	return section
}

// createBusinessSection creates a section for business-specific information
func createBusinessSection(business search.Business) fyne.CanvasObject {
	grid := container.New(layout.NewFormLayout())

	if len(business.AltNames) > 0 {
		grid.Add(widget.NewLabelWithStyle("Alt Names:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(strings.Join(business.AltNames, ", ")))
	}

	if business.Created != nil {
		grid.Add(widget.NewLabelWithStyle("Created:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(business.Created.Format("2006-01-02")))
	}

	if business.Dissolved != nil {
		grid.Add(widget.NewLabelWithStyle("Dissolved:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(business.Dissolved.Format("2006-01-02")))
	}

	if len(business.GovernmentIDs) > 0 {
		grid.Add(widget.NewLabelWithStyle("Gov't IDs:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		idStrings := make([]string, len(business.GovernmentIDs))
		for i, id := range business.GovernmentIDs {
			idStrings[i] = fmt.Sprintf("%s: %s (%s)", id.Type, id.Identifier, id.Country)
		}
		grid.Add(widget.NewLabel(strings.Join(idStrings, "\n")))
	}

	section := NewExpandableSection("Business Details", grid)
	section.Expanded = true
	return section
}

// createOrganizationSection creates a section for organization-specific information
func createOrganizationSection(org search.Organization) fyne.CanvasObject {
	grid := container.New(layout.NewFormLayout())

	if len(org.AltNames) > 0 {
		grid.Add(widget.NewLabelWithStyle("Alt Names:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(strings.Join(org.AltNames, ", ")))
	}

	if org.Created != nil {
		grid.Add(widget.NewLabelWithStyle("Created:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(org.Created.Format("2006-01-02")))
	}

	if org.Dissolved != nil {
		grid.Add(widget.NewLabelWithStyle("Dissolved:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(org.Dissolved.Format("2006-01-02")))
	}

	if len(org.GovernmentIDs) > 0 {
		grid.Add(widget.NewLabelWithStyle("Gov't IDs:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		idStrings := make([]string, len(org.GovernmentIDs))
		for i, id := range org.GovernmentIDs {
			idStrings[i] = fmt.Sprintf("%s: %s (%s)", id.Type, id.Identifier, id.Country)
		}
		grid.Add(widget.NewLabel(strings.Join(idStrings, "\n")))
	}

	section := NewExpandableSection("Organization Details", grid)
	section.Expanded = true
	return section
}

// createAircraftSection creates a section for aircraft-specific information
func createAircraftSection(aircraft search.Aircraft) fyne.CanvasObject {
	grid := container.New(layout.NewFormLayout())

	if len(aircraft.AltNames) > 0 {
		grid.Add(widget.NewLabelWithStyle("Alt Names:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(strings.Join(aircraft.AltNames, ", ")))
	}

	if aircraft.Type != "" {
		grid.Add(widget.NewLabelWithStyle("Type:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(string(aircraft.Type)))
	}

	if aircraft.Flag != "" {
		grid.Add(widget.NewLabelWithStyle("Flag:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(aircraft.Flag))
	}

	if aircraft.Built != nil {
		grid.Add(widget.NewLabelWithStyle("Built:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(aircraft.Built.Format("2006-01-02")))
	}

	if aircraft.ICAOCode != "" {
		grid.Add(widget.NewLabelWithStyle("ICAO Code:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(aircraft.ICAOCode))
	}

	if aircraft.Model != "" {
		grid.Add(widget.NewLabelWithStyle("Model:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(aircraft.Model))
	}

	if aircraft.SerialNumber != "" {
		grid.Add(widget.NewLabelWithStyle("Serial Number:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(aircraft.SerialNumber))
	}

	section := NewExpandableSection("Aircraft Details", grid)
	section.Expanded = true
	return section
}

// createVesselSection creates a section for vessel-specific information
func createVesselSection(vessel search.Vessel) fyne.CanvasObject {
	grid := container.New(layout.NewFormLayout())

	if len(vessel.AltNames) > 0 {
		grid.Add(widget.NewLabelWithStyle("Alt Names:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(strings.Join(vessel.AltNames, ", ")))
	}

	if vessel.IMONumber != "" {
		grid.Add(widget.NewLabelWithStyle("IMO Number:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(vessel.IMONumber))
	}

	if vessel.Type != "" {
		grid.Add(widget.NewLabelWithStyle("Type:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(string(vessel.Type)))
	}

	if vessel.Flag != "" {
		grid.Add(widget.NewLabelWithStyle("Flag:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(vessel.Flag))
	}

	if vessel.Built != nil {
		grid.Add(widget.NewLabelWithStyle("Built:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(vessel.Built.Format("2006-01-02")))
	}

	if vessel.Model != "" {
		grid.Add(widget.NewLabelWithStyle("Model:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(vessel.Model))
	}

	if vessel.Tonnage > 0 {
		grid.Add(widget.NewLabelWithStyle("Tonnage:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(fmt.Sprintf("%d", vessel.Tonnage)))
	}

	if vessel.MMSI != "" {
		grid.Add(widget.NewLabelWithStyle("MMSI:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(vessel.MMSI))
	}

	if vessel.CallSign != "" {
		grid.Add(widget.NewLabelWithStyle("Call Sign:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(vessel.CallSign))
	}

	if vessel.GrossRegisteredTonnage > 0 {
		grid.Add(widget.NewLabelWithStyle("GRT:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(fmt.Sprintf("%d", vessel.GrossRegisteredTonnage)))
	}

	if vessel.Owner != "" {
		grid.Add(widget.NewLabelWithStyle("Owner:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(vessel.Owner))
	}

	section := NewExpandableSection("Vessel Details", grid)
	section.Expanded = true
	return section
}

// createContactSection creates a section for contact information
func createContactSection(contact search.ContactInfo) fyne.CanvasObject {
	grid := container.New(layout.NewFormLayout())

	if len(contact.EmailAddresses) > 0 {
		grid.Add(widget.NewLabelWithStyle("Email:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(strings.Join(contact.EmailAddresses, "\n")))
	}

	if len(contact.PhoneNumbers) > 0 {
		grid.Add(widget.NewLabelWithStyle("Phone:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(strings.Join(contact.PhoneNumbers, "\n")))
	}

	if len(contact.FaxNumbers) > 0 {
		grid.Add(widget.NewLabelWithStyle("Fax:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(strings.Join(contact.FaxNumbers, "\n")))
	}

	if len(contact.Websites) > 0 {
		grid.Add(widget.NewLabelWithStyle("Websites:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(strings.Join(contact.Websites, "\n")))
	}

	section := NewExpandableSection("Contact Information", grid)
	section.Expanded = true
	return section
}

// createAddressesSection creates a section for addresses
func createAddressesSection(addresses []search.Address) fyne.CanvasObject {
	content := container.NewVBox()

	for i, addr := range addresses {
		addressCard := createAddressCard(i+1, addr)
		content.Add(addressCard)
	}

	section := NewExpandableSection("Addresses", content)
	section.Expanded = true
	return section
}

// createAddressCard creates a card for an individual address
func createAddressCard(index int, addr search.Address) fyne.CanvasObject {
	card := container.NewVBox()

	// Address title
	title := widget.NewLabelWithStyle(fmt.Sprintf("Address %d", index), fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	card.Add(title)

	// Address details
	grid := container.New(layout.NewFormLayout())

	if addr.Line1 != "" {
		grid.Add(widget.NewLabelWithStyle("Line 1:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(addr.Line1))
	}

	if addr.Line2 != "" {
		grid.Add(widget.NewLabelWithStyle("Line 2:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(addr.Line2))
	}

	if addr.City != "" {
		grid.Add(widget.NewLabelWithStyle("City:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(addr.City))
	}

	if addr.State != "" {
		grid.Add(widget.NewLabelWithStyle("State/Province:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(addr.State))
	}

	if addr.PostalCode != "" {
		grid.Add(widget.NewLabelWithStyle("Postal Code:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(addr.PostalCode))
	}

	if addr.Country != "" {
		grid.Add(widget.NewLabelWithStyle("Country:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(addr.Country))
	}

	if addr.Latitude != 0 || addr.Longitude != 0 {
		grid.Add(widget.NewLabelWithStyle("Coordinates:", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(fmt.Sprintf("%.6f, %.6f", addr.Latitude, addr.Longitude)))
	}

	card.Add(grid)

	// Add a divider after each address except the last one
	card.Add(widget.NewSeparator())

	return card
}

// createCryptoAddressesSection creates a section for crypto addresses
func createCryptoAddressesSection(cryptoAddresses []search.CryptoAddress) fyne.CanvasObject {
	grid := container.New(layout.NewFormLayout())

	for i, crypto := range cryptoAddresses {
		grid.Add(widget.NewLabelWithStyle(fmt.Sprintf("%s:", crypto.Currency), fyne.TextAlignTrailing, fyne.TextStyle{Bold: true}))
		grid.Add(widget.NewLabel(crypto.Address))

		// Add a bit of spacing every 2 addresses
		if i%2 == 1 && i < len(cryptoAddresses)-1 {
			grid.Add(widget.NewLabel(""))
			grid.Add(widget.NewLabel(""))
		}
	}

	section := NewExpandableSection("Crypto Addresses", grid)
	section.Expanded = true
	return section
}

// hasContactInfo checks if there is any contact information to display
func hasContactInfo(contact search.ContactInfo) bool {
	return len(contact.EmailAddresses) > 0 ||
		len(contact.PhoneNumbers) > 0 ||
		len(contact.FaxNumbers) > 0 ||
		len(contact.Websites) > 0
}

// Helper to format entity type for display
func formatEntityType(entityType search.EntityType) string {
	switch entityType {
	case search.EntityPerson:
		return "Person"
	case search.EntityBusiness:
		return "Business"
	case search.EntityOrganization:
		return "Organization"
	case search.EntityAircraft:
		return "Aircraft"
	case search.EntityVessel:
		return "Vessel"
	default:
		return string(entityType)
	}
}

func createMatchScoreBadge(match float64) fyne.CanvasObject {
	// Determine text color based on match percentage
	var textColor color.Color
	if match >= 0.85 {
		textColor = color.NRGBA{R: 220, G: 0, B: 0, A: 255} // Red for high match
	} else if match >= 0.7 {
		textColor = color.NRGBA{R: 255, G: 165, B: 0, A: 255} // Orange for medium match
	} else {
		textColor = color.NRGBA{R: 0, G: 128, B: 0, A: 255} // Green for low match
	}

	// Create the score text
	scoreText := fmt.Sprintf("Match Score: %.1f%%", match*100)
	scoreTextObj := canvas.NewText(scoreText, textColor)
	scoreTextObj.TextStyle = fyne.TextStyle{Bold: true}
	scoreTextObj.Alignment = fyne.TextAlignTrailing

	// Create a consistent container for the badge with fixed alignment
	badgeContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), scoreTextObj)

	return badgeContainer
}
