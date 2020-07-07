package dboperations

import (
	"database/sql"
	"fmt"
)

// Update function updates data in the table
func Update(db *sql.DB) {
	// Executing Update Query on the database
	sqlStatement := `
		UPDATE users
		SET first_name = $2, last_name = $3, email = $4
		WHERE id = $1
		RETURNING id, email;
	`
	var id int
	var email string
	err := db.QueryRow(sqlStatement, 2, "Brett", "Yang", "brettyang@gmail.com").Scan(&id, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, email)
}
