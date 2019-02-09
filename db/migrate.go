package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//declare command line argumet
var (
	command = flag.String("exec", "", "set up or down as a argument")
	fix     = flag.Bool("fix", false, "force exec fixed sql")
)

//available command list
var available_exec_commands = map[string]string{
	"up":   "Execute up sqls",
	"down": "Execute down sqls",
}

func main() {

	flag.Parse()
	if len(*command) < 1 {
		fmt.Println("\nerror: no argument\n")
		showUsageMessge()
		return
	}
	if len(available_exec_commands[*command]) < 1 {
		fmt.Println("\nerror: invalid command '" + *command + "'\n")
		showUsageMessge()
		return
	}

	m, err := migrate.New(
		"file://./sql/",
		"mysql://teixy:teixy@tcp(0.0.0.0:3306)/teixy_article")
	if err != nil {
		fmt.Println("err", err)
	}

	version, dirty, err := m.Version()
	fmt.Println("version: ", version)
	fmt.Println("dirty: ", dirty)
	fmt.Println("err: ", err)

	if *command == "up" {
		fmt.Println("command: exec up")
		if dirty && *fix {
			fmt.Println("fix=true: force execute current version sql")
			m.Force(int(version))
		}
		err := m.Up()
		if err != nil {
			fmt.Println("err", err)
		}
	}

	if *command == "down" {
		fmt.Println("command: exec down")
		if dirty && *fix {
			fmt.Println("fix=true: force execute current version sql")
			m.Force(int(version))
		}
		err := m.Down()
		if err != nil {
			fmt.Println("err", err)
		}
	}
}

func showUsageMessge() {
	fmt.Println("-------------------------------------")
	fmt.Println("Usage")
	fmt.Println("  go run migrate.go -exec <command>\n")
	fmt.Println("Available Exec Commands: ")
	for command, detail := range available_exec_commands {
		fmt.Println("  " + command + " : " + detail)
	}
	fmt.Println("-------------------------------------")
}
