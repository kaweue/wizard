package main

import (
	"github.com/kaweue/wizard/install"
	"github.com/kaweue/wizard/page"
	"github.com/rivo/tview"
)

const (
	installAppPage = "Install"
	progressPage   = "Installation progress"
)

func main() {
	progressCh := make(chan int)
	app := tview.NewApplication()
	pages := tview.NewPages()

	installAppDone := func(installApp bool) {
		if !installApp {
			app.Stop()
			return
		}
		install.Install(progressCh)
		pages.SwitchToPage(progressPage)
	}

	pages.AddPage(installAppPage, page.NewInstallApp(installAppDone), true, true)
	pages.AddPage(progressPage, page.NewProgress(app, progressCh, func() { app.Stop() }), true, false)

	// TODO: Add next pages

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
