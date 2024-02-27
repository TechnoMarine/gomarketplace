-- name: GetShippingAddress :one
SELECT *
FROM shipping_addresses
WHERE address_id = $1
LIMIT 1;

-- name: ListShippingAddresses :many
SELECT *
FROM shipping_addresses
ORDER BY address_id
LIMIT $1
    OFFSET $2;

-- name: CreateShippingAddress :one
INSERT INTO shipping_addresses (user_id, address, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateShippingAddress :one
UPDATE shipping_addresses
SET user_id = $1, address = $2, updated_at = $3
WHERE address_id = $4
RETURNING *;

-- name: DeleteShippingAddress :exec
DELETE FROM shipping_addresses WHERE address_id = $1;
