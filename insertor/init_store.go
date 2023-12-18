package insertor

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB_Conn *sqlx.DB

func InitStore(db string) error {

	conn, err := sqlx.Connect("mysql", db)
	if err != nil {
		return err
	}

	DB_Conn = conn

	return nil
}
