package gui

import (
	"fmt"

	"github.com/BeepLoop/registrar-student-insertor/insertor"
)

func InsertorBridge() error {
	if Options.Output == "" {
		fmt.Println("blank output file")
		Options.Output = "fails.txt"
	} else {
		fmt.Println("output file: ", Options.Output)
	}

	err := insertor.Insertor(&Options)
	if err != nil {
		return err
	}

	return nil
}
