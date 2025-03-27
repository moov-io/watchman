package ui

import (
	"cmp"
	"context"
	"fmt"
	"strings"

	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/ast"
	"github.com/moov-io/watchman/internal/ui/arch"
	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// SearchContainer creates a search interface with entity selection, form fields, and results display.
func SearchContainer(ctx context.Context, env Environment) fyne.CanvasObject {
	// Entity Type Section
	entityTypeLabel := widget.NewLabelWithStyle("Entity Type", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	entityTypes := []string{"Person", "Business", "Organization", "Aircraft", "Vessel"}
	entityTypeRadio := widget.NewRadioGroup(entityTypes, nil)
	entityTypeRadio.Horizontal = true

	// Entity-Specific Content
	personContent := createPersonContent()
	businessContent := createBusinessContent()
	organizationContent := createOrganizationContent()
	aircraftContent := createAircraftContent()
	vesselContent := createVesselContent()
	entityFieldsContainer := container.NewVBox()

	// Update entity-specific fields based on selected entity type
	entityTypeRadio.OnChanged = func(selected string) {
		entityFieldsContainer.RemoveAll()
		switch selected {
		case "Person":
			entityFieldsContainer.Add(personContent)
		case "Business":
			entityFieldsContainer.Add(businessContent)
		case "Organization":
			entityFieldsContainer.Add(organizationContent)
		case "Aircraft":
			entityFieldsContainer.Add(aircraftContent)
		case "Vessel":
			entityFieldsContainer.Add(vesselContent)
		}
		entityFieldsContainer.Refresh()
	}

	// Grab any prefilled values from the environment
	prefilledValues := arch.PrefillValues()

	// Common Search Fields
	commonFieldsLabel := widget.NewLabelWithStyle("Common Fields", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	warningLabel := widget.NewLabel("")
	warningLabel.TextStyle = fyne.TextStyle{Bold: true}
	warningLabel.Hide()
	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = cmp.Or(prefilledValues.Get("name"), "Enter name to search")
	sourceSelect := newSelect("SourceList") // Assume newSelect is a custom function returning a widget.Select
	sourceSelect.PlaceHolder = "Select source list (optional)"
	searchForm := widget.NewForm(
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("Source List", sourceSelect),
	)

	// Entity-Specific Fields Section
	entitySpecificLabel := widget.NewLabelWithStyle("Entity-Specific Fields", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	// Search Button
	searchButton := widget.NewButtonWithIcon("Search", theme.SearchIcon(), nil)
	searchButton.Importance = widget.HighImportance
	buttonContainer := container.NewHBox(warningLabel, layout.NewSpacer(), searchButton)

	// Top Content Layout
	topContent := container.NewVBox(
		entityTypeLabel,
		entityTypeRadio,
		widget.NewSeparator(),
		commonFieldsLabel,
		searchForm,
		widget.NewSeparator(),
		entitySpecificLabel,
		entityFieldsContainer,
		buttonContainer,
	)

	// Results Section
	resultsHeader := widget.NewLabelWithStyle("Results", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	resultsSeparator := widget.NewSeparator()
	resultsContent := container.NewVBox(widget.NewLabel("Search results will appear here"))

	// Main Layout
	center := container.NewVBox(resultsHeader, resultsSeparator, resultsContent)
	mainContent := container.NewVBox(topContent, center)

	// Search Button Callback
	searchButton.OnTapped = func() {
		warningLabel.Hide()

		// Validate required fields
		if entityTypeRadio.Selected == "" {
			warningLabel.SetText("Please select an entity type")
			warningLabel.Show()
			return
		}
		if nameEntry.Text == "" {
			warningLabel.SetText("Please enter a name to search")
			warningLabel.Show()
			return
		}

		// Clear and update results
		resultsContent.RemoveAll()
		resultsContent.Add(widget.NewLabel("Searching... Please wait"))
		resultsContent.Refresh()

		// Build search query (assume buildQueryEntity is defined elsewhere)
		query := buildQueryEntity(
			entityTypeRadio.Selected,
			nameEntry.Text,
			sourceSelect.Selected,
			personContent,
			businessContent,
			organizationContent,
			aircraftContent,
			vesselContent,
		)
		searchOpts := search.SearchOpts{Limit: 5, Debug: true}

		// Push a notification to the system
		sendSearchNotification(entityTypeRadio.Selected, query)

		// Perform search in a goroutine to avoid blocking the UI
		go func() {
			resp, err := env.Client.SearchByEntity(ctx, query, searchOpts)
			fyne.CurrentApp().Driver().CanvasForObject(mainContent).Content().Refresh()
			if err != nil {
				env.Logger.Error().LogErrorf("ERROR performing search: %v", err)
				resultsContent.RemoveAll()
				resultsContent.Add(widget.NewLabel(fmt.Sprintf("Error: %v", err)))
				resultsContent.Refresh()
				return
			}
			// Update results (assume updateResultsDisplay is defined elsewhere)
			updateResultsDisplay(env, resultsContent, resp.Entities)
		}()
	}

	return mainContent
}

func sendSearchNotification(entityType string, query search.Entity[search.Value]) {
	device := fyne.CurrentDevice()
	if device.IsBrowser() || device.IsMobile() {
		return
	}

	msg := fyne.NewNotification("Watchman Search", fmt.Sprintf("%s search for %s", entityType, query.Name))

	fyne.CurrentApp().SendNotification(msg)
}

// buildQueryEntity creates a search query from form data
func buildQueryEntity(
	entityType, name, sourceList string,
	personContent fyne.CanvasObject,
	businessContent fyne.CanvasObject,
	organizationContent fyne.CanvasObject,
	aircraftContent fyne.CanvasObject,
	vesselContent fyne.CanvasObject,
) search.Entity[search.Value] {
	var entity search.Entity[search.Value]

	// Set common fields
	entity.Name = name

	if sourceList != "" {
		entity.Source = search.SourceList(sourceList)
	}

	// Set entity type and specific fields based on selection
	switch entityType {
	case "Person":
		entity.Type = search.EntityPerson
		entity.Person = extractPersonFields(personContent)
		entity.Person.Name = name
	case "Business":
		entity.Type = search.EntityBusiness
		entity.Business = extractBusinessFields(businessContent)
		entity.Business.Name = name
	case "Organization":
		entity.Type = search.EntityOrganization
		entity.Organization = extractOrganizationFields(organizationContent)
		entity.Organization.Name = name
	case "Aircraft":
		entity.Type = search.EntityAircraft
		entity.Aircraft = extractAircraftFields(aircraftContent)
		entity.Aircraft.Name = name
	case "Vessel":
		entity.Type = search.EntityVessel
		entity.Vessel = extractVesselFields(vesselContent)
		entity.Vessel.Name = name
	}

	return entity
}

func extractPersonFields(content fyne.CanvasObject) *search.Person {
	person := &search.Person{}

	if content == nil {
		return person
	}

	container, isContainer := content.(*fyne.Container)
	if !isContainer || len(container.Objects) < 2 {
		return person
	}

	form, isForm := container.Objects[1].(*widget.Form)
	if !isForm {
		return person
	}

	// Extract alt names
	if len(form.Items) > 0 {
		if entry, isEntry := form.Items[0].Widget.(*widget.Entry); entry.MultiLine && isEntry {
			if entry.Text != "" {
				person.AltNames = strings.Split(entry.Text, "\n")
			}
		}
	}

	// Extract gender
	if len(form.Items) > 1 {
		if select_, isSelect := form.Items[1].Widget.(*widget.Select); isSelect {
			if select_.Selected != "" {
				person.Gender = search.Gender(strings.ToLower(select_.Selected))
			}
		}
	}

	return person
}

func extractBusinessFields(content fyne.CanvasObject) *search.Business {
	business := &search.Business{}

	if content == nil {
		return business
	}

	container, isContainer := content.(*fyne.Container)
	if !isContainer || len(container.Objects) < 2 {
		return business
	}

	form, isForm := container.Objects[1].(*widget.Form)
	if !isForm {
		return business
	}

	// Extract alt names
	if len(form.Items) > 0 {
		if entry, isEntry := form.Items[0].Widget.(*widget.Entry); entry.MultiLine && isEntry {
			if entry.Text != "" {
				business.AltNames = strings.Split(entry.Text, "\n")
			}
		}
	}

	return business
}

func extractOrganizationFields(content fyne.CanvasObject) *search.Organization {
	org := &search.Organization{}

	if content == nil {
		return org
	}

	container, isContainer := content.(*fyne.Container)
	if !isContainer || len(container.Objects) < 2 {
		return org
	}

	form, isForm := container.Objects[1].(*widget.Form)
	if !isForm {
		return org
	}

	// Extract alt names
	if len(form.Items) > 0 {
		if entry, isEntry := form.Items[0].Widget.(*widget.Entry); entry.MultiLine && isEntry {
			if entry.Text != "" {
				org.AltNames = strings.Split(entry.Text, "\n")
			}
		}
	}

	return org
}

func extractAircraftFields(content fyne.CanvasObject) *search.Aircraft {
	aircraft := &search.Aircraft{}

	if content == nil {
		return aircraft
	}

	container, isContainer := content.(*fyne.Container)
	if !isContainer || len(container.Objects) < 2 {
		return aircraft
	}

	form, isForm := container.Objects[1].(*widget.Form)
	if !isForm {
		return aircraft
	}

	// Extract alt names
	if len(form.Items) > 0 {
		if entry, isEntry := form.Items[0].Widget.(*widget.Entry); entry.MultiLine && isEntry {
			if entry.Text != "" {
				aircraft.AltNames = strings.Split(entry.Text, "\n")
			}
		}
	}

	// Extract type
	if len(form.Items) > 1 {
		if select_, isSelect := form.Items[1].Widget.(*widget.Select); isSelect {
			if select_.Selected != "" {
				aircraft.Type = search.AircraftType(strings.ToLower(select_.Selected))
			}
		}
	}

	return aircraft
}

func extractVesselFields(content fyne.CanvasObject) *search.Vessel {
	vessel := &search.Vessel{}

	if content == nil {
		return vessel
	}

	container, isContainer := content.(*fyne.Container)
	if !isContainer || len(container.Objects) < 2 {
		return vessel
	}

	form, isForm := container.Objects[1].(*widget.Form)
	if !isForm {
		return vessel
	}

	// Extract alt names
	if len(form.Items) > 0 {
		if entry, isEntry := form.Items[0].Widget.(*widget.Entry); entry.MultiLine && isEntry {
			if entry.Text != "" {
				vessel.AltNames = strings.Split(entry.Text, "\n")
			}
		}
	}

	// Extract IMO number
	if len(form.Items) > 1 {
		if entry, isEntry := form.Items[1].Widget.(*widget.Entry); isEntry {
			vessel.IMONumber = entry.Text
		}
	}

	// Extract type
	if len(form.Items) > 2 {
		if select_, isSelect := form.Items[2].Widget.(*widget.Select); isSelect {
			if select_.Selected != "" {
				vessel.Type = search.VesselType(strings.ToLower(select_.Selected))
			}
		}
	}

	return vessel
}

func createPersonContent() fyne.CanvasObject {
	form := widget.NewForm()

	altNamesEntry := widget.NewMultiLineEntry()
	altNamesEntry.PlaceHolder = "Alternative names, one per line"
	altNamesEntry.SetMinRowsVisible(2)
	form.Append("Alt Names", altNamesEntry)

	genderSelect := newSelect("Gender")
	genderSelect.PlaceHolder = "Select gender"
	form.Append("Gender", genderSelect)

	birthDateEntry := widget.NewEntry()
	birthDateEntry.PlaceHolder = "YYYY-MM-DD"
	form.Append("Birth Date", birthDateEntry)

	// Add an ID section
	idEntry := widget.NewMultiLineEntry()
	idEntry.PlaceHolder = "ID type:number:country, one per line"
	idEntry.SetMinRowsVisible(2)
	form.Append("Government IDs", idEntry)

	return container.NewVBox(
		widget.NewLabelWithStyle("Person Fields", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		form,
	)
}

func createBusinessContent() fyne.CanvasObject {
	form := widget.NewForm()

	altNamesEntry := widget.NewMultiLineEntry()
	altNamesEntry.PlaceHolder = "Alternative names, one per line"
	altNamesEntry.SetMinRowsVisible(2)
	form.Append("Alt Names", altNamesEntry)

	createdDateEntry := widget.NewEntry()
	createdDateEntry.PlaceHolder = "YYYY-MM-DD"
	form.Append("Created Date", createdDateEntry)

	dissolvedDateEntry := widget.NewEntry()
	dissolvedDateEntry.PlaceHolder = "YYYY-MM-DD"
	form.Append("Dissolved Date", dissolvedDateEntry)

	// Add an ID section
	idEntry := widget.NewMultiLineEntry()
	idEntry.PlaceHolder = "ID type:number:country, one per line"
	idEntry.SetMinRowsVisible(2)
	form.Append("Government IDs", idEntry)

	return container.NewVBox(
		widget.NewLabelWithStyle("Business Fields", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		form,
	)
}

func createOrganizationContent() fyne.CanvasObject {
	form := widget.NewForm()

	altNamesEntry := widget.NewMultiLineEntry()
	altNamesEntry.PlaceHolder = "Alternative names, one per line"
	altNamesEntry.SetMinRowsVisible(2)
	form.Append("Alt Names", altNamesEntry)

	createdDateEntry := widget.NewEntry()
	createdDateEntry.PlaceHolder = "YYYY-MM-DD"
	form.Append("Created Date", createdDateEntry)

	dissolvedDateEntry := widget.NewEntry()
	dissolvedDateEntry.PlaceHolder = "YYYY-MM-DD"
	form.Append("Dissolved Date", dissolvedDateEntry)

	// Add an ID section
	idEntry := widget.NewMultiLineEntry()
	idEntry.PlaceHolder = "ID type:number:country, one per line"
	idEntry.SetMinRowsVisible(2)
	form.Append("Government IDs", idEntry)

	return container.NewVBox(
		widget.NewLabelWithStyle("Organization Fields", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		form,
	)
}

func createAircraftContent() fyne.CanvasObject {
	form := widget.NewForm()

	altNamesEntry := widget.NewMultiLineEntry()
	altNamesEntry.PlaceHolder = "Alternative names, one per line"
	altNamesEntry.SetMinRowsVisible(2)
	form.Append("Alt Names", altNamesEntry)

	typeSelect := newSelect("AircraftType")
	typeSelect.PlaceHolder = "Select aircraft type"
	form.Append("Type", typeSelect)

	flagEntry := widget.NewEntry()
	flagEntry.PlaceHolder = "Country flag"
	form.Append("Flag", flagEntry)

	builtDateEntry := widget.NewEntry()
	builtDateEntry.PlaceHolder = "YYYY-MM-DD"
	form.Append("Built Date", builtDateEntry)

	icaoCodeEntry := widget.NewEntry()
	icaoCodeEntry.PlaceHolder = "ICAO code"
	form.Append("ICAO Code", icaoCodeEntry)

	modelEntry := widget.NewEntry()
	modelEntry.PlaceHolder = "Aircraft model"
	form.Append("Model", modelEntry)

	serialNumberEntry := widget.NewEntry()
	serialNumberEntry.PlaceHolder = "Serial number"
	form.Append("Serial Number", serialNumberEntry)

	return container.NewVBox(
		widget.NewLabelWithStyle("Aircraft Fields", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		form,
	)
}

func createVesselContent() fyne.CanvasObject {
	form := widget.NewForm()

	altNamesEntry := widget.NewMultiLineEntry()
	altNamesEntry.PlaceHolder = "Alternative names, one per line"
	altNamesEntry.SetMinRowsVisible(2)
	form.Append("Alt Names", altNamesEntry)

	imoNumberEntry := widget.NewEntry()
	imoNumberEntry.PlaceHolder = "IMO number"
	form.Append("IMO Number", imoNumberEntry)

	typeSelect := newSelect("VesselType")
	typeSelect.PlaceHolder = "Select vessel type"
	form.Append("Type", typeSelect)

	flagEntry := widget.NewEntry()
	flagEntry.PlaceHolder = "Country flag"
	form.Append("Flag", flagEntry)

	mmsiEntry := widget.NewEntry()
	mmsiEntry.PlaceHolder = "MMSI number"
	form.Append("MMSI", mmsiEntry)

	callSignEntry := widget.NewEntry()
	callSignEntry.PlaceHolder = "Call sign"
	form.Append("Call Sign", callSignEntry)

	return container.NewVBox(
		widget.NewLabelWithStyle("Vessel Fields", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		form,
	)
}

// Helper function to create a select widget with model values
func newSelect(modelName string) *widget.Select {
	var values []string

	// Try to get values from models, but provide defaults if it fails
	var err error
	values, err = ast.ExtractVariablesOfType(watchman.ModelsFilesystem, "pkg/search/models.go", modelName)
	if err != nil {
		// Use sensible defaults if extraction fails
		switch modelName {
		case "SourceList":
			values = []string{"SDN", "Consolidated", "DPL", "SSI", "FSE", "PLC"}
		case "Gender":
			values = []string{"Male", "Female", "Unknown"}
		case "AircraftType":
			values = []string{"Commercial", "Military", "Private"}
		case "VesselType":
			values = []string{"Cargo", "Tanker", "Passenger", "Fishing"}
		default:
			values = []string{}
		}

		fyne.LogError("Error extracting "+modelName+" values, using defaults", err)
	}

	return widget.NewSelect(values, nil)
}

func updateResultsDisplay(env Environment, resultsContent *fyne.Container, entities []search.SearchedEntity[search.Value]) {
	resultsContent.RemoveAll()

	if len(entities) == 0 {
		resultsContent.Add(widget.NewLabel("No matching results found"))
		resultsContent.Refresh()
		return
	}

	resultTabs := container.NewAppTabs()

	// Summary tab
	summaryTable := createSummaryTable(env, entities)
	resultTabs.Append(container.NewTabItem("Summary", summaryTable))

	// Entity tabs
	for _, entity := range entities {
		entityDetails := createEntityDetailsCard(env, entity)

		resultTabs.Append(container.NewTabItem(
			fmt.Sprintf("%s (%s)", entity.Entity.Name, formatEntityType(entity.Entity.Type)),
			entityDetails,
		))
	}

	resultsContent.Add(resultTabs)
	resultsContent.Refresh()
}

func createSummaryTable(env Environment, entities []search.SearchedEntity[search.Value]) fyne.CanvasObject {
	table := widget.NewTable(
		func() (int, int) { return len(entities) + 1, 5 },
		func() fyne.CanvasObject { return widget.NewLabel("                         ") },
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label, ok := cell.(*widget.Label)
			if !ok {
				env.Logger.Error().LogErrorf("unexpected %T (wanted *widget.Label)", cell)
				return
			}
			label.Alignment = fyne.TextAlignLeading
			label.Wrapping = fyne.TextTruncate
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
			entity := entities[id.Row-1]
			switch id.Col {
			case 0:
				label.SetText(entity.Entity.Name)
			case 1:
				label.SetText(formatEntityType(entity.Entity.Type))
			case 2:
				label.SetText(fmt.Sprintf("%.1f%%", entity.Match*100))
			case 3:
				label.SetText(string(entity.Entity.Source))
			case 4:
				label.SetText(entity.Entity.SourceID)
			}
		},
	)
	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 150)
	table.SetColumnWidth(4, 150)

	for idx := range entities {
		table.SetRowHeight(idx, 50)
	}

	return container.New(layout.NewMaxLayout(), table)
}

// Create a card that displays entity details
func createEntityDetailsCard(env Environment, entity search.SearchedEntity[search.Value]) fyne.CanvasObject {
	content := container.NewVBox()

	// Add match score indicator
	matchScore := widget.NewLabelWithStyle(
		fmt.Sprintf("Match Score: %.1f%%", entity.Match*100),
		fyne.TextAlignTrailing,
		fyne.TextStyle{Bold: true},
	)
	content.Add(matchScore)
	content.Add(widget.NewSeparator())

	// Add basic entity information
	basicInfo := widget.NewForm(
		widget.NewFormItem("Name", widget.NewLabel(entity.Entity.Name)),
		widget.NewFormItem("Type", widget.NewLabel(formatEntityType(entity.Entity.Type))),
	)

	if string(entity.Entity.Source) != "" {
		basicInfo.Append("Source", widget.NewLabel(string(entity.Entity.Source)))
	}

	if entity.Entity.SourceID != "" {
		basicInfo.Append("Source ID", widget.NewLabel(entity.Entity.SourceID))
	}

	basicInfoCard := widget.NewCard("Basic Information", "", basicInfo)
	content.Add(basicInfoCard)

	// Add entity-specific information based on type
	switch entity.Entity.Type {
	case search.EntityPerson:
		if entity.Entity.Person != nil {
			content.Add(createPersonInfoCard(*entity.Entity.Person))
		}
	case search.EntityBusiness:
		if entity.Entity.Business != nil {
			content.Add(createBusinessInfoCard(*entity.Entity.Business))
		}
	case search.EntityOrganization:
		if entity.Entity.Organization != nil {
			content.Add(createOrganizationInfoCard(*entity.Entity.Organization))
		}
	case search.EntityAircraft:
		if entity.Entity.Aircraft != nil {
			content.Add(createAircraftInfoCard(*entity.Entity.Aircraft))
		}
	case search.EntityVessel:
		if entity.Entity.Vessel != nil {
			content.Add(createVesselInfoCard(*entity.Entity.Vessel))
		}
	}

	// Add contact information if available
	if hasContactInfo(entity.Entity.Contact) {
		content.Add(createContactInfoCard(entity.Entity.Contact))
	}

	// Add addresses if available
	if len(entity.Entity.Addresses) > 0 {
		content.Add(createAddressesCard(entity.Entity.Addresses))
	}

	// Wrap in a scroll container
	detailsScroll := container.NewVScroll(content)

	device := fyne.CurrentDevice()
	if device.IsBrowser() {
		detailsScroll.SetMinSize(fyne.NewSize(0, env.Height*0.20))
	} else {
		detailsScroll.SetMinSize(fyne.NewSize(0, env.Height*0.40))
	}

	return detailsScroll
}

func wordWrappingLabel(text string) *widget.Label {
	label := widget.NewLabel(text)
	label.Wrapping = fyne.TextWrapWord

	return label
}

func createPersonInfoCard(person search.Person) fyne.CanvasObject {
	form := widget.NewForm()

	if len(person.AltNames) > 0 {
		form.Append("Alt Names", wordWrappingLabel(strings.Join(person.AltNames, ", ")))
	}

	if person.Gender != "" {
		form.Append("Gender", wordWrappingLabel(string(person.Gender)))
	}

	if person.BirthDate != nil {
		form.Append("Birth Date", wordWrappingLabel(person.BirthDate.Format("2006-01-02")))
	}

	if person.DeathDate != nil {
		form.Append("Death Date", wordWrappingLabel(person.DeathDate.Format("2006-01-02")))
	}

	if len(person.Titles) > 0 {
		form.Append("Titles", wordWrappingLabel(strings.Join(person.Titles, ", ")))
	}

	if len(person.GovernmentIDs) > 0 {
		idStrings := make([]string, len(person.GovernmentIDs))
		for i, id := range person.GovernmentIDs {
			idStrings[i] = fmt.Sprintf("%s: %s (%s)", id.Type, id.Identifier, id.Country)
		}
		form.Append("Government IDs", wordWrappingLabel(strings.Join(idStrings, "\n")))
	}

	return widget.NewCard("Person Details", "", form)
}

func createBusinessInfoCard(business search.Business) fyne.CanvasObject {
	form := widget.NewForm()

	if len(business.AltNames) > 0 {
		form.Append("Alt Names", wordWrappingLabel(strings.Join(business.AltNames, ", ")))
	}

	if business.Created != nil {
		form.Append("Created", wordWrappingLabel(business.Created.Format("2006-01-02")))
	}

	if business.Dissolved != nil {
		form.Append("Dissolved", wordWrappingLabel(business.Dissolved.Format("2006-01-02")))
	}

	if len(business.GovernmentIDs) > 0 {
		idStrings := make([]string, len(business.GovernmentIDs))
		for i, id := range business.GovernmentIDs {
			idStrings[i] = fmt.Sprintf("%s: %s (%s)", id.Type, id.Identifier, id.Country)
		}
		form.Append("Government IDs", wordWrappingLabel(strings.Join(idStrings, "\n")))
	}

	return widget.NewCard("Business Details", "", form)
}

func createOrganizationInfoCard(org search.Organization) fyne.CanvasObject {
	form := widget.NewForm()

	if len(org.AltNames) > 0 {
		form.Append("Alt Names", wordWrappingLabel(strings.Join(org.AltNames, ", ")))
	}

	if org.Created != nil {
		form.Append("Created", wordWrappingLabel(org.Created.Format("2006-01-02")))
	}

	if org.Dissolved != nil {
		form.Append("Dissolved", wordWrappingLabel(org.Dissolved.Format("2006-01-02")))
	}

	if len(org.GovernmentIDs) > 0 {
		idStrings := make([]string, len(org.GovernmentIDs))
		for i, id := range org.GovernmentIDs {
			idStrings[i] = fmt.Sprintf("%s: %s (%s)", id.Type, id.Identifier, id.Country)
		}
		form.Append("Government IDs", wordWrappingLabel(strings.Join(idStrings, "\n")))
	}

	return widget.NewCard("Organization Details", "", form)
}

func createAircraftInfoCard(aircraft search.Aircraft) fyne.CanvasObject {
	form := widget.NewForm()

	if len(aircraft.AltNames) > 0 {
		form.Append("Alt Names", wordWrappingLabel(strings.Join(aircraft.AltNames, ", ")))
	}

	if aircraft.Type != "" {
		form.Append("Type", wordWrappingLabel(string(aircraft.Type)))
	}

	if aircraft.Flag != "" {
		form.Append("Flag", wordWrappingLabel(aircraft.Flag))
	}

	if aircraft.Built != nil {
		form.Append("Built", wordWrappingLabel(aircraft.Built.Format("2006-01-02")))
	}

	if aircraft.ICAOCode != "" {
		form.Append("ICAO Code", wordWrappingLabel(aircraft.ICAOCode))
	}

	if aircraft.Model != "" {
		form.Append("Model", wordWrappingLabel(aircraft.Model))
	}

	if aircraft.SerialNumber != "" {
		form.Append("Serial Number", wordWrappingLabel(aircraft.SerialNumber))
	}

	return widget.NewCard("Aircraft Details", "", form)
}

func createVesselInfoCard(vessel search.Vessel) fyne.CanvasObject {
	form := widget.NewForm()

	if len(vessel.AltNames) > 0 {
		form.Append("Alt Names", wordWrappingLabel(strings.Join(vessel.AltNames, ", ")))
	}

	if vessel.IMONumber != "" {
		form.Append("IMO Number", wordWrappingLabel(vessel.IMONumber))
	}

	if vessel.Type != "" {
		form.Append("Type", wordWrappingLabel(string(vessel.Type)))
	}

	if vessel.Flag != "" {
		form.Append("Flag", wordWrappingLabel(vessel.Flag))
	}

	if vessel.Built != nil {
		form.Append("Built", wordWrappingLabel(vessel.Built.Format("2006-01-02")))
	}

	if vessel.Model != "" {
		form.Append("Model", wordWrappingLabel(vessel.Model))
	}

	if vessel.Tonnage > 0 {
		form.Append("Tonnage", wordWrappingLabel(fmt.Sprintf("%d", vessel.Tonnage)))
	}

	if vessel.MMSI != "" {
		form.Append("MMSI", wordWrappingLabel(vessel.MMSI))
	}

	if vessel.CallSign != "" {
		form.Append("Call Sign", wordWrappingLabel(vessel.CallSign))
	}

	if vessel.GrossRegisteredTonnage > 0 {
		form.Append("Gross Registered Tonnage", wordWrappingLabel(fmt.Sprintf("%d", vessel.GrossRegisteredTonnage)))
	}

	if vessel.Owner != "" {
		form.Append("Owner", wordWrappingLabel(vessel.Owner))
	}

	return widget.NewCard("Vessel Details", "", form)
}

func createContactInfoCard(contact search.ContactInfo) fyne.CanvasObject {
	form := widget.NewForm()

	if len(contact.EmailAddresses) > 0 {
		form.Append("Email Addresses", wordWrappingLabel(strings.Join(contact.EmailAddresses, "\n")))
	}

	if len(contact.PhoneNumbers) > 0 {
		form.Append("Phone Numbers", wordWrappingLabel(strings.Join(contact.PhoneNumbers, "\n")))
	}

	if len(contact.FaxNumbers) > 0 {
		form.Append("Fax Numbers", wordWrappingLabel(strings.Join(contact.FaxNumbers, "\n")))
	}

	if len(contact.Websites) > 0 {
		form.Append("Websites", wordWrappingLabel(strings.Join(contact.Websites, "\n")))
	}

	return widget.NewCard("Contact Information", "", form)
}

func createAddressesCard(addresses []search.Address) fyne.CanvasObject {
	container := container.NewVBox()

	for _, address := range addresses {
		addressForm := widget.NewForm()

		if address.Line1 != "" {
			addressForm.Append("Line 1", wordWrappingLabel(address.Line1))
		}

		if address.Line2 != "" {
			addressForm.Append("Line 2", wordWrappingLabel(address.Line2))
		}

		if address.City != "" {
			addressForm.Append("City", wordWrappingLabel(address.City))
		}

		if address.State != "" {
			addressForm.Append("State", wordWrappingLabel(address.State))
		}

		if address.PostalCode != "" {
			addressForm.Append("Postal Code", wordWrappingLabel(address.PostalCode))
		}

		if address.Country != "" {
			addressForm.Append("Country", wordWrappingLabel(address.Country))
		}

		if address.Latitude != 0 || address.Longitude != 0 {
			addressForm.Append("Coordinates", wordWrappingLabel(
				fmt.Sprintf("Lat: %.6f, Lon: %.6f", address.Latitude, address.Longitude),
			))
		}

		container.Add(widget.NewCard("Address", "", addressForm))
	}

	return container
}

// Helper function to check if contact info has any data
func hasContactInfo(contact search.ContactInfo) bool {
	return len(contact.EmailAddresses) > 0 ||
		len(contact.PhoneNumbers) > 0 ||
		len(contact.FaxNumbers) > 0 ||
		len(contact.Websites) > 0
}

// Helper function to format entity type for display
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
