-- name: GetDiscount :one
SELECT *
FROM discounts
WHERE discount_id = $1
LIMIT 1;

-- name: ListDiscounts :many
SELECT *
FROM discounts
ORDER BY discount_id
LIMIT $1
    OFFSET $2;

-- name: CreateDiscount :one
INSERT INTO discounts (product_id, discount_percentage, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateDiscount :one
UPDATE discounts
SET product_id = $1, discount_percentage = $2, updated_at = $3
WHERE discount_id = $4
RETURNING *;

-- name: DeleteDiscount :exec
DELETE FROM discounts WHERE discount_id = $1;
