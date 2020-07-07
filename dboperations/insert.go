package dboperations

import (
	"database/sql"
	"dbdesign"
	"fmt"
)

// Insert function inserts a row in table
func Insert(db *sql.DB, row dbdesign.Row) {
	// Executing Insert query on the database
	sqlStatement := `
		INSERT INTO
		users (age, first_name, last_name, email)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	id := 0
	err := db.QueryRow(sqlStatement, row.Age, row.FirstName, row.LastName, row.Email).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
