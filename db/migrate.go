package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

//sql and database info
const (
	source   = "file://./sql/"
	database = "mysql://teixy:teixy@tcp(0.0.0.0:3306)/teixy_article"
)

//declare command line options
var (
	command = flag.String("exec", "", "set up or down as a argument")
	force   = flag.Bool("f", false, "force exec fixed sql")
)

//available command list
var available_exec_commands = map[string]string{
	"up":      "Execute up sqls",
	"down":    "Execute down sqls",
	"version": "Just check current migrate version",
}

func main() {

	flag.Parse()
	if len(*command) < 1 {
		fmt.Println("\nerror: no argument\n")
		showUsageMessge()
		os.Exit(1)
		return
	}

	m, err := migrate.New(source, database)
	if err != nil {
		fmt.Println("err", err)
	}
	version, dirty, err := m.Version()
	showVersionInfo(version, dirty, err)

	fmt.Println("command: exec", *command)
	switch *command {
	case "up":
		upSql(m, version, dirty)
	case "down":
		downSql(m, version, dirty)
	case "version":
		//do nothing
	default:
		fmt.Println("\nerror: invalid command '" + *command + "'\n")
		showUsageMessge()
		os.Exit(1)
	}
}

//exec up sqls
//with force option if needed
func upSql(m *migrate.Migrate, version uint, dirty bool) {
	if dirty && *force {
		fmt.Println("force=true: force execute current version sql")
		m.Force(int(version))
	}
	err := m.Up()
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	} else {
		fmt.Println("success:", *command+"\n")
		fmt.Println("updated version info")
		version, dirty, err := m.Version()
		showVersionInfo(version, dirty, err)
	}
}

//exec up sqls
//with force option if needed
func downSql(m *migrate.Migrate, version uint, dirty bool) {
	if dirty && *force {
		fmt.Println("force=true: force execute current version sql")
		m.Force(int(version))
	}
	err := m.Down()
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	} else {
		fmt.Println("success:", *command+"\n")
		fmt.Println("updated version info")
		version, dirty, err := m.Version()
		showVersionInfo(version, dirty, err)
	}
}

func showUsageMessge() {
	fmt.Println("-------------------------------------")
	fmt.Println("Usage")
	fmt.Println("  go run migrate.go -exec <command>\n")
	fmt.Println("Available Exec Commands: ")
	for available_command, detail := range available_exec_commands {
		fmt.Println("  " + available_command + " : " + detail)
	}
	fmt.Println("-------------------------------------")
}

func showVersionInfo(version uint, dirty bool, err error) {
	fmt.Println("-------------------")
	fmt.Println("version  : ", version)
	fmt.Println("dirty    : ", dirty)
	fmt.Println("error    : ", err)
	fmt.Println("-------------------")
}
