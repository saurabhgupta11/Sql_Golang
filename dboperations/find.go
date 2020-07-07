package dboperations

import (
	"database/sql"
	"dbdesign"
	"fmt"
)

// Find let's you query for data in the database
func Find(db *sql.DB, row dbdesign.Row) dbdesign.Row {
	sqlStatement := `SELECT id, age, first_name, last_name, email FROM users WHERE id=$1;`
	reply := dbdesign.Row{
		ID:        row.ID,
		Age:       row.Age,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Email:     row.Email,
	}
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	r := db.QueryRow(sqlStatement, row.ID)
	switch err := r.Scan(&reply.ID, &reply.Age, &reply.FirstName, &reply.LastName, &reply.Email); err {
	case sql.ErrNoRows:
		fmt.Println("No User Found")
		return reply
	case nil:
		return reply
	default:
		panic(err)
	}
}
