
-- +migrate Up

CREATE TABLE `users` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  `deleted_at` DATETIME NULL DEFAULT NULL,
  `name` VARCHAR(191) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `uid` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

-- +migrate Down
