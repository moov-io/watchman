package ui

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/ast"
	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var errMinMatch = errors.New("min score must be between 0 and 1, or a percent from 1 to 100")

type personForm struct {
	altNames      *widget.Entry
	gender        *widget.Select
	birthDate     *widget.Entry
	governmentIDs *widget.Entry
	root          fyne.CanvasObject
}

func newPersonForm() personForm {
	f := personForm{
		altNames:      multiEntry("Alternative names, one per line"),
		gender:        newModelSelect("Gender", []string{"Male", "Female", "Unknown"}),
		birthDate:     singleEntry("YYYY-MM-DD"),
		governmentIDs: multiEntry("ID type:number:country, one per line"),
	}
	f.gender.PlaceHolder = "Select gender"
	f.root = formOf(
		widget.NewFormItem("Alt names", f.altNames),
		widget.NewFormItem("Gender", f.gender),
		widget.NewFormItem("Birth date", f.birthDate),
		widget.NewFormItem("Government IDs", f.governmentIDs),
	)
	return f
}

func (f personForm) fields(name string) *search.Person {
	person := &search.Person{
		Name:          name,
		AltNames:      linesOf(f.altNames.Text),
		GovernmentIDs: parseGovernmentIDs(f.governmentIDs.Text),
		BirthDate:     parseDate(f.birthDate.Text),
	}
	if f.gender.Selected != "" {
		person.Gender = search.Gender(strings.ToLower(f.gender.Selected))
	}
	return person
}

type businessForm struct {
	altNames      *widget.Entry
	created       *widget.Entry
	dissolved     *widget.Entry
	governmentIDs *widget.Entry
	root          fyne.CanvasObject
}

func newBusinessForm() businessForm {
	f := businessForm{
		altNames:      multiEntry("Alternative names, one per line"),
		created:       singleEntry("YYYY-MM-DD"),
		dissolved:     singleEntry("YYYY-MM-DD"),
		governmentIDs: multiEntry("ID type:number:country, one per line"),
	}
	f.root = formOf(
		widget.NewFormItem("Alt names", f.altNames),
		widget.NewFormItem("Created", f.created),
		widget.NewFormItem("Dissolved", f.dissolved),
		widget.NewFormItem("Government IDs", f.governmentIDs),
	)
	return f
}

func (f businessForm) fields(name string) *search.Business {
	return &search.Business{
		Name:          name,
		AltNames:      linesOf(f.altNames.Text),
		Created:       parseDate(f.created.Text),
		Dissolved:     parseDate(f.dissolved.Text),
		GovernmentIDs: parseGovernmentIDs(f.governmentIDs.Text),
	}
}

type organizationForm struct {
	altNames      *widget.Entry
	created       *widget.Entry
	dissolved     *widget.Entry
	governmentIDs *widget.Entry
	root          fyne.CanvasObject
}

func newOrganizationForm() organizationForm {
	f := organizationForm{
		altNames:      multiEntry("Alternative names, one per line"),
		created:       singleEntry("YYYY-MM-DD"),
		dissolved:     singleEntry("YYYY-MM-DD"),
		governmentIDs: multiEntry("ID type:number:country, one per line"),
	}
	f.root = formOf(
		widget.NewFormItem("Alt names", f.altNames),
		widget.NewFormItem("Created", f.created),
		widget.NewFormItem("Dissolved", f.dissolved),
		widget.NewFormItem("Government IDs", f.governmentIDs),
	)
	return f
}

func (f organizationForm) fields(name string) *search.Organization {
	return &search.Organization{
		Name:          name,
		AltNames:      linesOf(f.altNames.Text),
		Created:       parseDate(f.created.Text),
		Dissolved:     parseDate(f.dissolved.Text),
		GovernmentIDs: parseGovernmentIDs(f.governmentIDs.Text),
	}
}

type aircraftForm struct {
	altNames     *widget.Entry
	kind         *widget.Select
	flag         *widget.Entry
	built        *widget.Entry
	icaoCode     *widget.Entry
	model        *widget.Entry
	serialNumber *widget.Entry
	root         fyne.CanvasObject
}

