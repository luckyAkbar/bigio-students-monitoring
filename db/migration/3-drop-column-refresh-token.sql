-- +migrate Up notransaction
ALTER TABLE sessions DROP COLUMN IF EXISTS "refresh_token";
-- +migrate Down
ALTER TABLE sessions ADD COLUMN IF NOT EXISTS "refresh_token" TEXT NOT NULL;
