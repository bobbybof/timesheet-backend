-- name: CreateActivity :one
INSERT INTO activities (
    name,
    date_start,
    date_end,
    project_id,
    user_id
) VALUES ($1,$2,$3,$4,$5)
RETURNING *;


-- name: UpdateActivity :one
UPDATE activities
SET 
    name = $1,
    date_start = $2,
    date_end = $3,
    project_id = $4
WHERE 
    id = $5
RETURNING *;


-- name: DeleteActivity :exec
DELETE FROM activities
WHERE id = $1;