func newAircraftForm() aircraftForm {
	f := aircraftForm{
		altNames:     multiEntry("Alternative names, one per line"),
		kind:         newModelSelect("AircraftType", []string{"Cargo", "Unknown"}),
		flag:         singleEntry("Country flag"),
		built:        singleEntry("YYYY-MM-DD"),
		icaoCode:     singleEntry("ICAO code"),
		model:        singleEntry("Aircraft model"),
		serialNumber: singleEntry("Serial number"),
	}
	f.kind.PlaceHolder = "Select aircraft type"
	f.root = formOf(
		widget.NewFormItem("Alt names", f.altNames),
		widget.NewFormItem("Type", f.kind),
		widget.NewFormItem("Flag", f.flag),
		widget.NewFormItem("Built", f.built),
		widget.NewFormItem("ICAO code", f.icaoCode),
		widget.NewFormItem("Model", f.model),
		widget.NewFormItem("Serial number", f.serialNumber),
	)
	return f
}

func (f aircraftForm) fields(name string) *search.Aircraft {
	aircraft := &search.Aircraft{
		Name:         name,
		AltNames:     linesOf(f.altNames.Text),
		Flag:         strings.TrimSpace(f.flag.Text),
		Built:        parseDate(f.built.Text),
		ICAOCode:     strings.TrimSpace(f.icaoCode.Text),
		Model:        strings.TrimSpace(f.model.Text),
		SerialNumber: strings.TrimSpace(f.serialNumber.Text),
	}
	if f.kind.Selected != "" {
		aircraft.Type = search.AircraftType(strings.ToLower(f.kind.Selected))
	}
	return aircraft
}

type vesselForm struct {
	altNames               *widget.Entry
	imoNumber              *widget.Entry
	kind                   *widget.Select
	flag                   *widget.Entry
	mmsi                   *widget.Entry
	callSign               *widget.Entry
	tonnage                *widget.Entry
	grossRegisteredTonnage *widget.Entry
	owner                  *widget.Entry
	built                  *widget.Entry
	model                  *widget.Entry
	root                   fyne.CanvasObject
}

func newVesselForm() vesselForm {
	f := vesselForm{
		altNames:               multiEntry("Alternative names, one per line"),
		imoNumber:              singleEntry("IMO number"),
		kind:                   newModelSelect("VesselType", []string{"Cargo", "Unknown"}),
		flag:                   singleEntry("Country flag"),
		mmsi:                   singleEntry("MMSI number"),
		callSign:               singleEntry("Call sign"),
		tonnage:                singleEntry("e.g. 85000"),
		grossRegisteredTonnage: singleEntry("Gross registered tonnage"),
		owner:                  singleEntry("Owner name"),
		built:                  singleEntry("YYYY-MM-DD"),
		model:                  singleEntry("Vessel model"),
	}
	f.kind.PlaceHolder = "Select vessel type"
	f.root = formOf(
		widget.NewFormItem("Alt names", f.altNames),
		widget.NewFormItem("IMO number", f.imoNumber),
		widget.NewFormItem("Type", f.kind),
		widget.NewFormItem("Flag", f.flag),
		widget.NewFormItem("MMSI", f.mmsi),
		widget.NewFormItem("Call sign", f.callSign),
		widget.NewFormItem("Tonnage", f.tonnage),
		widget.NewFormItem("Gross registered tonnage", f.grossRegisteredTonnage),
		widget.NewFormItem("Owner", f.owner),
		widget.NewFormItem("Built", f.built),
		widget.NewFormItem("Model", f.model),
	)
	return f
}

func (f vesselForm) fields(name string) *search.Vessel {
	vessel := &search.Vessel{
		Name:                   name,
		AltNames:               linesOf(f.altNames.Text),
		IMONumber:              strings.TrimSpace(f.imoNumber.Text),
		Flag:                   strings.TrimSpace(f.flag.Text),
		MMSI:                   strings.TrimSpace(f.mmsi.Text),
		CallSign:               strings.TrimSpace(f.callSign.Text),
		Tonnage:                parsePositiveInt(f.tonnage.Text),
		GrossRegisteredTonnage: parsePositiveInt(f.grossRegisteredTonnage.Text),
		Owner:                  strings.TrimSpace(f.owner.Text),
		Built:                  parseDate(f.built.Text),
		Model:                  strings.TrimSpace(f.model.Text),
	}
	if f.kind.Selected != "" {
		vessel.Type = search.VesselType(strings.ToLower(f.kind.Selected))
	}
	return vessel
}

