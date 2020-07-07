package dbdetails

import (
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgresDatabase"
	dbname   = "postgres"
)

// SQLInfo is a function that returns a string to connect to the database
func SQLInfo() string {
	// psqlInfo is a string containing details to access database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlInfo
}
