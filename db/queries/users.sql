-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;


-- name: UpdateUser :one
UPDATE users
SET name = $1, rate = $2
WHERE id = $3
RETURNING *;