
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `weight_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `unit` varchar(4) NOT NULL,
  `value` double(15,8) NOT NULL DEFAULT '0.00000000',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO `weight_class` (`id`, `name`, `unit`, `value`) VALUES
(1, 'Kilogram', 'kg', '1.00000000'),
(2, 'Gram', 'g', '1000.00000000'),
(5, 'Pound ', 'lb', '2.20460000'),
(6, 'Ounce', 'oz', '35.27400000');


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `weight_class`;
