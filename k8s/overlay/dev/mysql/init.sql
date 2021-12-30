CREATE DATABASE app_develop;
USE app_develop;

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

CREATE TABLE `questions` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME NULL DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT NULL,
  `deleted_at` DATETIME NULL DEFAULT NULL,
  `title` VARCHAR(255) NOT NULL,
  `first_answer` VARCHAR(255) NOT NULL,
  `second_answer` VARCHAR(255) NOT NULL,
  `first_count` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0,
  `second_count` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0,
  `first_img_url` VARCHAR(255) NOT NULL,
  `second_img_url` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `app_develop`.`questions` (`id`, `created_at`, `updated_at`, `title`, `first_answer`, `second_answer`, `first_count`, `second_count`, `first_img_url`, `second_img_url`) VALUES ('1', '2021-11-21 03:32:58', '2021-12-11 16:36:57', '犬派?それとも猫派?', 'いぬ', 'ねこ', '46', '76', '/images/answer_img/dog.png', '/images/answer_img/cat.png');
INSERT INTO `app_develop`.`questions` (`id`, `created_at`, `updated_at`, `title`, `first_answer`, `second_answer`, `first_count`, `second_count`, `first_img_url`, `second_img_url`) VALUES ('2', '2021-11-21 03:32:58', '2021-12-11 16:36:39', 'きのこ派?それともたけのこ派?', 'きのこ', 'たけのこ', '118', '79', '/images/answer_img/mushroom.png', '/images/answer_img/bambooshoot.png');
