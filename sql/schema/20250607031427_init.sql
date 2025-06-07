-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_profile (
    user_id UUID PRIMARY KEY NOT NULL,
    fullname VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    avatar_url TEXT,
    bio TEXT,
    slug VARCHAR(255) UNIQUE,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_profile;
-- +goose StatementEnd