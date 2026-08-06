-- name: CreateSession :one
INSERT INTO sessions (
    user_id,
    token_hash,
    created_at,
    expires_at,
    last_used_at,
    user_agent,
    ip_address
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: GetSessionByTokenHash :one
SELECT *
FROM sessions
WHERE token_hash = $1
AND expires_at > NOW();

-- name: UpdateSessionLastUsed :exec
UPDATE sessions
SET last_used_at = NOW()
WHERE id = $1;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1;

-- name: DeleteSessionsByUserID :exec
DELETE FROM sessions
WHERE user_id = $1;

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions
WHERE expires_at <= NOW();

-- name: DeleteSessionByTokenHash :exec
DELETE FROM sessions
WHERE token_hash = $1;