-- name: GetProduct :one
SELECT *
FROM products
WHERE product_id = $1
LIMIT 1;

-- name: ListProducts :many
SELECT *
FROM products
ORDER BY product_id
LIMIT $1
    OFFSET $2;

-- name: CreateProduct :one
INSERT INTO products (name, description, price, stock_quantity, seller_id, category_id, is_active, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
SET name = $1, description = $2, price = $3, stock_quantity = $4, seller_id = $5, category_id = $6, is_active = $7, updated_at = $8
WHERE product_id = $9
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products WHERE product_id = $1;
