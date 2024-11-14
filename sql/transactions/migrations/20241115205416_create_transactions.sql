-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions (
  id BIGSERIAL PRIMARY KEY,
  amount DECIMAL(10, 2) NOT NULL,
  date DATE NOT NULL,
  user_id BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE transactions IS 'Table to store user transactions';

COMMENT ON COLUMN transactions.amount IS 'Amount of the transaction. Credit transactions are
indicated with a plus sign like +60.5. Debit transactions are indicated by a minus sign like -20.46';

COMMENT ON COLUMN transactions.date IS 'The date the transaction was made';

COMMENT ON COLUMN transactions.user_id IS 'The user who made the transaction';

CREATE INDEX transactions_user_id_idx ON transactions (user_id, created_at desc);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;

-- +goose StatementEnd
