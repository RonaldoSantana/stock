
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `address` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `supplier_id` int(11) NOT NULL,
  `line_1` varchar(255) NOT NULL,
  `line_2` varchar(255) NOT NULL,
  `line_3` varchar(255) NOT NULL,
  `city` varchar(150) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `postcode` varchar(10) NOT NULL,
  `country_id` int(11) NULL,
  `region_id` int(11) NULL,
  PRIMARY KEY (`id`),
  KEY `supplier_id` (`supplier_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `address`;
