-- +goose Up
-- +goose StatementBegin
-- Add a unique constraint check to ensure only one row exists in scraper_status
-- This table is meant to store a single row with the last processed date
ALTER TABLE scraper_status ADD CONSTRAINT single_row_check CHECK (id = 1);

-- Insert the initial row if it doesn't exist
INSERT INTO scraper_status (id, last_processed_date)
VALUES (1, '2024-01-01')
ON CONFLICT (id) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE scraper_status DROP CONSTRAINT IF EXISTS single_row_check;
-- +goose StatementEnd
