// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package testDatabase

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (
  id, name, bio
) VALUES (
  $1, $2, $3 
)
RETURNING id, bio, name
`

type CreateAuthorParams struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.ID, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(&i.ID, &i.Bio, &i.Name)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, bio, name FROM authors
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Bio, &i.Name)
	return i, err
}

const getAuthorById = `-- name: GetAuthorById :one
SELECT name FROM authors
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthorById(ctx context.Context, id int64) (string, error) {
	row := q.db.QueryRowContext(ctx, getAuthorById, id)
	var name string
	err := row.Scan(&name)
	return name, err
}

const getAuthorsWithBooks = `-- name: GetAuthorsWithBooks :many
SELECT authors.id, bio, name, books.id, authorid, title FROM authors
INNER JOIN books on books.authorid = authors.id
WHERE authors.id = ANY($1::int[])
`

type GetAuthorsWithBooksRow struct {
	ID       int64
	Bio      sql.NullString
	Name     string
	ID_2     int64
	Authorid int64
	Title    string
}

// WHERE authors.id in (sqlc.slice('ids'));
func (q *Queries) GetAuthorsWithBooks(ctx context.Context, dollar_1 []int32) ([]GetAuthorsWithBooksRow, error) {
	rows, err := q.db.QueryContext(ctx, getAuthorsWithBooks, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAuthorsWithBooksRow
	for rows.Next() {
		var i GetAuthorsWithBooksRow
		if err := rows.Scan(
			&i.ID,
			&i.Bio,
			&i.Name,
			&i.ID_2,
			&i.Authorid,
			&i.Title,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, bio, name FROM authors
ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Bio, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
