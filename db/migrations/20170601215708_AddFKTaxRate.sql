
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE tax_rate ADD CONSTRAINT fk_tax_rate_geo_zone FOREIGN KEY (geo_zone_id) REFERENCES geo_zone(id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE tax_rate DROP FOREIGN KEY fk_tax_rate_geo_zone;