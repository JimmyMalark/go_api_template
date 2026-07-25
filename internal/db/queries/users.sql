-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CountUsers :one
SELECT COUNT(*)
FROM users;

-- name: CreateUser :one
INSERT INTO users (xid, username, email, first_name, last_name, birth_date, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;