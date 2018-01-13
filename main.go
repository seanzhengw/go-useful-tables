package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/seanzhengw/go-useful-tables/user"

	// DB driver
	_ "github.com/go-sql-driver/mysql"
)

var name = flag.String("name", "", "database user name")
var password = flag.String("passwd", "", "database user password")
var host = flag.String("host", "localhost", "database server hostname")
var port = flag.Int("port", 3306, "database server port")
var dbname = flag.String("db", "", "enter db name")

type tablesPair struct {
	name       string
	createFunc func(sqldb *sql.DB) (sql.Result, error)
}

var tablesPairs map[string](tablesPair)

func initTablesPairs() {
	tablesPairs = make(map[string](tablesPair))
	tablesPairs["user"] = tablesPair{name: "user", createFunc: user.CreateTable}
	tablesPairs["user_detail"] = tablesPair{name: "user_detail", createFunc: user.CreateDetail}
	tablesPairs["user_emails"] = tablesPair{name: "user_emails", createFunc: user.CreateEmails}
	tablesPairs["user_emails_no_id"] = tablesPair{name: "user_emails", createFunc: user.CreateEmailsWithoutID}
}

func printTableNames() {
	for k := range tablesPairs {
		fmt.Printf("%s ", k)
	}
}

func main() {
	initTablesPairs()
	flag.Usage = func() {
		fmt.Printf("Usage: go-useful-tables -name <username> -passwd <password> -db <dbname> [options] <table name> [[table name 2]...] \n\n")
		flag.PrintDefaults()
		fmt.Printf("\n")
		fmt.Printf("Available tables:\n")
		fmt.Printf("\t")
		printTableNames()
		fmt.Printf("\n\n")
	}
	flag.Parse()

	if len(*name) == 0 {
		fmt.Println("Input database user name with flag: -name <username>")
		flag.Usage()
		os.Exit(1)
	}
	if len(*password) == 0 {
		fmt.Println("Input database user password with flag: -passwd <password>")
		flag.Usage()
		os.Exit(1)
	}
	if len(*dbname) == 0 {
		fmt.Println("Input database name with flag: -db <dbname>")
		flag.Usage()
		os.Exit(1)
	}

	tables := flag.Args()
	if len(tables) == 0 {
		fmt.Println("Input at least one table name for creation.")
		flag.Usage()
		os.Exit(1)
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", *name, *password, *host, *port, *dbname)

	sqldb, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer sqldb.Close()

	for _, table := range tables {
		pair, ok := tablesPairs[table]
		if ok {
			_, err = pair.createFunc(sqldb)
			if err == nil {
				fmt.Printf("Table %s created. (%s)\n", pair.name, table)
			}
		} else {
			err = fmt.Errorf("Unknown table %s", table)
		}
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}
}
