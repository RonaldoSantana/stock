
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE country DROP COLUMN address_format;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE country ADD COLUMN address_format TEXT NULL AFTER iso_code_3;
