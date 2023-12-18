package insertor

import (
	"fmt"
	"os"
)

func WriteToFile(file *os.File, info []string, err error) error {
	text := fmt.Sprintf("%s, %s, %s, %s, %s, reason=%s  \n", info[0], info[2], info[3], info[4], info[5], err.Error())
	if _, err := file.WriteString(text); err != nil {
		return err
	}

	return nil
}
