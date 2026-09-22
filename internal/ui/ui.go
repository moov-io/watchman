package ui

import (
	"cmp"
	"context"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Environment struct {
	Logger log.Logger

	Client search.Client

	Width, Height float32
}

func New(ctx context.Context, env Environment) fyne.App {
	a := app.New()

	// fyne package --app-version sets Metadata.Version. The server build
	// sets watchman.Version with -ldflags. Keep whichever one is present.
	version := cmp.Or(watchman.Version, a.Metadata().Version)
	app.SetMetadata(fyne.AppMetadata{
		Name:    "Watchman",
		Version: version,
	})

	a.Settings().SetTheme(newWatchmanTheme())

	device := fyne.CurrentDevice()
	env.Logger.Debug().Logf("device: mobile=%v browser=%v keyboard=%v",
		device.IsMobile(), device.IsBrowser(), device.HasKeyboard())

	w := a.NewWindow("Watchman")

	if device.IsBrowser() {
		env.Width = 1200
		env.Height = 800
	} else {
		env.Width = 1500
		env.Height = 900
		w.Resize(fyne.NewSize(env.Width, env.Height))
		w.CenterOnScreen()
	}

	w.SetContent(shell(ctx, env))
	w.Show()

	return a
}

func shell(ctx context.Context, env Environment) fyne.CanvasObject {
	title := widget.NewLabel("Watchman")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.SizeName = theme.SizeNameHeadingText

	subtitle := widget.NewLabel("Sanctions and watchlist screening")
	subtitle.Importance = widget.LowImportance

	version := widget.NewLabel(fyne.CurrentApp().Metadata().Version)
	version.Importance = widget.LowImportance
	version.Alignment = fyne.TextAlignTrailing
	if version.Text == "" {
		version.Hide()
	}

	header := container.NewBorder(
		nil,
		widget.NewSeparator(),
		container.NewVBox(title, subtitle),
		version,
	)

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Search", theme.SearchIcon(), SearchContainer(ctx, env)),
		container.NewTabItemWithIcon("Lists", theme.ListIcon(), AdminContainer(ctx, env)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	return container.NewBorder(container.NewPadded(header), nil, nil, nil, tabs)
}
