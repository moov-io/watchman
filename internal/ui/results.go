package ui

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// resultsPane is the middle and right columns: the match list, and the
// selected entity. Detail widgets are built for the selected row only.
type resultsPane struct {
	listColumn   fyne.CanvasObject
	detailColumn fyne.CanvasObject
	status       *widget.Label
	list         *widget.List
	detail       *fyne.Container
	scroll       *container.Scroll

	entities []search.SearchedEntity[search.Value]
}

func newResultsPane() *resultsPane {
	pane := &resultsPane{}

	pane.status = widget.NewLabel("Run a search to see matches.")
	pane.status.Wrapping = fyne.TextWrapWord
	pane.status.Importance = widget.LowImportance

	pane.detail = container.NewVBox(quietLabel("Select a match to see details."))
	pane.scroll = container.NewVScroll(pane.detail)

	pane.list = widget.NewList(
		func() int { return len(pane.entities) },
		func() fyne.CanvasObject { return newResultRow() },
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*resultRow).set(pane.entities[id])
		},
	)
	pane.list.OnSelected = func(id widget.ListItemID) {
		if id < 0 || id >= len(pane.entities) {
			return
		}
		pane.showDetail(pane.entities[id])
	}

	pane.listColumn = container.NewBorder(
		container.NewVBox(sectionLabel("Results"), pane.status),
		nil, nil, nil,
		pane.list,
	)
	pane.detailColumn = container.NewBorder(
		sectionLabel("Details"),
		nil, nil, nil,
		pane.scroll,
	)
	return pane
}

func (p *resultsPane) showSearching() {
	p.status.Importance = widget.MediumImportance
	p.status.SetText("Searching…")
	p.entities = nil
	p.list.UnselectAll()
	p.list.Refresh()
	p.setDetail(quietLabel("Searching…"))
}

func (p *resultsPane) showError(err error) {
	p.status.Importance = widget.DangerImportance
	p.status.SetText(err.Error())
	p.setDetail(quietLabel("Search failed."))
}

func (p *resultsPane) show(entities []search.SearchedEntity[search.Value], elapsed time.Duration) {
	p.entities = entities
	p.list.UnselectAll()
	p.list.Refresh()

	if len(entities) == 0 {
		p.status.Importance = widget.LowImportance
		p.status.SetText(fmt.Sprintf("No matches in %s.", formatElapsed(elapsed)))
		p.setDetail(quietLabel("Nothing on the loaded lists matched this query."))
		return
	}

	p.status.Importance = widget.LowImportance
	noun := "matches"
	if len(entities) == 1 {
		noun = "match"
	}
	p.status.SetText(fmt.Sprintf("%s %s in %s.", formatCount(len(entities)), noun, formatElapsed(elapsed)))
	p.list.Select(0)
}

func (p *resultsPane) showDetail(entity search.SearchedEntity[search.Value]) {
	p.setDetail(entityDetails(entity))
	p.scroll.ScrollToTop()
}

func (p *resultsPane) setDetail(obj fyne.CanvasObject) {
	p.detail.RemoveAll()
	p.detail.Add(obj)
	p.detail.Refresh()
}

func quietLabel(text string) *widget.Label {
	label := wordWrappingLabel(text)
	label.Importance = widget.LowImportance
	return label
}

type resultRow struct {
	widget.BaseWidget

	name  *widget.Label
	meta  *widget.Label
	score *widget.Label
}

func newResultRow() *resultRow {
	row := &resultRow{
		name:  widget.NewLabel(""),
		meta:  widget.NewLabel(""),
		score: widget.NewLabel("100.0%"),
	}
	row.name.Truncation = fyne.TextTruncateEllipsis
	row.meta.Truncation = fyne.TextTruncateEllipsis
	row.meta.Importance = widget.LowImportance
	row.score.Alignment = fyne.TextAlignTrailing
	row.score.TextStyle = fyne.TextStyle{Bold: true}
	row.ExtendBaseWidget(row)
	return row
}

