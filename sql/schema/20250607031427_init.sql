-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_profile (
    user_id UUID PRIMARY KEY NOT NULL,
    fullname VARCHAR(255) NOT NULL DEFAULT '',
    username VARCHAR(255) UNIQUE NOT NULL DEFAULT '',
    avatar_url TEXT DEFAULT '',
    bio TEXT DEFAULT '',
    slug VARCHAR(255) UNIQUE DEFAULT '',
    date_of_birth DATE NOT NULL DEFAULT CURRENT_DATE,
    gender BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_profile;
-- +goose StatementEnd