package ui

import (
	"cmp"
	"context"
	"image/color"
	"slices"

	"github.com/moov-io/watchman/pkg/search"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func AdminContainer(ctx context.Context, env Environment) fyne.CanvasObject {
	versionValue := statValue("—")
	startedValue := statValue("—")
	endedValue := statValue("—")
	totalValue := statValue("—")

	status := widget.NewLabel("Loading lists…")
	status.Wrapping = fyne.TextWrapWord
	status.Importance = widget.LowImportance

	var entries []listEntry
	table := widget.NewTable(
		func() (int, int) {
			return len(entries) + 1, 3
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("template")
			label.Truncation = fyne.TextTruncateEllipsis
			return label
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			label.TextStyle = fyne.TextStyle{}
			label.Importance = widget.MediumImportance
			label.Selectable = false
			label.Truncation = fyne.TextTruncateEllipsis
			if id.Row == 0 {
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.Alignment = fyne.TextAlignLeading
				switch id.Col {
				case 0:
					label.SetText("List")
				case 1:
					label.SetText("Entities")
					label.Alignment = fyne.TextAlignTrailing
				case 2:
					label.SetText("Hash")
				}
				return
			}
			if id.Row-1 >= len(entries) {
				label.SetText("")
				return
			}
			entry := entries[id.Row-1]
			switch id.Col {
			case 0:
				label.Alignment = fyne.TextAlignLeading
				label.SetText(entry.name)
			case 1:
				label.Alignment = fyne.TextAlignTrailing
				label.SetText(formatCount(entry.count))
			default:
				label.Alignment = fyne.TextAlignLeading
				label.TextStyle = fyne.TextStyle{Monospace: true}
				label.Selectable = true
				label.Truncation = fyne.TextTruncateOff
				if entry.hash == "" {
					label.SetText("—")
				} else {
					label.SetText(entry.hash)
				}
			}
		},
	)
	table.SetColumnWidth(0, 160)
	table.SetColumnWidth(1, 120)
	table.SetColumnWidth(2, 560)

	// The table scrolls itself. An outer scroll offsets the header from the rows.

	var refreshBtn *widget.Button
	load := func() {
		refreshBtn.Disable()
		status.Importance = widget.LowImportance
		status.SetText("Loading lists…")
		status.Show()

		go func() {
			info, err := env.Client.ListInfo(ctx)
			fyne.Do(func() {
				refreshBtn.Enable()
				if err != nil {
					versionValue.SetText("—")
					startedValue.SetText("—")
					endedValue.SetText("—")
					totalValue.SetText("—")
					entries = nil
					table.Refresh()
					table.Hide()
					status.Importance = widget.DangerImportance
					status.SetText(err.Error())
					status.Show()
					return
				}
				applyListInfo(info, versionValue, startedValue, endedValue, totalValue, &entries)
				table.Refresh()
				if len(entries) == 0 {
					table.Hide()
					status.Importance = widget.LowImportance
					status.SetText("No lists are loaded.")
					status.Show()
					return
				}
				status.Hide()
				table.Show()
			})
		}()
	}

	refreshBtn = widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), load)
	table.Hide()
	load()

	stats := container.NewHBox(
		statBlock("Version", versionValue),
		statGap(),
		statBlock("Started", startedValue),
		statGap(),
		statBlock("Finished", endedValue),
		statGap(),
		statBlock("Entities", totalValue),
	)

	header := container.NewVBox(
		container.NewBorder(nil, nil, sectionLabel("Loaded lists"), refreshBtn),
		stats,
		status,
		statGap(),
	)

	return container.NewPadded(container.NewBorder(header, nil, nil, nil, table))
}

func statGap() fyne.CanvasObject {
	gap := canvas.NewRectangle(color.Transparent)
	gap.SetMinSize(fyne.NewSize(36, 8))
	return gap
}

type listEntry struct {
	name  string
	count int
	hash  string
}

func applyListInfo(
	info search.ListInfoResponse,
	versionValue, startedValue, endedValue, totalValue *widget.Label,
	entries *[]listEntry,
) {
	versionValue.SetText(cmp.Or(info.Version, "—"))
	startedValue.SetText(formatTime(info.StartedAt))
	endedValue.SetText(formatTime(info.EndedAt))

	next := make([]listEntry, 0, len(info.Lists))
	total := 0
	for name, count := range info.Lists {
		total += count
		next = append(next, listEntry{
			name:  name,
			count: count,
			hash:  info.ListHashes[name],
		})
	}
	slices.SortFunc(next, func(a, b listEntry) int {
		return -1 * cmp.Compare(a.count, b.count)
	})
	*entries = next
	totalValue.SetText(formatCount(total))
}

func statBlock(title string, value *widget.Label) fyne.CanvasObject {
	caption := widget.NewLabel(title)
	caption.Importance = widget.LowImportance
	return container.NewVBox(caption, value)
}

func statValue(text string) *widget.Label {
	label := widget.NewLabel(text)
	label.TextStyle = fyne.TextStyle{Bold: true}
	return label
}
