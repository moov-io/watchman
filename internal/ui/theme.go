package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// watchmanTheme is a light and dark palette built around the shield blue in
// the app icon. Sizes stay close to Fyne's defaults so forms remain compact
// without the cramped 4px inner padding the previous theme used.
type watchmanTheme struct {
	base fyne.Theme
}

func newWatchmanTheme() fyne.Theme {
	return &watchmanTheme{base: theme.DefaultTheme()}
}

var _ fyne.Theme = (*watchmanTheme)(nil)

func (t *watchmanTheme) Font(style fyne.TextStyle) fyne.Resource {
	return t.base.Font(style)
}

func (t *watchmanTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return t.base.Icon(name)
}

func (t *watchmanTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameInnerPadding, theme.SizeNamePadding:
		return 6
	case theme.SizeNameInputRadius, theme.SizeNameButtonRadius:
		return 6
	case theme.SizeNameHeadingText:
		return 22
	default:
		return t.base.Size(name)
	}
}

func (t *watchmanTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if variant == theme.VariantDark {
		if c, ok := darkColors[name]; ok {
			return c
		}
	} else if c, ok := lightColors[name]; ok {
		return c
	}
	return t.base.Color(name, variant)
}

var lightColors = map[fyne.ThemeColorName]color.Color{
	theme.ColorNameBackground:          color.NRGBA{R: 0xF4, G: 0xF7, B: 0xFB, A: 0xFF},
	theme.ColorNameButton:              color.NRGBA{R: 0xE7, G: 0xEE, B: 0xF8, A: 0xFF},
	theme.ColorNameDisabled:            color.NRGBA{R: 0x8E, G: 0x9A, B: 0xAD, A: 0xFF},
	theme.ColorNameDisabledButton:      color.NRGBA{R: 0xEE, G: 0xF2, B: 0xF7, A: 0xFF},
	theme.ColorNameError:               color.NRGBA{R: 0xC0, G: 0x39, B: 0x2B, A: 0xFF},
	theme.ColorNameFocus:               color.NRGBA{R: 0x1B, G: 0x4F, B: 0x9C, A: 0xFF},
	theme.ColorNameForeground:          color.NRGBA{R: 0x15, G: 0x20, B: 0x33, A: 0xFF},
	theme.ColorNameForegroundOnError:   color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	theme.ColorNameForegroundOnPrimary: color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	theme.ColorNameForegroundOnSuccess: color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	theme.ColorNameForegroundOnWarning: color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	theme.ColorNameHeaderBackground:    color.NRGBA{R: 0xE7, G: 0xEE, B: 0xF8, A: 0xFF},
	theme.ColorNameHover:               color.NRGBA{R: 0xE4, G: 0xED, B: 0xF8, A: 0xFF},
	theme.ColorNameHyperlink:           color.NRGBA{R: 0x1B, G: 0x4F, B: 0x9C, A: 0xFF},
	theme.ColorNameInputBackground:     color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	theme.ColorNameInputBorder:         color.NRGBA{R: 0xC5, G: 0xD2, B: 0xE4, A: 0xFF},
	theme.ColorNameMenuBackground:      color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	theme.ColorNameOverlayBackground:   color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	theme.ColorNamePlaceHolder:         color.NRGBA{R: 0x6D, G: 0x7B, B: 0x90, A: 0xFF},
	theme.ColorNamePressed:             color.NRGBA{R: 0xD0, G: 0xDF, B: 0xF3, A: 0xFF},
	theme.ColorNamePrimary:             color.NRGBA{R: 0x1B, G: 0x4F, B: 0x9C, A: 0xFF},
	theme.ColorNameScrollBar:           color.NRGBA{R: 0xC5, G: 0xD0, B: 0xE0, A: 0xFF},
	theme.ColorNameSelection:           color.NRGBA{R: 0xD3, G: 0xE3, B: 0xFA, A: 0xFF},
	theme.ColorNameSeparator:           color.NRGBA{R: 0xD5, G: 0xDE, B: 0xEA, A: 0xFF},
	theme.ColorNameShadow:              color.NRGBA{R: 0x15, G: 0x20, B: 0x33, A: 0x22},
	theme.ColorNameSuccess:             color.NRGBA{R: 0x0E, G: 0x7A, B: 0x45, A: 0xFF},
	theme.ColorNameWarning:             color.NRGBA{R: 0xA1, G: 0x5C, B: 0x00, A: 0xFF},
}

