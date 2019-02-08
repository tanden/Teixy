package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"	
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://sql/",
		"mysql://teixy:teixy@tcp(0.0.0.0:3306)/teixy_article")
	if err != nil {
		fmt.Println("err", err)
	}
	m.Steps(-1)
}
