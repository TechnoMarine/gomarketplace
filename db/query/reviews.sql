-- name: GetReview :one
SELECT *
FROM reviews
WHERE review_id = $1
LIMIT 1;

-- name: ListReviews :many
SELECT *
FROM reviews
ORDER BY review_id
LIMIT $1
    OFFSET $2;

-- name: CreateReview :one
INSERT INTO reviews (product_id, user_id, rating, comment, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateReview :one
UPDATE reviews
SET product_id = $1, user_id = $2, rating = $3, comment = $4, updated_at = $5
WHERE review_id = $6
RETURNING *;

-- name: DeleteReview :exec
DELETE FROM reviews WHERE review_id = $1;
