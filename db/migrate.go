package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	command = flag.String("exec", "", "set up or down as a argument")
)

func main() {

	flag.Parse()
	if len(*command) < 1 {
		fmt.Println("\nerror: no argument\n")
		showUsageMessge()
		return
	}
	if *command != "up" || *command != "down" {
		fmt.Println("\nerror: invalid command\n")
		showUsageMessge()
		return
	}

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

func showUsageMessge() {
	fmt.Println("-------------------------------------")
	fmt.Println("Usage")
	fmt.Println(" go run migrate.go -exec <command>\n")
	fmt.Println("Available Commands: ")
	fmt.Println("  up    execute up sqls")
	fmt.Println("  down  execute down sqls")
	fmt.Println("-------------------------------------")
}
