-- name: InsertUser :one
INSERT INTO users (id, email, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING id, email, created_at, updated_at;
