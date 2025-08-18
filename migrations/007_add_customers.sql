-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customers (
    id BIGSERIAL PRIMARY KEY,
    external_id BIGINT UNIQUE,
    name VARCHAR(255) NOT NULL,
    login VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS task_customers (
    task_id BIGINT NOT NULL REFERENCES tasks(task_id) ON DELETE CASCADE,
    customer_id BIGINT NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    PRIMARY KEY (task_id, customer_id)
);

ALTER TABLE IF EXISTS tasks DROP COLUMN customer_name, DROP COLUMN customer_login;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS task_customers;
-- +goose StatementEnd
