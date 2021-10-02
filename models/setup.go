package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDataBase() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "marreta"
		dbname   = "shortener"
	)

	connInf := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", connInf)

	if err != nil {
		panic(err)
	}

	DB = db
	return db
}
