package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	WIDTH  = 600
	HEIGHT = 500
)

func GUI() fyne.App {
	app := app.New()
	win := app.NewWindow("Student Insertor")
	win.Resize(fyne.NewSize(WIDTH, HEIGHT))

	fileBtn := FileSelect(win)

	outputFile := widget.NewEntry()
	outputFile.SetPlaceHolder("defaults to fails.txt if non given")

	databaseDsn := widget.NewEntry()

	rowsAmount := widget.NewEntry()
    rowsAmount.SetPlaceHolder("10")
	rowsAmount.Disable()
	limitInsert := widget.NewCheck("Limit rows to insert", func(b bool) {
		if b {
			rowsAmount.Enable()
		} else {
			rowsAmount.Disable()
		}
	})

	limitOpts := container.NewHBox(limitInsert, rowsAmount)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Source File", Widget: fileBtn},
			{Text: "Output File", Widget: outputFile},
			{Text: "Database DSN", Widget: databaseDsn},
			{Text: "Limit rows to insert", Widget: limitOpts},
		},
	}

	content := container.NewVBox(form)
	win.SetContent(content)
	win.Show()

	return app
}
