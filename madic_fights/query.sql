-- name: CreateRoom :one

INSERT INTO rooms (key)
VALUES ($1)
RETURNING *;