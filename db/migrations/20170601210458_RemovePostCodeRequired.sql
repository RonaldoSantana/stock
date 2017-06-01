
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE country DROP COLUMN postcode_required;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE country ADD COLUMN address_format TINYINT(1) NULL DEFAULT 0 AFTER iso_code_3;
