-- name: AddTransactions :many
INSERT INTO
  transactions (id, amount, date, user_id)
SELECT
  unnest(sqlc.arg (ids)::BIGINT[]) AS id,
  unnest(sqlc.arg (amounts)::DECIMAL(10, 2) []) AS amount,
  unnest(sqlc.arg (dates)::DATE[]) AS date,
  unnest(sqlc.arg (user_ids)::BIGINT[]) AS user_id
RETURNING
  *;
