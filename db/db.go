package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectDataBase() *sql.DB {
	conn := "user=postgres dbname=loja_go password=1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
