package pqhelper_test

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/northbright/pqhelper"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "postgres"
)

func ExampleTableExists() {
	var err error
	var exists bool
	var db *sql.DB

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	if db, err = sql.Open("postgres", connStr); err != nil {
		log.Printf("sql.Open() error: %v\n", err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Printf("db.Ping() error: %v\n", err)
		return
	}

	// Check non-exist table of public scheme.
	if exists, err = pqhelper.TableExists(db, "public", "non_exist_table"); err != nil {
		log.Printf("pqhelper.TableExists() error: %v\n", err)
	}

	fmt.Printf("%v\n", exists)

	// Check built-in "pg_database" table of pg_catalog scheme.
	if exists, err = pqhelper.TableExists(db, "pg_catalog", "pg_database"); err != nil {
		log.Printf("pqhelper.TableExists() error: %v\n", err)
	}

	fmt.Printf("%v\n", exists)

	// Output:
	//false
	//true
}
