-- Queries for the refresh_tokens table

-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetRefreshTokenByToken :one
SELECT * FROM refresh_tokens WHERE token = $1;

-- name: GetUserIDByRefreshToken :one
SELECT user_id FROM refresh_tokens WHERE token = $1;

-- name: GetAllRefreshTokensForUser :many
SELECT * FROM refresh_tokens WHERE user_id = $1;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens SET revoked_at = $2, updated_at = $3 WHERE token = $1;
