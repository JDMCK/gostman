package pages

import (
	"fmt"
	"gostman/request"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func LoadSavedRequest(app *tview.Application, pages *tview.Pages, selectedReq *request.Request) *tview.List {
	reqs, _ := request.DeserializeRequest()

	list := tview.NewList()

	list.SetInputCapture(func(key *tcell.EventKey) *tcell.EventKey {
		if key.Key() == tcell.KeyEsc {
			pages.SwitchToPage("launch")
			return nil
		}
		return key
	})

	for i := len(reqs) - 1; i >= 0; i-- {
		req := reqs[i]
		list.AddItem(req.URL, string(req.Method), 0, func() {
			*selectedReq = req
			pages.AddAndSwitchToPage("Execute Request", ExecutePage(app, pages, selectedReq), true)
		})
	}

	return list
}

func ExecutePage(app *tview.Application, pages *tview.Pages, req *request.Request) *tview.Flex {
	headersFormatted := ""
	for key, value := range req.Headers {
		headersFormatted += fmt.Sprintf("  %s: %s\n", key, value)
	}

	requestDetails := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetText(fmt.Sprintf(
			"URL: %s\nMethod: %s\nHeaders:\n%sBody:\n%s",
			req.URL, req.Method, headersFormatted, req.Body,
		))

	requestDetails.SetBorder(true)
	requestDetails.SetTitle("Request Details")

	// Results display
	results := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetText("Results will appear here...")

	results.SetBorder(true)

	buttons := tview.NewForm().
		SetButtonBackgroundColor(tcell.ColorNavy).
		SetHorizontal(true).
		AddButton(
			"Execute Request", func() {
				results.SetText("[green]Executing request...[/green]\n")

				go func() { // Run request execution in a goroutine
					resp, err := request.ExecuteRequest(req)
					if err != nil {
						app.QueueUpdateDraw(func() {
							results.SetText(fmt.Sprintf("[red]Error: %s\n", err.Error()))
						})
						return
					}

					// Format response headers
					headersFormatted := ""
					for key, value := range resp.Headers {
						headersFormatted += fmt.Sprintf("  %s: %s\n", key, value)
					}

					// Construct the full response message
					responseText := fmt.Sprintf(
						"[green]Response received: [yellow]%d\n\n"+
							"[blue]Headers:\n%s\n"+
							"[blue]Body:\n%s",
						resp.StatusCode, headersFormatted, resp.Body,
					)

					// Update UI safely
					app.QueueUpdateDraw(func() {
						results.SetText(responseText)
					})
				}()
			},
		).
		AddButton(
			"Go Back", func() {
				pages.SwitchToPage("launch")
			},
		)

	// Main layout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(requestDetails, 0, 2, false).
		AddItem(buttons, 0, 1, true).
		AddItem(results, 0, 4, false)

	flex.SetInputCapture(func(key *tcell.EventKey) *tcell.EventKey {
		if key.Key() == tcell.KeyEsc {
			pages.SwitchToPage("launch")
			return nil
		}
		return key
	})

	return flex
}
