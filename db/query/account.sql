-- name: CreateAccount :one
INSERT INTO accounts (
  first_name,
  last_name,
  email,
  phone_number,
  balance,
  currency
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts 
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;