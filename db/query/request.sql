-- name: CreateRequest :one
INSERT INTO requests (email, status) VALUES ($1, $2) RETURNING *;

-- name: GetRequest :one
SELECT * FROM requests
WHERE id = $1 LIMIT 1;

-- name: GetByStatus :many
SELECT * FROM requests
WHERE status = $1 
FOR NO KEY UPDATE;

-- name: GetByEmail :many
SELECT * FROM requests
WHERE email = $1;

-- name: DeleteRequest :exec
DELETE FROM requests WHERE id = $1;

-- name: UpdateStatus :one
UPDATE requests
SET status = $2
WHERE id = $1
RETURNING *;

-- name: UpdateSongID :one
UPDATE requests
SET songid = $2
WHERE id = $1
RETURNING *;