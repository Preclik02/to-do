package main

// -- IMPORT LIBRARIES -- //
import (

	// -- fyne -- //
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	// -- general -- //
	"fmt"
	"slices"
	"os"
	"io"
	"strings"
	"strconv"

)

// -- FUNCTIONS -- //
func read_tasks() ([]string, []int) {

	// -- Opening file, checking err and closing it at the end -- //
	file, err := os.OpenFile("tasks.md", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("[+] Error while opening the file")
	}
	defer file.Close()

	// -- Setting the cursor at the begining of the file and start reading -- //
	file.Seek(0, io.SeekStart)
	var file_string string
	var marks_slice_string_read string
	fmt.Fscanf(file, "%s\n%s", &file_string, &marks_slice_string_read)
	// ^^ needs to read the 

	// -- Split the string into a slice separated words with ";" -- //
	tasks_slice := strings.Split(file_string, ";")
	marks_slice_string := strings.Split(marks_slice_string_read, ";")

	var marks_slice []int
	for _, i := range marks_slice_string {
		marks_slice = append(marks_slice, strconv.Atoi(i))
	}

	// -- return slice since it will be called like `x, y := read_tasks()` -- //
	return tasks_slice, marks_slice
}

func write_tasks(tasks_slice []string, marks_slice []int) {

	file, err := os.OpenFile("tasks.md", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("[+] Error while opening the file")
	}
	defer file.Close()

	for _, i := range tasks_slice {
		fmt.Fprintf(file, "%s;", i)
	}
	for _, i := range marks_slice {
		fmt.Fprintf(file, "%s;", i)
	}
	
	// -- make marking logic -- //



}
// ^^ FUNCTIONS ^^ //

// -- main function -- //
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
	var marks_slice []int
	tasks_slice, marks_slice = read_tasks()
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
	var selected_task = prefs.StringWithFallback("task", "none", "mark", 0)
	var marks_state []int
	tasks.OnSelected = func(id widget.ListItemID) {
		selected_task = tasks_slice[id]
		prefs.SetString("task", selected_task, "mark", mark_state[id])
		fmt.Println("selected task --> "+selected_task+"mark --> "+mark_state[id])
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
			marks_state = append(marks_state, 0)
			input.SetText("")
			prefs.SetString("task", "", "mark", 0)
			tasks.Refresh()
			// -- WRITE FUNCTION -- //
			write_tasks(tasks_slice, marks_slice)
		}
	})
	btn_remove := widget.NewButton("Remove", func() {
		fmt.Println("[+] Remove --> "+selected_task)
			
		j := 0
		for _, i := range tasks_slice {
			if i == selected_task {
				tasks_slice = slices.Delete(tasks_slice, j, j+1)
				marks_state = slices.Delete(marks_state, j, j+1)
			}
			j++
		}
		prefs.SetString("task", "", "mark", 0)
		tasks.Refresh()
		// -- WRITE FUNCTION -- //
		write_tasks(tasks_slice, marks_slice)
	})
	btn_mark := widget.NewButton("Mark Task", func() {
		fmt.Println("Mark "+selected_task)
		// -- get the index of selected_task from tasks_slice and use it on marks_state and change it -- //
		// -- it can be done with for loop -- //
		// -- WRITE FUNCTION -- //
		write_tasks(tasks_slice, marks_slice)
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
	Window.ShowAndRun()

}
