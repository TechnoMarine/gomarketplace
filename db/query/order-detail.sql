-- name: GetOrderDetail :one
SELECT *
FROM order_details
WHERE detail_id = $1
LIMIT 1;

-- name: ListOrderDetails :many
SELECT *
FROM order_details
ORDER BY detail_id
LIMIT $1
    OFFSET $2;

-- name: CreateOrderDetail :one
INSERT INTO order_details (order_id, product_id, quantity, total_amount, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateOrderDetail :one
UPDATE order_details
SET order_id = $1, product_id = $2, quantity = $3, total_amount = $4, updated_at = $5
WHERE detail_id = $6
RETURNING *;

-- name: DeleteOrderDetail :exec
DELETE FROM order_details WHERE detail_id = $1;
