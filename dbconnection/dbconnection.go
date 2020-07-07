package dbconnection

import (
	"database/sql" // this package let's us use general sql features in golang
	"dbdetails"    // the details of database name and password and port
	"fmt"          // the standard format library

	_ "github.com/lib/pq" // features specific to postgresql
)

// ConnectDB return a pointer of type db.sql
func ConnectDB() *sql.DB {
	// sql.Open() is used to validate our database details
	// database details are returned as a string
	db, err := sql.Open("postgres", dbdetails.SQLInfo())
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// db.Ping() actually conncects to the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
