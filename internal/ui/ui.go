package ui

import (
	"context"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Environment struct {
	Logger log.Logger

	Client search.Client

	Width, Height float32
}

func New(ctx context.Context, env Environment) fyne.App {
	a := app.New()

	device := fyne.CurrentDevice()
	env.Logger.Debug().Logf("device: mobile=%v browser=%v keyboard=%v",
		device.IsMobile(), device.IsBrowser(), device.HasKeyboard())

	w := a.NewWindow("Hello World")
	w.SetTitle("Watchman")

	// w.IsMobile() bool
	// w.IsBrowser() bool
	// w.SetFullScreen(bool)

	// Center the overall window and make it a reasonable size
	env.Width = 800.0
	env.Height = 500.0
	w.Resize(fyne.NewSize(env.Width, env.Height))
	w.Show()
	w.CenterOnScreen()

	// Set app tabs along the top
	tabs := container.NewAppTabs(
		container.NewTabItem("Search", SearchContainer(ctx, env)),
		container.NewTabItem("Admin", widget.NewLabel("TODO - admin operations")),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)

	return a
}
