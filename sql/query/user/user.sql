-- name: GetUserProfile :one
select * from users where id = $1;