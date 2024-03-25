-- +goose Up
-- +goose StatementBegin
CREATE TABLE Adverts(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(256) NOT NULL,
    text TEXT NOT NULL,
    image_url VARCHAR(256) NOT NULL,
    price BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_adverts_user_id
        FOREIGN KEY (user_id)
        REFERENCES Users (id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Adverts;
-- +goose StatementEnd
