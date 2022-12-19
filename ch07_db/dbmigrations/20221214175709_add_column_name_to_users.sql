-- +goose Up
ALTER TABLE users ADD COLUMN name VARCHAR(255) NOT NULL;

-- +goose Down
ALTER TABLE users DROP COLUMN name;
