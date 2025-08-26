-- +goose Up
-- +goose StatementBegin
ALTER TABLE employees ADD COLUMN is_admin BOOLEAN DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