var darkColors = map[fyne.ThemeColorName]color.Color{
	theme.ColorNameBackground:          color.NRGBA{R: 0x10, G: 0x17, B: 0x22, A: 0xFF},
	theme.ColorNameButton:              color.NRGBA{R: 0x24, G: 0x30, B: 0x44, A: 0xFF},
	theme.ColorNameDisabled:            color.NRGBA{R: 0x6C, G: 0x78, B: 0x8C, A: 0xFF},
	theme.ColorNameDisabledButton:      color.NRGBA{R: 0x1C, G: 0x25, B: 0x33, A: 0xFF},
	theme.ColorNameError:               color.NRGBA{R: 0xFF, G: 0x8A, B: 0x80, A: 0xFF},
	theme.ColorNameFocus:               color.NRGBA{R: 0x8F, G: 0xB4, B: 0xFF, A: 0xFF},
	theme.ColorNameForeground:          color.NRGBA{R: 0xE6, G: 0xED, B: 0xF7, A: 0xFF},
	theme.ColorNameForegroundOnError:   color.NRGBA{R: 0x10, G: 0x17, B: 0x22, A: 0xFF},
	theme.ColorNameForegroundOnPrimary: color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	theme.ColorNameForegroundOnSuccess: color.NRGBA{R: 0x10, G: 0x17, B: 0x22, A: 0xFF},
	theme.ColorNameForegroundOnWarning: color.NRGBA{R: 0x10, G: 0x17, B: 0x22, A: 0xFF},
	theme.ColorNameHeaderBackground:    color.NRGBA{R: 0x1A, G: 0x24, B: 0x36, A: 0xFF},
	theme.ColorNameHover:               color.NRGBA{R: 0x24, G: 0x30, B: 0x44, A: 0xFF},
	theme.ColorNameHyperlink:           color.NRGBA{R: 0x8F, G: 0xB4, B: 0xFF, A: 0xFF},
	theme.ColorNameInputBackground:     color.NRGBA{R: 0x1A, G: 0x23, B: 0x33, A: 0xFF},
	theme.ColorNameInputBorder:         color.NRGBA{R: 0x3A, G: 0x4A, B: 0x63, A: 0xFF},
	theme.ColorNameMenuBackground:      color.NRGBA{R: 0x1A, G: 0x23, B: 0x33, A: 0xFF},
	theme.ColorNameOverlayBackground:   color.NRGBA{R: 0x1A, G: 0x23, B: 0x33, A: 0xFF},
	theme.ColorNamePlaceHolder:         color.NRGBA{R: 0x8B, G: 0x98, B: 0xAD, A: 0xFF},
	theme.ColorNamePressed:             color.NRGBA{R: 0x2C, G: 0x3C, B: 0x56, A: 0xFF},
	theme.ColorNamePrimary:             color.NRGBA{R: 0x3B, G: 0x82, B: 0xF6, A: 0xFF},
	theme.ColorNameScrollBar:           color.NRGBA{R: 0x3A, G: 0x4A, B: 0x63, A: 0xFF},
	theme.ColorNameSelection:           color.NRGBA{R: 0x24, G: 0x36, B: 0x56, A: 0xFF},
	theme.ColorNameSeparator:           color.NRGBA{R: 0x2C, G: 0x3A, B: 0x50, A: 0xFF},
	theme.ColorNameShadow:              color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x66},
	theme.ColorNameSuccess:             color.NRGBA{R: 0x3D, G: 0xDC, B: 0x97, A: 0xFF},
	theme.ColorNameWarning:             color.NRGBA{R: 0xF5, G: 0xC1, B: 0x6C, A: 0xFF},
}
