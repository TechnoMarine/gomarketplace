-- name: GetPost :one
SELECT *
FROM posts
WHERE id = $1
LIMIT 1;

-- name: GetPostForUpdate :one
SELECT *
FROM posts
WHERE id = $1
for update
LIMIT 1;


-- name: ListPosts :many
SELECT *
FROM posts
order by id
limit $1
    offset $2;

-- name: CreatePost :one
INSERT INTO posts (title, body, user_id, status)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdatePost :one
UPDATE posts
SET title = $1, body = $2, status = $3
where id = $4
returning *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;