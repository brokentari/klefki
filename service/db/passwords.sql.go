// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: passwords.sql

package db

import (
	"context"

	uuid "github.com/gofrs/uuid/v5"
)

const createPassword = `-- name: CreatePassword :one
INSERT INTO passwords(
  id, user_id, username, password, name, website
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, user_id, username, password, name, website
`

type CreatePasswordParams struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Website  string    `json:"website"`
}

func (q *Queries) CreatePassword(ctx context.Context, arg CreatePasswordParams) (Password, error) {
	row := q.db.QueryRow(ctx, createPassword,
		arg.ID,
		arg.UserID,
		arg.Username,
		arg.Password,
		arg.Name,
		arg.Website,
	)
	var i Password
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.Website,
	)
	return i, err
}

const deletePassword = `-- name: DeletePassword :exec
DELETE FROM passwords
WHERE id = $1
`

func (q *Queries) DeletePassword(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deletePassword, id)
	return err
}

const getAllPasswords = `-- name: GetAllPasswords :many
SELECT id, user_id, username, password, name, website FROM passwords
ORDER BY name
`

func (q *Queries) GetAllPasswords(ctx context.Context) ([]Password, error) {
	rows, err := q.db.Query(ctx, getAllPasswords)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Password
	for rows.Next() {
		var i Password
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Username,
			&i.Password,
			&i.Name,
			&i.Website,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserPasswords = `-- name: GetUserPasswords :many
SELECT id, user_id, username, password, name, website FROM passwords
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUserPasswords(ctx context.Context, userID uuid.UUID) ([]Password, error) {
	rows, err := q.db.Query(ctx, getUserPasswords, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Password
	for rows.Next() {
		var i Password
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Username,
			&i.Password,
			&i.Name,
			&i.Website,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE passwords
  SET name = $2,
  username = $3, 
  password = $4, 
  website = $5
WHERE id = $1
`

type UpdatePasswordParams struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Website  string    `json:"website"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.Exec(ctx, updatePassword,
		arg.ID,
		arg.Name,
		arg.Username,
		arg.Password,
		arg.Website,
	)
	return err
}