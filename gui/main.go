package gui

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/BeepLoop/our-srts-student-insertor/insertor"
)

const (
	WIDTH  = 600
	HEIGHT = 500
)

var Options insertor.Config

func GUI() fyne.App {
	app := app.New()
	win := app.NewWindow("Student Insertor")
	win.Resize(fyne.NewSize(WIDTH, HEIGHT))

	loadingBar := widget.NewProgressBarInfinite()
	loadingBar.Hide()

	fileBtn := FileSelect(win)

	outputFile := widget.NewEntry()
	outputFile.SetPlaceHolder("defaults to fails.txt if non given")

	databaseDsn := widget.NewEntry()

	rowsAmount := widget.NewEntry()
	rowsAmount.SetText("10")
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
		OnSubmit: func() {
			amount, err := strconv.Atoi(rowsAmount.Text)
			if err != nil {
				log.Fatal("err: ", err)
			}
			Options.Output = outputFile.Text
			Options.Database = databaseDsn.Text
			Options.Limit = limitInsert.Checked
			Options.LimitValue = amount

			loadingBar.Show()

			if err := InsertorBridge(); err != nil {
				loadingBar.Hide()
				errPrompt := dialog.NewError(err, win)
				errPrompt.Show()
			} else {
				loadingBar.Hide()
				succPrompt := dialog.NewInformation("Success", "Done inserting students", win)
				succPrompt.Show()
			}

		},
	}

	content := container.NewVBox(form, loadingBar)
	win.SetContent(content)
	win.Show()

	return app
}
