-- name: InsertUser :one
INSERT INTO users (id, email, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING id, email, created_at, updated_at;

-- name: GetUserProfile :one
SELECT
    u.id          AS user_id,
    u.fullname    AS fullname,
    u.email       AS email,
    u.avatar_url  AS avatar_url,
    u.bio         AS bio,
    u.slug        AS slug,
    u.date_of_birth AS date_of_birth,
    u.gender      AS gender,
    u.created_at  AS created_at,
    u.updated_at  AS updated_at,
    u.sentence    AS sentence,
    u.author      AS author
FROM users u
WHERE u.id = $1;

-- name: UpdateUserProfile :one
UPDATE users
SET fullname      = $2,
    bio           = $3,
    date_of_birth = $4,
    gender        = $5,
    sentence      = $6,
    author        = $7,
    updated_at    = NOW()
WHERE id = $1
RETURNING id;


-- name: UpdateSlugById :one
UPDATE users
SET slug = $2
WHERE id = $1
RETURNING id;

