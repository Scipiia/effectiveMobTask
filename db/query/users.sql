-- name: CreateUser :one
INSERT INTO users (
  name,
  surname,
  patronymic
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
  set name = $2,
  surname = $3
WHERE id = $1
RETURNING *;