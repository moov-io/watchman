package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// ExpandableSection is a custom widget that can be expanded/collapsed
type ExpandableSection struct {
	widget.BaseWidget
	Title      string
	Content    fyne.CanvasObject
	Expanded   bool
	titleLabel *widget.Label
	icon       *widget.Icon
	container  *fyne.Container
}

// NewExpandableSection creates a new expandable section
func NewExpandableSection(title string, content fyne.CanvasObject) *ExpandableSection {
	section := &ExpandableSection{
		Title:    title,
		Content:  content,
		Expanded: false,
	}
	section.ExtendBaseWidget(section)
	return section
}

// CreateRenderer implements the widget.Widget interface
func (e *ExpandableSection) CreateRenderer() fyne.WidgetRenderer {
	e.titleLabel = widget.NewLabel(e.Title)
	e.titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Use NavigateNextIcon for collapsed state (points right)
	e.icon = widget.NewIcon(theme.NavigateNextIcon())

	header := container.NewHBox(
		e.icon,
		e.titleLabel,
	)

	// Make the header clickable
	headerButton := widget.NewButton("", func() {
		e.Expanded = !e.Expanded
		e.Refresh()
	})
	headerButton.Importance = widget.LowImportance

	// Overlay the button on the header for clickability
	headerWithButton := container.NewStack(
		header,
		headerButton,
	)

	// Main container that will show/hide content
	e.container = container.NewVBox(
		headerWithButton,
	)

	if e.Expanded {
		// Use MenuDropDownIcon for expanded state (points down)
		e.icon.SetResource(theme.MenuDropDownIcon())
		e.container.Add(e.Content)
	}

	return &expandableSectionRenderer{
		section:    e,
		background: canvas.NewRectangle(color.Transparent),
		objects:    []fyne.CanvasObject{e.container},
	}
}

// Fixed version of the Refresh method in ExpandableSection
func (e *ExpandableSection) Refresh() {
	if e.Expanded {
		e.icon.SetResource(theme.MenuDropDownIcon())
		// Only add if not already there
		if len(e.container.Objects) == 1 {
			e.container.Add(e.Content)
		}
	} else {
		e.icon.SetResource(theme.NavigateNextIcon())
		// Only remove if it's there
		if len(e.container.Objects) > 1 {
			e.container.Remove(e.Content)
		}
	}

	// No parent access, just refresh the container
	e.container.Refresh()
	e.BaseWidget.Refresh()
	// Force a resize on browser to ensure proper layout
	if fyne.CurrentDevice().IsBrowser() {
		e.container.Resize(e.container.MinSize())
	}
}

// expandableSectionRenderer implements the fyne.WidgetRenderer interface
type expandableSectionRenderer struct {
	section    *ExpandableSection
	background *canvas.Rectangle
	objects    []fyne.CanvasObject
}

func (r *expandableSectionRenderer) Destroy() {}

func (r *expandableSectionRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.objects[0].Resize(size)
}

func (r *expandableSectionRenderer) MinSize() fyne.Size {
	return r.objects[0].MinSize()
}

func (r *expandableSectionRenderer) Objects() []fyne.CanvasObject {
	return append([]fyne.CanvasObject{r.background}, r.objects...)
}

func (r *expandableSectionRenderer) Refresh() {
	r.background.Refresh()
	r.objects[0].Refresh()
}
