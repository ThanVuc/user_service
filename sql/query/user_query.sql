-- name: CreateUserAndAuthor :exec

WITH new_user AS (
  INSERT INTO users (name, age, gender, address)
  VALUES ($1, $2, $3, $4)
  RETURNING id
)
INSERT INTO authors (name, bio, user_id)
SELECT $5, $6, id FROM new_user;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;
