package ui

import (
	"context"
	"fmt"
	"time"

	"github.com/moov-io/watchman/internal/ui/arch"
	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// SearchContainer is three columns: the query, the match list, and the
// selected entity. Each column scrolls on its own.
func SearchContainer(ctx context.Context, env Environment) fyne.CanvasObject {
	prefilled := arch.PrefillValues()

	entityType := widget.NewRadioGroup([]string{"Person", "Business", "Organization", "Aircraft", "Vessel"}, nil)
	entityType.Horizontal = true
	entityType.Required = true

	person := newPersonForm()
	business := newBusinessForm()
	organization := newOrganizationForm()
	aircraft := newAircraftForm()
	vessel := newVesselForm()

	forms := map[string]fyne.CanvasObject{
		"Person":       person.root,
		"Business":     business.root,
		"Organization": organization.root,
		"Aircraft":     aircraft.root,
		"Vessel":       vessel.root,
	}
	entityFields := container.NewVBox()
	entityType.OnChanged = func(selected string) {
		entityFields.RemoveAll()
		if form, ok := forms[selected]; ok {
			entityFields.Add(form)
		}
		entityFields.Refresh()
	}

	selected := entityTypeFromLabel(prefilled.Get("type"))
	if selected == "" {
		selected = "Person"
	}
	entityType.SetSelected(selected)

	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "Name to search"
	if name := prefilled.Get("name"); name != "" {
		nameEntry.SetText(name)
	}

	sourceSelect := newSourceSelect()
	sourceSelect.PlaceHolder = "All lists"

	limitSelect := widget.NewSelect([]string{"5", "10", "25", "50"}, nil)
	limitSelect.SetSelected("5")

	minMatchEntry := widget.NewEntry()
	minMatchEntry.PlaceHolder = "0.00–1.00, or a percent"

	// Debug scoring writes a breakdown for every candidate. It used to be on
	// for every search and the UI never rendered it.
	debugCheck := widget.NewCheck("Include score breakdown", nil)

	warning := widget.NewLabel("")
	warning.Importance = widget.DangerImportance
	warning.Wrapping = fyne.TextWrapWord
	warning.Hide()

	searchButton := widget.NewButtonWithIcon("Search", theme.SearchIcon(), nil)
	searchButton.Importance = widget.HighImportance

	results := newResultsPane()
	searching := false

	showWarning := func(text string) {
		warning.SetText(text)
		warning.Show()
		warning.Refresh()
	}

	runSearch := func() {
		if searching {
			return
		}
		warning.Hide()

		if entityType.Selected == "" {
			showWarning("Choose an entity type.")
			return
		}
		name := nameEntry.Text
		if name == "" {
			showWarning("Enter a name to search.")
			return
		}
		minMatch, err := parseMinMatch(minMatchEntry.Text)
		if err != nil {
			showWarning(err.Error())
			return
		}

		query := buildQuery(
			entityType.Selected,
			name,
			sourceSelect.Selected,
			person,
			business,
			organization,
			aircraft,
			vessel,
		)
		opts := search.SearchOpts{
			Limit:    parseLimit(limitSelect.Selected),
			MinMatch: minMatch,
			Debug:    debugCheck.Checked,
		}

		searching = true
		searchButton.Disable()
		searchButton.SetText("Searching…")
		results.showSearching()
		sendSearchNotification(entityType.Selected, query)

		go func() {
			started := time.Now()
			resp, err := env.Client.SearchByEntity(ctx, query, opts)
			elapsed := time.Since(started)
			fyne.Do(func() {
				searching = false
				searchButton.Enable()
				searchButton.SetText("Search")
				if err != nil {
					env.Logger.Error().LogErrorf("ERROR performing search: %v", err)
					results.showError(err)
					return
				}
				results.show(resp.Entities, elapsed)
			})
		}()
	}

	searchButton.OnTapped = runSearch
	nameEntry.OnSubmitted = func(string) { runSearch() }

	queryForm := widget.NewForm(
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("Source", sourceSelect),
		widget.NewFormItem("Limit", limitSelect),
		widget.NewFormItem("Min score", minMatchEntry),
	)

	form := container.NewVBox(
		sectionLabel("Entity type"),
		entityType,
		widget.NewSeparator(),
		sectionLabel("Query"),
		queryForm,
		debugCheck,
		widget.NewSeparator(),
		sectionLabel("Additional fields"),
		entityFields,
		warning,
		container.NewBorder(nil, nil, nil, searchButton),
	)
	formScroll := container.NewVScroll(form)
	formScroll.SetMinSize(fyne.NewSize(320, 240))

	matches := container.NewHSplit(results.listColumn, results.detailColumn)
	matches.SetOffset(0.40)

	screen := container.NewHSplit(formScroll, matches)
	screen.SetOffset(0.28)
	return screen
}

func buildQuery(
	entityType, name, sourceList string,
	person personForm,
	business businessForm,
	organization organizationForm,
	aircraft aircraftForm,
	vessel vesselForm,
) search.Entity[search.Value] {
	entity := search.Entity[search.Value]{Name: name}
	if sourceList != "" {
		entity.Source = search.SourceList(sourceList)
	}

	switch entityType {
	case "Person":
		entity.Type = search.EntityPerson
		entity.Person = person.fields(name)
	case "Business":
		entity.Type = search.EntityBusiness
		entity.Business = business.fields(name)
	case "Organization":
		entity.Type = search.EntityOrganization
		entity.Organization = organization.fields(name)
	case "Aircraft":
		entity.Type = search.EntityAircraft
		entity.Aircraft = aircraft.fields(name)
	case "Vessel":
		entity.Type = search.EntityVessel
		entity.Vessel = vessel.fields(name)
	}
	return entity
}

func sendSearchNotification(entityType string, query search.Entity[search.Value]) {
	device := fyne.CurrentDevice()
	if device.IsBrowser() || device.IsMobile() {
		return
	}
	fyne.CurrentApp().SendNotification(fyne.NewNotification(
		"Watchman search",
		fmt.Sprintf("%s search for %s", entityType, query.Name),
	))
}
