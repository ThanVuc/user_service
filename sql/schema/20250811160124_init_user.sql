-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id uuid NOT NULL,
    fullname character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    avatar_url text,
    bio text,
    slug character varying(255),
    date_of_birth date NOT NULL,
    gender character varying(50) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    sentence character varying NOT NULL,
    author character varying DEFAULT 'Khuyết Danh',
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_slug_key UNIQUE (slug)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
