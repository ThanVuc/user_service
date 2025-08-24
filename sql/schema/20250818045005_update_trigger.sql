-- +goose Up
-- +goose StatementBegin

-- Create the timestamp trigger function (if not exists)
CREATE OR REPLACE FUNCTION set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    -- Only update updated_at on UPDATE
    IF TG_ARGV[0] = 'updated_at' AND TG_OP = 'UPDATE' THEN
        NEW.updated_at = CURRENT_TIMESTAMP;
        RETURN NEW;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger for users table
CREATE TRIGGER trg_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_timestamp('updated_at');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop trigger for users table
DROP TRIGGER IF EXISTS trg_users_updated_at ON users;

-- Optionally drop function if no other table uses it
DROP FUNCTION IF EXISTS set_timestamp();

-- +goose StatementEnd
