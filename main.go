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
	app := app.New()
	Window := app.NewWindow("To-Do Application") // --> title // 

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
		fmt.Println("[+] Remove --> "+input.Text)
			
		j := 0
		for _, i := range tasks_slice {
			if i == input.Text {
				tasks_slice = slices.Delete(tasks_slice, j, j+1)
			}
			j++
		}
		input.SetText("")
		tasks.Refresh()
	})
	btn_mark := widget.NewButton("Mark Task", func() {
		fmt.Println("Mark "+input.Text)
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
