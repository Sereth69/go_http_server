-- name: CreateUser :one
INSERT INTO users (id, email, hashed_password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, created_at, updated_at;
-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, email, hashed_password FROM users
WHERE email = $1;