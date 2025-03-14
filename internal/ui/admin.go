package ui

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AdminContainer(ctx context.Context, env Environment) fyne.CanvasObject {
	// Create the main container
	mainContainer := container.NewVBox()

	// Create header
	header := widget.NewLabelWithStyle("List Information", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	mainContainer.Add(header)

	// Create status labels (will be updated with data)
	versionLabel := widget.NewLabel("Version: Loading...")
	startedLabel := widget.NewLabel("Started: -")
	endedLabel := widget.NewLabel("Ended: -")

	// Status container
	statusContainer := container.NewVBox(
		versionLabel,
		startedLabel,
		endedLabel,
	)

	// Create a container for the lists table
	listsLabel := widget.NewLabelWithStyle("Available Lists", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Create table with headers
	table := widget.NewTable(
		func() (int, int) {
			return 1, 3 // Start with header row only, 3 columns
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			label.SetText("Loading...")

			// Set header row styling
			if i.Row == 0 {
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.Alignment = fyne.TextAlignCenter
			} else {
				label.TextStyle = fyne.TextStyle{}

				// Align count to right, others to left
				if i.Col == 1 { // Count column
					label.Alignment = fyne.TextAlignTrailing
				} else {
					label.Alignment = fyne.TextAlignLeading
				}
			}
		},
	)

	// Set column widths
	table.SetColumnWidth(0, 100) // List name
	table.SetColumnWidth(1, 80)  // Count
	table.SetColumnWidth(2, 600) // Hash

	listsScroll := container.NewScroll(table)
	listsScroll.SetMinSize(fyne.NewSize(0, env.Height*0.4))

	listsContainer := container.NewBorder(listsLabel, nil, nil, nil, listsScroll)

	// Create refresh button
	refreshBtn := widget.NewButton("Refresh Data", func() {
		updateListInfo(env, ctx, versionLabel, startedLabel, endedLabel, listsScroll)
	})

	// Add components to main container
	mainContainer.Add(refreshBtn)
	mainContainer.Add(statusContainer)
	mainContainer.Add(listsContainer)

	// Initial data fetch
	updateListInfo(env, ctx, versionLabel, startedLabel, endedLabel, listsScroll)

	return container.NewPadded(mainContainer)
}

// updateListInfo fetches the list info and updates the UI
func updateListInfo(env Environment, ctx context.Context, versionLabel, startedLabel, endedLabel *widget.Label, listsContent fyne.CanvasObject) {
	listInfo, err := env.Client.ListInfo(ctx)
	if err != nil {
		versionLabel.SetText(fmt.Sprintf("Error: %s", err))
		return
	}

	// Update version and timestamps
	versionLabel.SetText(fmt.Sprintf("Version: %s", listInfo.Version))
	startedLabel.SetText(fmt.Sprintf("Started: %s", formatTime(listInfo.StartedAt)))
	endedLabel.SetText(fmt.Sprintf("Ended: %s", formatTime(listInfo.EndedAt)))

	// Get the table widget
	table := listsContent.(*container.Scroll).Content.(*widget.Table)

	// Create slice of list entries for sorting (optional)
	type listEntry struct {
		name  string
		count int
		hash  string
	}

	entries := make([]listEntry, 0, len(listInfo.Lists))
	for listName, count := range listInfo.Lists {
		hash := listInfo.ListHashes[listName]
		entries = append(entries, listEntry{
			name:  listName,
			count: count,
			hash:  hash,
		})
	}
	slices.SortFunc(entries, func(e1, e2 listEntry) int {
		return -1 * cmp.Compare(e1.count, e2.count) // DESC
	})

	// Resize table based on data
	table.CreateCell = func() fyne.CanvasObject {
		return widget.NewLabel("Template")
	}

	// Update table data
	if len(entries) == 0 {
		// Only show headers and "No lists" message
		table.Length = func() (int, int) {
			return 2, 3
		}

		table.UpdateCell = func(id widget.TableCellID, object fyne.CanvasObject) {
			label := object.(*widget.Label)

			// Headers
			if id.Row == 0 {
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.Alignment = fyne.TextAlignCenter

				switch id.Col {
				case 0:
					label.SetText("List Name")
				case 1:
					label.SetText("Count")
				case 2:
					label.SetText("Hash")
				}
			} else if id.Row == 1 {
				// No lists message, centered across all columns
				label.Alignment = fyne.TextAlignCenter
				label.TextStyle = fyne.TextStyle{}

				if id.Col == 0 {
					label.SetText("No lists available")
				} else {
					label.SetText("")
				}
			}
		}
	} else {
		// Show headers and data
		table.Length = func() (int, int) {
			return len(entries) + 1, 3 // +1 for header row
		}

		table.UpdateCell = func(id widget.TableCellID, object fyne.CanvasObject) {
			label := object.(*widget.Label)

			// Headers
			if id.Row == 0 {
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.Alignment = fyne.TextAlignCenter

				switch id.Col {
				case 0:
					label.SetText("List Name")
				case 1:
					label.SetText("Count")
				case 2:
					label.SetText("Hash")
				}
			} else {
				// Data rows
				entry := entries[id.Row-1]
				label.TextStyle = fyne.TextStyle{}

				switch id.Col {
				case 0:
					label.Alignment = fyne.TextAlignLeading
					label.SetText(entry.name)
				case 1:
					label.Alignment = fyne.TextAlignTrailing
					label.SetText(fmt.Sprintf("%d", entry.count))
				case 2:
					label.Alignment = fyne.TextAlignLeading
					label.SetText(entry.hash)
				}
			}
		}
	}

	// Refresh the table
	table.Refresh()
}

// formatTime returns a user-friendly time string
func formatTime(t time.Time) string {
	if t.IsZero() {
		return "N/A"
	}
	return t.Format(time.RFC3339)
}
