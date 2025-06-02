-- name: CreateUser :one
INSERT INTO users (username, password, firstname, lastname)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;
