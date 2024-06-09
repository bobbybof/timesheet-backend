-- name: GetProjects :many
SELECT * FROM projects;

-- name: CreateProject :one
INSERT INTO projects (name)
VALUES ($1)
RETURNING *;