func formOf(items ...*widget.FormItem) fyne.CanvasObject {
	form := widget.NewForm(items...)
	return container.NewVBox(form)
}

func singleEntry(placeholder string) *widget.Entry {
	entry := widget.NewEntry()
	entry.PlaceHolder = placeholder
	return entry
}

func multiEntry(placeholder string) *widget.Entry {
	entry := widget.NewMultiLineEntry()
	entry.PlaceHolder = placeholder
	entry.SetMinRowsVisible(2)
	return entry
}

func newSourceSelect() *widget.Select {
	values := modelValues("SourceList", []string{
		"us_ofac", "us_csl", "us_non_sdn", "us_tel", "us_fincen_311", "eu_csl", "uk_csl", "un_csl",
	})
	filtered := make([]string, 0, len(values))
	for _, value := range values {
		switch value {
		case "", "api-request", "mcp-request":
			continue
		default:
			filtered = append(filtered, value)
		}
	}
	return widget.NewSelect(filtered, nil)
}

func newModelSelect(modelName string, defaults []string) *widget.Select {
	values := modelValues(modelName, defaults)
	display := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" {
			continue
		}
		display = append(display, displayEnum(value))
	}
	return widget.NewSelect(display, nil)
}

func modelValues(modelName string, defaults []string) []string {
	values, err := ast.ExtractVariablesOfType(watchman.ModelsFilesystem, "pkg/search/models.go", modelName)
	if err != nil || len(values) == 0 {
		if err != nil {
			fyne.LogError("reading "+modelName+" values", err)
		}
		return defaults
	}
	return values
}

func displayEnum(value string) string {
	if value == "" {
		return value
	}
	return strings.ToUpper(value[:1]) + value[1:]
}

func entityTypeFromLabel(label string) string {
	switch strings.ToLower(strings.TrimSpace(label)) {
	case "person":
		return "Person"
	case "business":
		return "Business"
	case "organization":
		return "Organization"
	case "aircraft":
		return "Aircraft"
	case "vessel":
		return "Vessel"
	default:
		return ""
	}
}

func linesOf(text string) []string {
	if strings.TrimSpace(text) == "" {
		return nil
	}
	var out []string
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			out = append(out, line)
		}
	}
	return out
}

func parseDate(text string) *time.Time {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", text)
	if err != nil {
		return nil
	}
	return &t
}

func parsePositiveInt(text string) int {
	text = strings.TrimSpace(text)
	if text == "" {
		return 0
	}
	n, err := strconv.Atoi(text)
	if err != nil || n <= 0 {
		return 0
	}
	return n
}

func parseLimit(selected string) int {
	n, err := strconv.Atoi(strings.TrimSpace(selected))
	if err != nil || n <= 0 {
		return 5
	}
	return n
}

// parseMinMatch accepts a 0–1 score ("0.80" or "1") or a percent ("80" or "80%").
// A trailing % always means a percent, including "1%". A bare number above 1 is
// read as a percent; 1 itself is an exact match.
func parseMinMatch(text string) (float64, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return 0, nil
	}
	percent := strings.HasSuffix(text, "%")
	if percent {
		text = strings.TrimSpace(strings.TrimSuffix(text, "%"))
	}
	value, err := strconv.ParseFloat(text, 64)
	if err != nil || value < 0 || text == "" {
		return 0, errMinMatch
	}
	if percent || value > 1 {
		if value > 100 {
			return 0, errMinMatch
		}
		value /= 100
	}
	return value, nil
}

func parseGovernmentIDs(text string) []search.GovernmentID {
	lines := linesOf(text)
	if len(lines) == 0 {
		return nil
	}
	ids := make([]search.GovernmentID, 0, len(lines))
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 3)
		var id search.GovernmentID
		switch len(parts) {
		case 3:
			id.Type = search.GovernmentIDType(strings.TrimSpace(parts[0]))
			id.Identifier = strings.TrimSpace(parts[1])
			id.Country = strings.TrimSpace(parts[2])
		case 2:
			id.Type = search.GovernmentIDType(strings.TrimSpace(parts[0]))
			id.Identifier = strings.TrimSpace(parts[1])
		default:
			id.Identifier = strings.TrimSpace(parts[0])
		}
		if id.Identifier != "" {
			ids = append(ids, id)
		}
	}
	return ids
}
