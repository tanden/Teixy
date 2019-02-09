package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	command = flag.String("exec", "no command", "set up or down as a command to exec sql")
)

func main() {

	m, err := migrate.New(
		"file://sql/",
		"mysql://teixy:teixy@tcp(0.0.0.0:3306)/teixy_article")

	if err != nil {
		fmt.Println("err", err)
	}

	version, dirty, err := m.Version()
	fmt.Println(version)
	fmt.Println(dirty)
	fmt.Println(err)

	flag.Parse()
	if *command == "up" {
		fmt.Println("command: exec up")
		err := m.Up()
		if err != nil {
			fmt.Println("err", err)
		}
	}

	if *command == "down" {
		fmt.Println("command: exec down")
		err := m.Down()
		if err != nil {
			fmt.Println("err", err)
		}
	}
}
