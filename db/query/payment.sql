-- name: GetPayment :one
SELECT *
FROM payments
WHERE payment_id = $1
LIMIT 1;

-- name: ListPayments :many
SELECT *
FROM payments
ORDER BY payment_id
LIMIT $1
    OFFSET $2;

-- name: CreatePayment :one
INSERT INTO payments (order_id, amount, status, payment_date, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdatePayment :one
UPDATE payments
SET order_id = $1, amount = $2, status = $3, payment_date = $4, updated_at = $5
WHERE payment_id = $6
RETURNING *;

-- name: DeletePayment :exec
DELETE FROM payments WHERE payment_id = $1;
