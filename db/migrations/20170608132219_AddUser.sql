
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id` char(36) NOT NULL,
  `first_name` varchar(255) NULL,
  `last_name` varchar(255) NULL,
  `avatar` varchar(255) NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `mobile` varchar(20) NULL,
  `status` enum('active','disabled') NOT NULL DEFAULT 'disabled',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `user`;