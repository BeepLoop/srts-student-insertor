package main

import (
	"fmt"
	"os"
)

func TrailingCommaRemover(str string) string {
	if len(str) > 0 && str[len(str)-1] == ',' {
		return str[:len(str)-1]
	}

	file, err := os.OpenFile("with-trailing-comma.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("failed to open file")
		return str
	}
	defer file.Close()

	_, err = file.WriteString(str + "\n")
	if err != nil {
		fmt.Println("failed to write to file str: ", str)
		return str
	}

	return str
}
