
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO `tax_rate` (`id`, `geo_zone_id`, `name`, `rate`, `type`) VALUES
(1, 1, 'GST (15%)', '15.0000', 'percent');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM tax_rate WHERE id = 1;
