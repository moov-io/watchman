package ui

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AdminContainer(ctx context.Context, env Environment) fyne.CanvasObject {
	mainContainer := container.NewVBox()

	mainContainer.Add(widget.NewLabel("TODO - admin operations"))

	return mainContainer
}
