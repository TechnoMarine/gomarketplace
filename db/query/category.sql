-- name: GetCategory :one
SELECT *
FROM categories
WHERE category_id = $1
LIMIT 1;

-- name: ListCategories :many
SELECT *
FROM categories
ORDER BY category_id
LIMIT $1
    OFFSET $2;

-- name: CreateCategory :one
INSERT INTO categories (category_name, parent_id, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateCategory :one
UPDATE categories
SET category_name = $1, parent_id = $2, updated_at = $3
WHERE category_id = $4
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE category_id = $1;