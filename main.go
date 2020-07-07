package main

import (
	"database/sql"
	"dbconnection" // connects the go code with the database
	"dbdesign"     // database schema design (no. of fieilds)
	"dboperations" // dboperations contains the basic CRUD opetions of the database

	_ "fmt" // standard format library
	"log"
	"net"
	"net/http"
	"net/rpc"

	_ "github.com/lib/pq" // features specific to postgresql
)

// global pointer for database access in API struct methods
var db *sql.DB

// API struct has methods associated with them
type API struct{}

// Add inserts data into the table
func (a *API) Add(row dbdesign.Row, reply *dbdesign.Row) error {
	dboperations.Insert(db, row)
	*reply = row
	return nil
}

func main() {
	db = dbconnection.ConnectDB()

	// defer postpones the execution of db.Close() function until main() is completed
	defer db.Close()

	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error in API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 3000)

	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

	// dboperations.Insert(db)

	// dboperations.Update(db)

	// dboperations.Delete(db)
}
