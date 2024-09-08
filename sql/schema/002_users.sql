-- +goose Up
ALTER TABLE users ADD COLUMN api_key VARCHAR(128) UNIQUE NOT NULL DEFAULT (
  encode(sha512(random()::text::bytea), 'hex') 
);
-- +goose Down
ALTER TABLE users DROP COLUMN api_key;
