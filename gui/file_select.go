package gui

import (
	"fmt"
	"io"
	"os"

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

		file := uc.URI().Path()
		fmt.Println("file: ", file)
		src, err := os.Open(file)
		if err != nil {
			fmt.Println("error open file: ", err)
			return
		}

		bytes, err := io.ReadAll(src)
		if err != nil {
			fmt.Println("error reading file: ", err)
			return
		}

		fmt.Println(string(bytes))
	}, win)

	fileBtn := widget.NewButton("selec file", func() {
		fileSelect.Show()
	})

	return fileBtn
}
