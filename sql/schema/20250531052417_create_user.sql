-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id      BIGSERIAL PRIMARY KEY, -- must match authors.id
  name    TEXT NOT NULL,
  age     INTEGER,
  gender  TEXT,
  address TEXT,
  CONSTRAINT fk_author
    FOREIGN KEY (id)
    REFERENCES authors(id)
    ON DELETE CASCADE
);

ALTER TABLE authors
ADD COLUMN user_id BIGINT UNIQUE;

ALTER TABLE authors
ADD CONSTRAINT fk_user
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
ALTER TABLE authors
DROP CONSTRAINT IF EXISTS fk_user;
ALTER TABLE authors
DROP COLUMN IF EXISTS user_id;
-- +goose StatementEnd
