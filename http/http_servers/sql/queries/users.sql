-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, password_hash)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: LoginUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users SET id = $1, email = $2, password_hash = $3, created_at = $4, updated_at = $5 WHERE id = $1
RETURNING *;

-- name: UpgradeUserToChirpyRed :one
UPDATE users SET is_chirpy_red = TRUE, updated_at = NOW() WHERE id = $1
RETURNING *;
