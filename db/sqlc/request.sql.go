// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: request.sql

package db

import (
	"context"
)

const createRequest = `-- name: CreateRequest :one
INSERT INTO requests (email, status) VALUES ($1, $2) RETURNING id, email, status, songid
`

type CreateRequestParams struct {
	Email  string `json:"email"`
	Status string `json:"status"`
}

func (q *Queries) CreateRequest(ctx context.Context, arg CreateRequestParams) (Request, error) {
	row := q.db.QueryRowContext(ctx, createRequest, arg.Email, arg.Status)
	var i Request
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Status,
		&i.Songid,
	)
	return i, err
}

const deleteRequest = `-- name: DeleteRequest :exec
DELETE FROM requests WHERE id = $1
`

func (q *Queries) DeleteRequest(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteRequest, id)
	return err
}

const getByEmail = `-- name: GetByEmail :many
SELECT id, email, status, songid FROM requests
WHERE email = $1
`

func (q *Queries) GetByEmail(ctx context.Context, email string) ([]Request, error) {
	rows, err := q.db.QueryContext(ctx, getByEmail, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Request
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Status,
			&i.Songid,
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

const getByStatus = `-- name: GetByStatus :many
SELECT id, email, status, songid FROM requests
WHERE status = $1 
FOR NO KEY UPDATE
`

func (q *Queries) GetByStatus(ctx context.Context, status string) ([]Request, error) {
	rows, err := q.db.QueryContext(ctx, getByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Request
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Status,
			&i.Songid,
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

const getRequest = `-- name: GetRequest :one
SELECT id, email, status, songid FROM requests
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRequest(ctx context.Context, id int64) (Request, error) {
	row := q.db.QueryRowContext(ctx, getRequest, id)
	var i Request
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Status,
		&i.Songid,
	)
	return i, err
}

const updateSongID = `-- name: UpdateSongID :one
UPDATE requests
SET songid = $2
WHERE id = $1
RETURNING id, email, status, songid
`

type UpdateSongIDParams struct {
	ID     int64  `json:"id"`
	Songid string `json:"songid"`
}

func (q *Queries) UpdateSongID(ctx context.Context, arg UpdateSongIDParams) (Request, error) {
	row := q.db.QueryRowContext(ctx, updateSongID, arg.ID, arg.Songid)
	var i Request
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Status,
		&i.Songid,
	)
	return i, err
}

const updateStatus = `-- name: UpdateStatus :one
UPDATE requests
SET status = $2
WHERE id = $1
RETURNING id, email, status, songid
`

type UpdateStatusParams struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

func (q *Queries) UpdateStatus(ctx context.Context, arg UpdateStatusParams) (Request, error) {
	row := q.db.QueryRowContext(ctx, updateStatus, arg.ID, arg.Status)
	var i Request
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Status,
		&i.Songid,
	)
	return i, err
}
