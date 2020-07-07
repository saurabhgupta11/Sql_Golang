package main

import (
	"database/sql"
	"dbconnection" // connects the go code with the database
	"dbdesign"     // database schema design (no. of fieilds)
	"dboperations" // dboperations contains the basic CRUD opetions of the database
	"errors"

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
	*reply = dboperations.Insert(db, row)
	return nil
}

// Update Method updates the data in the database
func (a *API) Update(row dbdesign.Row, reply *dbdesign.Row) error {
	*reply = dboperations.Update(db, row)
	return nil
}

// Delete Method deletes the data in the database
func (a *API) Delete(row dbdesign.Row, reply *dbdesign.Row) error {
	rowsAffected := dboperations.Delete(db, row)
	if rowsAffected == 0 {
		return errors.New("No User Found")
	}
	return nil
}

// Find Method Find the data in the database
func (a *API) Find(row dbdesign.Row, reply *dbdesign.Row) error {
	*reply = dboperations.Find(db, row)
	if reply.Email == "" {
		return errors.New("No User Found")
	}
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

	// dboperations.Delete(db)
}
