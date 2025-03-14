package ui

import (
	"context"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type Environment struct {
	Logger log.Logger

	Client search.Client

	Width, Height float32
}

func New(ctx context.Context, env Environment) fyne.App {
	a := app.New()

	app.SetMetadata(fyne.AppMetadata{
		Name:    "Watchman",
		Version: watchman.Version,
	})

	// Don't allow for Dark Mode
	a.Settings().SetTheme(newDefaultTheme(theme.LightTheme()))

	device := fyne.CurrentDevice()

	env.Logger.Debug().Logf("device: mobile=%v browser=%v keyboard=%v",
		device.IsMobile(), device.IsBrowser(), device.HasKeyboard())

	w := a.NewWindow("Hello World")
	w.SetTitle("Watchman")

	// Center the overall window and make it a reasonable size
	if device.IsBrowser() {
		// Browser-specific settings
		env.Width = 1200.0 // Slightly narrower for browser
		env.Height = 800.0
	} else {
		// Desktop settings
		env.Width = 1500.0
		env.Height = 900.0
		w.Resize(fyne.NewSize(env.Width, env.Height))
		w.CenterOnScreen()
	}
	w.Show()

	// Set app tabs along the top
	tabs := container.NewAppTabs(
		container.NewTabItem("Search", SearchContainer(ctx, env)),
		container.NewTabItem("Admin", AdminContainer(ctx, env)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)

	return a
}
