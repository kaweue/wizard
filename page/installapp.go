package page

import (
	"github.com/rivo/tview"
)

// NewInstallApp is used to create a page that collect information from user whether App should be installed
func NewInstallApp(done func(installApp bool)) tview.Primitive {
	m := tview.NewModal().
		SetText("Would you like to install app").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(index int, label string) {
			if label == "Yes" {
				done(true)
			} else {
				done(false)
			}
		})
	f := tview.NewFlex().AddItem(m, 1, 1, true)
	f.SetBorder(true).SetTitle("Installation wizard")
	return f
}
