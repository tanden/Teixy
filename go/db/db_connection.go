package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const Driver = "mysql"
const ConnInfoBooks = "teixy:teixy@tcp(db_books:3306)/books"

//Create mysql connection
func CreateConectionTeixyBooks() *sqlx.DB {
	db, err := sqlx.Connect(Driver, ConnInfoBooks)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
