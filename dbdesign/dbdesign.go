package dbdesign

// Row defines database schema
type Row struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

// id is defined by the database while creating table and is automatically given
// by the database using SERIAL in sql
// This id is used to execute update and delete queries
/*
	delete some user where id is Row.ID
*/
