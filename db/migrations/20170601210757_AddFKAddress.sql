
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE address ADD CONSTRAINT fk_address_supplier FOREIGN KEY (supplier_id) REFERENCES supplier(id);
ALTER TABLE address ADD CONSTRAINT fk_address_country FOREIGN KEY (country_id) REFERENCES country(id);
ALTER TABLE address ADD CONSTRAINT fk_address_region FOREIGN KEY (region_id) REFERENCES region(id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE address DROP FOREIGN KEY fk_address_supplier;
ALTER TABLE address DROP FOREIGN KEY fk_address_country;
ALTER TABLE address DROP FOREIGN KEY fk_address_region;
