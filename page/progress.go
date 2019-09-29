package page

import (
	"fmt"

	"github.com/rivo/tview"
)

type updater interface {
	QueueUpdateDraw(f func()) *tview.Application
}

const installProgress = "Installation progress ( %d %%)"

// NewProgress is used to create a page that shows installation progress
func NewProgress(u updater, progress <-chan int, done func()) tview.Primitive {
	m := tview.NewModal().SetText(fmt.Sprintf(installProgress, 0))
	f := tview.NewFlex().AddItem(m, 1, 1, true)
	f.SetBorder(true).SetTitle("Installation progress")

	go func() {
		for p := range progress {
			u.QueueUpdateDraw(func() {
				m.SetText(fmt.Sprintf(installProgress, p))
			})
		}
		done()
	}()
	return f
}
