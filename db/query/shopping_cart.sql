-- name: GetShoppingCart :one
SELECT *
FROM shopping_cart
WHERE cart_id = $1
LIMIT 1;

-- name: ListShoppingCarts :many
SELECT *
FROM shopping_cart
ORDER BY cart_id
LIMIT $1
    OFFSET $2;

-- name: CreateShoppingCart :one
INSERT INTO shopping_cart (user_id, product_id, quantity, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateShoppingCart :one
UPDATE shopping_cart
SET user_id = $1, product_id = $2, quantity = $3, updated_at = $4
WHERE cart_id = $5
RETURNING *;

-- name: DeleteShoppingCart :exec
DELETE FROM shopping_cart WHERE cart_id = $1;
