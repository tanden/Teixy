package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const Driver = "mysql"
const ConnInfoTeixyBooks = "teixy:teixy@tcp(db_books:3306)/teixy_books"

//Create mysql connection
//to teixy_article database
func CreateConectionTeixyBooks() *sqlx.DB {
	db, err := sqlx.Connect(Driver, ConnInfoTeixyBooks)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
