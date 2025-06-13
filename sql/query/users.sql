-- name: GetUsers :many
SELECT 
    user_id,
    fullname,
    username
FROM 
    user_profile;