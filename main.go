package main

import (
	gostpages "gostman/pages"
	"gostman/request"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	pages := tview.NewPages()
	var selectedReq request.Request

	var modal = tview.NewModal().
		SetBackgroundColor(tcell.ColorBlack).
		SetButtonBackgroundColor(tcell.ColorNavy).
		SetText("Welcome to Gostman👻!").
		AddButtons([]string{"Create New Request", "Load Saved Request"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonIndex == 0 {
				pages.AddAndSwitchToPage(
					"new request",
					tview.NewFlex().
						SetDirection(tview.FlexColumn).
						AddItem(
							gostpages.NewRequestForm(app, pages),
							125, 1, true,
						),
					true,
				)
			} else if buttonIndex == 1 {
				pages.AddAndSwitchToPage(
					"load saved request",
					gostpages.LoadSavedRequest(app, pages, &selectedReq),
					true,
				)
			} else {
				app.Stop()
			}
		})
	modal.Box.SetBackgroundColor(tcell.ColorBlack)
	pages.AddPage(
		"launch",
		modal,
		true,
		true,
	)
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}
