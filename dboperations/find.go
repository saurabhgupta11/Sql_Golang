package dboperations

import (
	"database/sql"
	"fmt"
)

// Find let's you query for data in the database
func Find(db *sql.DB) {
	sqlStatement := `SELECT id, email FROM users WHERE id=$1;`
	var email string
	var id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement, 1)
	switch err := row.Scan(&id, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, email)
	default:
		panic(err)
	}
}
