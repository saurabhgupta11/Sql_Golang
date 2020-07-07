package dboperations

import (
	"database/sql"
	"dbdesign"
)

// Update function updates data in the table
func Update(db *sql.DB, row dbdesign.Row) dbdesign.Row {
	// Executing Update Query on the database
	sqlStatement := `
		UPDATE users
		SET age = $2, first_name = $3, last_name = $4, email = $5
		WHERE id = $1
		RETURNING id;
	`
	var id int
	err := db.QueryRow(sqlStatement, row.ID, row.Age, row.FirstName, row.LastName, row.Email).Scan(&id)
	if err != nil {
		panic(err)
	}

	reply := dbdesign.Row{
		ID:        id,
		Age:       row.Age,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Email:     row.Email,
	}

	return reply
}
