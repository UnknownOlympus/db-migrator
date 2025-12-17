-- +goose Up
-- +goose StatementBegin
ALTER TABLE bot_users ADD COLUMN IF NOT EXISTS locale TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