func (r *resultRow) set(entity search.SearchedEntity[search.Value]) {
	r.name.SetText(entity.Name)
	r.meta.SetText(resultMeta(entity))
	r.score.Importance = matchImportance(entity.Match)
	r.score.SetText(formatMatch(entity.Match))
}

func (r *resultRow) CreateRenderer() fyne.WidgetRenderer {
	texts := container.NewVBox(r.name, r.meta)
	return widget.NewSimpleRenderer(container.NewBorder(nil, nil, nil, r.score, texts))
}

func resultMeta(entity search.SearchedEntity[search.Value]) string {
	parts := []string{formatEntityType(entity.Type)}
	if src := string(entity.Source); src != "" {
		parts = append(parts, src)
	}
	if entity.SourceID != "" {
		parts = append(parts, entity.SourceID)
	}
	return strings.Join(parts, " · ")
}

func entityDetails(entity search.SearchedEntity[search.Value]) fyne.CanvasObject {
	content := container.NewVBox()

	score := widget.NewLabelWithStyle(formatMatch(entity.Match)+" match", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	score.Importance = matchImportance(entity.Match)
	score.Selectable = true
	content.Add(score)

	basic := widget.NewForm(
		widget.NewFormItem("Name", wordWrappingLabel(entity.Name)),
		widget.NewFormItem("Type", wordWrappingLabel(formatEntityType(entity.Type))),
	)
	if src := string(entity.Source); src != "" {
		basic.Append("Source", wordWrappingLabel(src))
	}
	if entity.SourceID != "" {
		basic.Append("Source ID", selectableLabel(entity.SourceID))
	}
	content.Add(widget.NewCard("Identity", "", basic))

	switch entity.Type {
	case search.EntityPerson:
		if entity.Person != nil {
			content.Add(personCard(*entity.Person))
		}
	case search.EntityBusiness:
		if entity.Business != nil {
			content.Add(businessCard(*entity.Business))
		}
	case search.EntityOrganization:
		if entity.Organization != nil {
			content.Add(organizationCard(*entity.Organization))
		}
	case search.EntityAircraft:
		if entity.Aircraft != nil {
			content.Add(aircraftCard(*entity.Aircraft))
		}
	case search.EntityVessel:
		if entity.Vessel != nil {
			content.Add(vesselCard(*entity.Vessel))
		}
	}

	if hasContactInfo(entity.Contact) {
		content.Add(contactCard(entity.Contact))
	}
	if len(entity.Addresses) > 0 {
		content.Add(addressesCard(entity.Addresses))
	}
	if len(entity.CryptoAddresses) > 0 {
		content.Add(cryptoCard(entity.CryptoAddresses))
	}
	if entity.SanctionsInfo != nil {
		content.Add(sanctionsCard(*entity.SanctionsInfo))
	}
	if len(entity.Affiliations) > 0 {
		content.Add(affiliationsCard(entity.Affiliations))
	}
	if len(entity.HistoricalInfo) > 0 {
		content.Add(historicalCard(entity.HistoricalInfo))
	}
	if breakdown := scoreBreakdown(entity); breakdown != nil {
		content.Add(breakdown)
	}

	return content
}

func selectableLabel(text string) *widget.Label {
	label := wordWrappingLabel(text)
	label.Selectable = true
	label.TextStyle = fyne.TextStyle{Monospace: true}
	return label
}

func personCard(person search.Person) fyne.CanvasObject {
	form := widget.NewForm()
	appendLines(form, "Alt names", person.AltNames)
	if person.Gender != "" {
		form.Append("Gender", wordWrappingLabel(displayEnum(string(person.Gender))))
	}
	if person.BirthDate != nil {
		form.Append("Birth date", wordWrappingLabel(person.BirthDate.Format("2006-01-02")))
	}
	if person.DeathDate != nil {
		form.Append("Death date", wordWrappingLabel(person.DeathDate.Format("2006-01-02")))
	}
	if person.PlaceOfBirth != "" {
		form.Append("Place of birth", wordWrappingLabel(person.PlaceOfBirth))
	}
	appendLines(form, "Titles", person.Titles)
	appendGovernmentIDs(form, person.GovernmentIDs)
	return widget.NewCard("Person", "", form)
}

func businessCard(business search.Business) fyne.CanvasObject {
	form := widget.NewForm()
	appendLines(form, "Alt names", business.AltNames)
	if business.Created != nil {
		form.Append("Created", wordWrappingLabel(business.Created.Format("2006-01-02")))
	}
	if business.Dissolved != nil {
		form.Append("Dissolved", wordWrappingLabel(business.Dissolved.Format("2006-01-02")))
	}
	appendGovernmentIDs(form, business.GovernmentIDs)
	return widget.NewCard("Business", "", form)
}

func organizationCard(org search.Organization) fyne.CanvasObject {
	form := widget.NewForm()
	appendLines(form, "Alt names", org.AltNames)
	if org.Created != nil {
		form.Append("Created", wordWrappingLabel(org.Created.Format("2006-01-02")))
	}
	if org.Dissolved != nil {
		form.Append("Dissolved", wordWrappingLabel(org.Dissolved.Format("2006-01-02")))
	}
	appendGovernmentIDs(form, org.GovernmentIDs)
	return widget.NewCard("Organization", "", form)
}

func aircraftCard(aircraft search.Aircraft) fyne.CanvasObject {
	form := widget.NewForm()
	appendLines(form, "Alt names", aircraft.AltNames)
	if aircraft.Type != "" {
		form.Append("Type", wordWrappingLabel(displayEnum(string(aircraft.Type))))
	}
	appendText(form, "Flag", aircraft.Flag)
	if aircraft.Built != nil {
		form.Append("Built", wordWrappingLabel(aircraft.Built.Format("2006-01-02")))
	}
	appendText(form, "ICAO code", aircraft.ICAOCode)
	appendText(form, "Model", aircraft.Model)
	appendText(form, "Serial number", aircraft.SerialNumber)
	return widget.NewCard("Aircraft", "", form)
}

func vesselCard(vessel search.Vessel) fyne.CanvasObject {
	form := widget.NewForm()
	appendLines(form, "Alt names", vessel.AltNames)
	appendText(form, "IMO number", vessel.IMONumber)
	if vessel.Type != "" {
		form.Append("Type", wordWrappingLabel(displayEnum(string(vessel.Type))))
	}
	appendText(form, "Flag", vessel.Flag)
	if vessel.Built != nil {
		form.Append("Built", wordWrappingLabel(vessel.Built.Format("2006-01-02")))
	}
	appendText(form, "Model", vessel.Model)
	if vessel.Tonnage > 0 {
		form.Append("Tonnage", wordWrappingLabel(formatCount(vessel.Tonnage)))
	}
	if vessel.GrossRegisteredTonnage > 0 {
		form.Append("Gross registered tonnage", wordWrappingLabel(formatCount(vessel.GrossRegisteredTonnage)))
	}
	appendText(form, "MMSI", vessel.MMSI)
	appendText(form, "Call sign", vessel.CallSign)
	appendText(form, "Owner", vessel.Owner)
	return widget.NewCard("Vessel", "", form)
}

func contactCard(contact search.ContactInfo) fyne.CanvasObject {
	form := widget.NewForm()
	appendLines(form, "Email", contact.EmailAddresses)
	appendLines(form, "Phone", contact.PhoneNumbers)
	appendLines(form, "Fax", contact.FaxNumbers)
	appendLines(form, "Websites", contact.Websites)
	return widget.NewCard("Contact", "", form)
}

func addressesCard(addresses []search.Address) fyne.CanvasObject {
	box := container.NewVBox()
	for i, address := range addresses {
		if i > 0 {
			box.Add(widget.NewSeparator())
		}
		box.Add(wordWrappingLabel(address.Format()))
		if address.Latitude != 0 || address.Longitude != 0 {
			coord := wordWrappingLabel(fmt.Sprintf("%.5f, %.5f", address.Latitude, address.Longitude))
			coord.Importance = widget.LowImportance
			box.Add(coord)
		}
	}
	return widget.NewCard("Addresses", "", box)
}

func cryptoCard(addresses []search.CryptoAddress) fyne.CanvasObject {
	lines := make([]string, 0, len(addresses))
	for _, address := range addresses {
		lines = append(lines, strings.TrimSpace(address.Currency+" "+address.Address))
	}
	return widget.NewCard("Crypto addresses", "", wordWrappingLabel(strings.Join(lines, "\n")))
}

func sanctionsCard(info search.SanctionsInfo) fyne.CanvasObject {
	form := widget.NewForm()
	appendLines(form, "Programs", info.Programs)
	if info.Secondary {
		form.Append("Secondary sanctions", wordWrappingLabel("Yes"))
	}
	appendText(form, "Description", info.Description)
	return widget.NewCard("Sanctions", "", form)
}

func affiliationsCard(affiliations []search.Affiliation) fyne.CanvasObject {
	lines := make([]string, 0, len(affiliations))
	for _, affiliation := range affiliations {
		line := affiliation.EntityName
		if affiliation.Type != "" {
			line = affiliation.Type + ": " + affiliation.EntityName
		}
		if affiliation.Details != "" {
			line += " (" + affiliation.Details + ")"
		}
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	return widget.NewCard("Affiliations", "", wordWrappingLabel(strings.Join(lines, "\n")))
}

func historicalCard(items []search.HistoricalInfo) fyne.CanvasObject {
	lines := make([]string, 0, len(items))
	for _, item := range items {
		line := strings.TrimSpace(item.Type + ": " + item.Value)
		if !item.Date.IsZero() {
			line += " (" + item.Date.Format("2006-01-02") + ")"
		}
		if line != ":" && line != "" {
			lines = append(lines, line)
		}
	}
	return widget.NewCard("History", "", wordWrappingLabel(strings.Join(lines, "\n")))
}

func scoreBreakdown(entity search.SearchedEntity[search.Value]) fyne.CanvasObject {
	box := container.NewVBox()
	for _, piece := range entity.Details.Pieces {
		if piece.FieldsCompared == 0 && piece.Score == 0 && !piece.Matched {
			continue
		}
		label := wordWrappingLabel(fmt.Sprintf("%s  %s", piece.PieceType, formatMatch(piece.Score)))
		label.Importance = matchImportance(piece.Score)
		if piece.Exact {
			label.SetText(label.Text + "  exact")
		}
		box.Add(label)
	}
	if text := decodeDebug(entity.Debug); text != "" {
		log := wordWrappingLabel(text)
		log.TextStyle = fyne.TextStyle{Monospace: true}
		box.Add(log)
	}
	if len(box.Objects) == 0 {
		return nil
	}
	return widget.NewAccordion(widget.NewAccordionItem("Score breakdown", box))
}

func decodeDebug(encoded string) string {
	if encoded == "" {
		return ""
	}
	raw, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return encoded
	}
	return strings.TrimSpace(string(raw))
}

func appendText(form *widget.Form, label, value string) {
	if strings.TrimSpace(value) == "" {
		return
	}
	form.Append(label, wordWrappingLabel(value))
}

func appendLines(form *widget.Form, label string, values []string) {
	if len(values) == 0 {
		return
	}
	form.Append(label, wordWrappingLabel(strings.Join(values, "\n")))
}

func appendGovernmentIDs(form *widget.Form, ids []search.GovernmentID) {
	if len(ids) == 0 {
		return
	}
	lines := make([]string, len(ids))
	for i, id := range ids {
		parts := make([]string, 0, 3)
		if id.Type != "" {
			parts = append(parts, string(id.Type))
		}
		if id.Identifier != "" {
			parts = append(parts, id.Identifier)
		}
		if id.Country != "" {
			parts = append(parts, id.Country)
		}
		lines[i] = strings.Join(parts, " · ")
	}
	form.Append("Government IDs", wordWrappingLabel(strings.Join(lines, "\n")))
}

func hasContactInfo(contact search.ContactInfo) bool {
	return len(contact.EmailAddresses) > 0 ||
		len(contact.PhoneNumbers) > 0 ||
		len(contact.FaxNumbers) > 0 ||
		len(contact.Websites) > 0
}

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
		if entityType == "" {
			return "Unknown"
		}
		return displayEnum(string(entityType))
	}
}
