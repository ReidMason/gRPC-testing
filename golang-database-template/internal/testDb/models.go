// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package testDb

import (
	"database/sql"
)

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

type Book struct {
	ID       int64
	Authorid int64
	Title    string
}
