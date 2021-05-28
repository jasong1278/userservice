// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package repository

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users ( 
  first_name,
  last_name, 
  email
) VALUES (
  $1, 
  $2, 
  $3
)
RETURNING first_name, last_name, email
`

type CreateUserParams struct {
	FirstName string
	LastName  string
	Email     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.FirstName, arg.LastName, arg.Email)
	var i User
	err := row.Scan(&i.FirstName, &i.LastName, &i.Email)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users
WHERE email = $1
RETURNING first_name, last_name, email
`

func (q *Queries) DeleteUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, email)
	var i User
	err := row.Scan(&i.FirstName, &i.LastName, &i.Email)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT first_name, last_name, email FROM users 
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(&i.FirstName, &i.LastName, &i.Email)
	return i, err
}