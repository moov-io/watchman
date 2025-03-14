package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type defaultTheme struct {
	theme fyne.Theme
}

func newDefaultTheme(theme fyne.Theme) fyne.Theme {
	return &defaultTheme{theme: theme}
}

var _ fyne.Theme = (&defaultTheme{})

func (t *defaultTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return t.theme.Color(name, variant)
}
func (t *defaultTheme) Font(style fyne.TextStyle) fyne.Resource {
	return t.theme.Font(style)
}

func (t *defaultTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return t.theme.Icon(name)
}

func (t *defaultTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameInnerPadding:
		return 4.0
	}
	return t.theme.Size(name)
}
