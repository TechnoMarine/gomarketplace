-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserForUpdate :one
SELECT *
FROM users
WHERE id = $1
for update
LIMIT 1;


-- name: ListUsers :many
SELECT *
FROM users
order by id
limit $1
offset $2;

-- name: CreateUser :one
INSERT INTO users (username, role)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET username = $1, role = $2
where id = $3
returning *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;