package utils

import (
	"database/sql"
	"log"
)

func ConnectDB() {
	conn := "https://myuser:mypassword@magicfights-postgres-1:5432/magic_fights?sslmode=disable"
	db, err := sql.Open("postgres", conn)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
}
