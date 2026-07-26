-- name: GetUser :one
SELECT *
FROM users
WHERE xid = $1;

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
INSERT INTO users (xid, username, email, first_name, last_name, password_hash, birth_date, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;