-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_transactions (
   id int AUTO_INCREMENT PRIMARY KEY,
   date DATE,
   amount NUMERIC(15,2) NOT NULL,
   observation VARCHAR(120),
   protocol VARCHAR(225),
   status VARCHAR(12),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_transactions;
-- +goose StatementEnd
