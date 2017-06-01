
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `geo_zone_matrix` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `country_id` int(11) NOT NULL,
  `region_id` int(11) NULL,
  `geo_zone_id` int(11) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

ALTER TABLE geo_zone_matrix ADD CONSTRAINT fk_geo_zone_matrix_country FOREIGN KEY (country_id) REFERENCES country(id);
ALTER TABLE geo_zone_matrix ADD CONSTRAINT fk_geo_zone_matrix_region FOREIGN KEY (region_id) REFERENCES region(id);
ALTER TABLE geo_zone_matrix ADD CONSTRAINT fk_geo_zone_matrix_geo_zone FOREIGN KEY (geo_zone_id) REFERENCES geo_zone(id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `geo_zone_matrix`;
