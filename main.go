package main

import (

	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"slices"
	
)

func main() {

	// -- intilize app -- //
	app := app.NewWithID("jaylub.com")
	Window := app.NewWindow("To-Do Application") // --> title // 
	prefs := app.Preferences()

	// -- Text -- //
	label := widget.NewLabel("To-Do List")
	// ^^ Text ^^ //

	// -- Tasks -- //
	var tasks_slice []string
	tasks := widget.NewList(
		func() int {
			return len(tasks_slice)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.ListItemID, task fyne.CanvasObject) {
			task.(*widget.Label).SetText(tasks_slice[id])
		},

	)
	// ^^ NewList logic ^^ //
	var selected_task = prefs.StringWithFallback("task", "none")
	tasks.OnSelected = func(id widget.ListItemID) {
		selected_task = tasks_slice[id]
		prefs.SetString("task", selected_task)
		fmt.Println("selected task --> "+selected_task)
	}
	// ^^ on selected ^^ //
	// ^^ Tasks ^^ //

	// -- Entry -- //
	input := widget.NewEntry()
	input.SetPlaceHolder("Task name . . .")
	// ^^ Entry ^^ //
	

	// -- Buttons -- //
	btn_add := widget.NewButton("Add", func() {
		if input.Text != "" {
			fmt.Println("[+] Add --> "+input.Text)
			tasks_slice = append(tasks_slice, input.Text)
			input.SetText("")
			tasks.Refresh()
		}
	})
	btn_remove := widget.NewButton("Remove", func() {
		fmt.Println("[+] Remove --> "+selected_task)
			
		j := 0
		for _, i := range tasks_slice {
			if i == selected_task {
				tasks_slice = slices.Delete(tasks_slice, j, j+1)
			}
			j++
		}
		input.SetText("")
		tasks.Refresh()
	})
	btn_mark := widget.NewButton("Mark Task", func() {
		fmt.Println("Mark "+selected_task)
	})
	btn_row := container.NewGridWithColumns(2, btn_add, btn_remove)
	// ^^ Buttons ^^ //

	interactive_ui := container.NewBorder(
		container.NewVBox(label, input, btn_row, btn_mark, widget.NewSeparator(),),
		nil,
		nil,
		nil,
		tasks,
	)

	Window.SetContent(interactive_ui)
	Window.Resize(fyne.NewSize(400, 600))
	Window.ShowAndRun()

}
