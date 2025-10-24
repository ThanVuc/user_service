-- name: InsertUser :one
INSERT INTO users (id, email,fullname, created_at, updated_at, avatar_url)
VALUES ($1, NULLIF($2, ''), NULLIF($3, ''), $4, $5, NULLIF($6, ''))
ON CONFLICT (id) DO UPDATE SET
    email = COALESCE(EXCLUDED.email, users.email),
    fullname = COALESCE(EXCLUDED.fullname, users.fullname),
    avatar_url = COALESCE(EXCLUDED.avatar_url, users.avatar_url),
    updated_at = EXCLUDED.updated_at
RETURNING id, email,fullname, created_at, updated_at, avatar_url;

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
    avatar_url    = $3,
    bio           = $4,
    date_of_birth = $5,
    gender        = $6,
    sentence      = $7,
    author        = $8,
    updated_at    = NOW()
WHERE id = $1
RETURNING id;

-- name: UpdateAvatarById :one
UPDATE users
SET avatar_url = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id;

-- name: UpdateSlugById :one
UPDATE users
SET slug = $2
WHERE id = $1
RETURNING id;


