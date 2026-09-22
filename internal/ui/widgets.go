package ui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func sectionLabel(text string) *widget.Label {
	label := widget.NewLabel(text)
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.SizeName = theme.SizeNameSubHeadingText
	return label
}

func wordWrappingLabel(text string) *widget.Label {
	label := widget.NewLabel(text)
	label.Wrapping = fyne.TextWrapWord
	label.Selectable = true
	return label
}

func formatCount(n int) string {
	negative := n < 0
	if negative {
		n = -n
	}
	digits := strconv.Itoa(n)
	if len(digits) <= 3 {
		if negative {
			return "-" + digits
		}
		return digits
	}

	var b strings.Builder
	if negative {
		b.WriteByte('-')
	}
	lead := len(digits) % 3
	if lead == 0 {
		lead = 3
	}
	b.WriteString(digits[:lead])
	for i := lead; i < len(digits); i += 3 {
		b.WriteByte(',')
		b.WriteString(digits[i : i+3])
	}
	return b.String()
}

func formatElapsed(d time.Duration) string {
	switch {
	case d < time.Millisecond:
		return d.Truncate(time.Microsecond).String()
	case d < time.Second:
		return d.Truncate(time.Millisecond).String()
	default:
		return d.Truncate(10 * time.Millisecond).String()
	}
}

func formatMatch(score float64) string {
	return fmt.Sprintf("%.1f%%", score*100)
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return "—"
	}
	return t.Local().Format("2006-01-02 15:04 MST")
}

// matchImportance colors a score the way a reviewer would read it:
// strong matches in green, review-range matches in amber, weak ones in red.
func matchImportance(score float64) widget.Importance {
	switch {
	case score >= 0.90:
		return widget.SuccessImportance
	case score >= 0.70:
		return widget.WarningImportance
	default:
		return widget.DangerImportance
	}
}
