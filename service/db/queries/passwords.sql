-- name: GetUserPasswords :many
SELECT * FROM passwords
WHERE user_id = $1;

-- name: GetAllPasswords :many
SELECT * FROM passwords
ORDER BY name;

-- name: CreatePassword :one
INSERT INTO passwords(
  id, user_id, username, password, name, website
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: DeletePassword :exec
DELETE FROM passwords
WHERE id = $1;

-- name: UpdatePassword :exec
UPDATE passwords
  SET name = $2,
  username = $3, 
  password = $4, 
  website = $5
WHERE id = $1;