
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE region ADD CONSTRAINT fk_region_country FOREIGN KEY (country_id) REFERENCES country(id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE region DROP FOREIGN KEY fk_region_country;
