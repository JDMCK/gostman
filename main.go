package main

import (
	"bytes"
	gostpages "gostman/pages"
	"gostman/request"
	"image/jpeg"
	"os"

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
		SetText("Welcome to Gostman 👻!").
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

	// Open the image file
	file, _ := os.Open("./gosty.jpeg")
	defer file.Close()
	// Read the file into a byte slice
	imgBytes, _ := os.ReadFile("./gosty.jpeg")
	// Decode the image from bytes
	gosty, _ := jpeg.Decode(bytes.NewReader(imgBytes))
	// Assuming form is an object with an AddImage method
	image := tview.NewImage().
		SetImage(gosty).
		SetColors(100)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(image, 15, 0, false).
		AddItem(modal, 0, 4, true)

	pages.AddPage(
		"launch",
		flex,
		true,
		true,
	)
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}
