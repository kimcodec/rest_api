-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS USERS(
    login VARCHAR(256) UNIQUE NOT NULL,
    password varchar(256) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS USERS;
-- +goose StatementEnd
