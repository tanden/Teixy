package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

const Driver = "mysql"
const ConnInfoTeixyArticle = "teixy:teixy@tcp(db_article:3306)/teixy_article"

//Create mysql connection
//to teixy_article database
func CreateConectionArticle() *sqlx.DB {
	db, err := sqlx.Connect(Driver, ConnInfoTeixyArticle)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
