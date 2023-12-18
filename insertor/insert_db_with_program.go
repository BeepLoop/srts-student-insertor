package insertor

import "fmt"

func InsertToDBWithProgram(info []string, programId int) error {

	query := `
        INSERT INTO 
            Student (
                controlNumber,
                lastname,
                firstname,
                middlename,
                fileLocation,
                programId
            )
        VALUES (
            ?, ?, ?, ?, ?, ?
        )
    `

	_, err := DB_Conn.Exec(query, info[0], info[2], info[3], info[4], info[5], programId)
	if err != nil {
		fmt.Printf("Error inserting: %v, programId: %d", info, programId)
		return err
	}

	return nil
}
