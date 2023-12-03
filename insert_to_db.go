package main

import (
	"errors"
)

func InsertToDB(info []string) error {

	query := `
        INSERT INTO 
            Student (
                controlNumber,
                lastname,
                firstname,
                middlename,
                fileLocation
            )
        VALUES (
            ?, ?, ?, ?, ?
        )
    `

	if info[0] == "" {
		return errors.New("control number is empty")
	}

	_, err := DB_Conn.Exec(query, info[0], info[2], info[3], info[4], info[5])
	if err != nil {
		return err
	}

	return nil
}
