package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

const Driver = "mysql"
const ConnInfoTeixyArticles = "teixy:teixy@tcp(db_articles:3306)/teixy_articles"

//Create mysql connection
//to teixy_article database
func CreateConectionTeixyArticles() *sqlx.DB {
	db, err := sqlx.Connect(Driver, ConnInfoTeixyArticles)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
