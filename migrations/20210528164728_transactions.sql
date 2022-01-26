-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_transactions (
   
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_transactions;
-- +goose StatementEnd
