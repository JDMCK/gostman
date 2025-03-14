package pages

import (
	"fmt"
	"gostman/request"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type FormState struct {
	HTTPMethod request.HTTPMethod
	URL        string
	Headers    map[int]([2]string)
	Body       string
}

// NewRequestForm initializes the form
func NewRequestForm(app *tview.Application, pages *tview.Pages) *tview.Form {
	state := &FormState{
		HTTPMethod: "GET",
		URL:        "",
		Headers:    make(map[int][2]string),
		Body:       "",
	}

	form := tview.NewForm().
		SetFieldBackgroundColor(tcell.ColorNavy).
		SetButtonBackgroundColor(tcell.ColorNavy).
		AddDropDown("HTTP Method:", []string{"GET", "POST", "PUT", "DELETE"}, 0,
			func(option string, optionIndex int) {
				state.HTTPMethod = request.HTTPMethod(option)
			},
		).
		AddInputField("URL:", "", 100, nil,
			func(text string) {
				state.URL = text
			},
		).
		AddTextArea("Body", "", 50, 30, 200, func(text string) {
			state.Body = text
		})

	AddMapInput("Header", app, pages, form, state, 1)
	form.SetBorder(true).SetTitle("Create a New Request👻")

	form.SetInputCapture(func(key *tcell.EventKey) *tcell.EventKey {
		if key.Key() == tcell.KeyEsc {
			pages.SwitchToPage("launch")
			return nil
		}
		return key
	})

	return form
}

// AddMapInput adds a new key-value pair input field with remove functionality
func AddMapInput(
	label string,
	app *tview.Application,
	pages *tview.Pages,
	form *tview.Form,
	state *FormState,
	index int,
) {
	form.
		AddInputField(fmt.Sprintf("%s %d name:", label, index), "", 35, nil,
			func(text string) {
				temp := (*state).Headers[index]
				temp[0] = text
				(*state).Headers[index] = temp
			},
		).
		AddInputField(fmt.Sprintf("%s %d value:", label, index), "", 35, nil,
			func(text string) {
				temp := (*state).Headers[index]
				temp[1] = text
				(*state).Headers[index] = temp
			},
		)
	AddButtons(label, app, pages, form, state, index)
}

func AddButtons(
	label string,
	app *tview.Application,
	pages *tview.Pages,
	form *tview.Form,
	state *FormState,
	index int,
) {
	form.
		AddButton("+", func() {
			form.ClearButtons()

			AddMapInput(label, app, pages, form, state, index+1)
			// Move focus to the newly added input field
			app.SetFocus(form)
			app.SetFocus(form.GetFormItem(3 + index*2))
		}).
		AddButton("-", func() {
			if form.GetFormItemCount() <= 5 {
				return
			}
			form.ClearButtons()
			delete((*state).Headers, index)
			RemoveHeaderField(app, form, index) // Remove from form
			index--
			AddButtons(label, app, pages, form, state, index)
		}).
		AddButton("Save", func() {
			headers := make(map[string]string)
			for _, pair := range (*state).Headers {
				headers[pair[0]] = pair[1]
			}
			req := request.NewRequest(state.HTTPMethod, state.URL, headers, state.Body)
			request.SerializeRequest(req)
			pages.SwitchToPage("launch")
		}).
		AddButton("Cancel", func() {
			pages.SwitchToPage("launch")
		})
}

func RemoveHeaderField(app *tview.Application, form *tview.Form, index int) {
	form.RemoveFormItem(1 + index*2) // Header Name
	form.RemoveFormItem(1 + index*2) // Header Value

	app.SetFocus(form)
	app.SetFocus(form.GetFormItem(index * 2))
}
