package dboperations

import (
	"database/sql"
	"fmt"
)

// Delete function deletes a row in the table
func Delete(db *sql.DB) {
	// Executing Delete Query on the database
	sqlStatement := `
		DELETE FROM users
		WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, 7)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
