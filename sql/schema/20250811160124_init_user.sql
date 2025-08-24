-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id uuid NOT NULL,
    fullname character varying(255),
    email character varying(255) NOT NULL,
    avatar_url text,
    bio text,
    slug character varying(255),
    date_of_birth date,
    gender boolean DEFAULT false,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    sentence text,
    author varchar(256) DEFAULT 'Khuyết Danh',
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_slug_key UNIQUE (slug)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
