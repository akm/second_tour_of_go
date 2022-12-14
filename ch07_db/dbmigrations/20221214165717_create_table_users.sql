-- +goose Up
CREATE TABLE users (
    id serial PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    created_at datetime  default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    UNIQUE INDEX unq_users_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_bin;
-- +goose Down
DROP TABLE users;
