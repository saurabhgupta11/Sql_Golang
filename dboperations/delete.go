package dboperations

import (
	"database/sql"
	"dbdesign"
)

// Delete function deletes a row in the table
func Delete(db *sql.DB, row dbdesign.Row) int64 {
	// Executing Delete Query on the database
	sqlStatement := `
		DELETE FROM users
		WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, row.ID)
	if err != nil {
		panic(err)
	}

	// used to verify how many rows were affected
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return count
}
