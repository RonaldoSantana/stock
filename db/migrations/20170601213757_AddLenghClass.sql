
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `length_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `unit` varchar(4) NOT NULL,
  `value` double(15,8) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO `length_class` (`id`, `name`, `unit`, `value`) VALUES
(1, 'Centimeter', 'cm', '1.00000000'),
(2, 'Millimeter', 'mm', '10.00000000'),
(3, 'Inch', 'in', '0.39370000');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `length_class`;
