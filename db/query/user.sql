-- name: GetUser :one
SELECT *
FROM users
WHERE user_id = $1
LIMIT 1;

-- name: GetUserForUpdate :one
SELECT *
FROM users
WHERE user_id = $1
for update
LIMIT 1;


-- name: ListUsers :many
SELECT *
FROM users
order by user_id
limit $1
offset $2;

-- name: CreateUser :one
INSERT INTO users (first_name, last_name, sur_name, email, password, address, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET first_name = $1, last_name = $2, sur_name = $3, email = $4, password = $5, address = $6
where user_id = $7
returning *;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;