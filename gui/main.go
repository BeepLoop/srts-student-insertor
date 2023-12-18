package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

	content := container.NewVBox(fileBtn)
	win.SetContent(content)
	win.Show()

	return app
}
