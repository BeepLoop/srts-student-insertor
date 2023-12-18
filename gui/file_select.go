package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func FileSelect(win fyne.Window) *widget.Button {
	fileSelect := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		if err != nil {
			fmt.Println("errored: ", err)
			return
		}

		if uc == nil {
			fmt.Println("no source file selected")
			return
		}

		file := uc.URI().Path()
		fmt.Println("file: ", file)
		Options.Source = file

	}, win)

	fileBtn := widget.NewButton("select a file", func() {
		fileSelect.Show()
	})

	return fileBtn
}
