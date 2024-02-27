-- name: GetOrder :one
SELECT *
FROM orders
WHERE order_id = $1
LIMIT 1;

-- name: ListOrders :many
SELECT *
FROM orders
ORDER BY order_id
LIMIT $1
    OFFSET $2;

-- name: CreateOrder :one
INSERT INTO orders (order_date, buyer_id, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateOrder :one
UPDATE orders
SET order_date = $1, buyer_id = $2, updated_at = $3
WHERE order_id = $4
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE order_id = $1;
