package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

//Create mysql connection
//to teixy_article database
func CreateConectionArticle() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "teixy:teixy@tcp(db_article:3306)/teixy_article")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
