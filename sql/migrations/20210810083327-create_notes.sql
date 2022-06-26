-- +migrate Up
CREATE TABLE `notes`(
    `id` INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(50) NOT NULL,
    `gender` VARCHAR(1) NOT NULL,
    `age` VARCHAR(2) NOT NULL,
    `vaccine_type` VARCHAR(1) NOT NULL,
    `number_of_vaccination` TINYINT NOT NULL,
    `max_temperature` VARCHAR(2) NOT NULL,
    `log` TEXT NULL,
    `remarks` TEXT NULL,
    `good_count` INT NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL
) DEFAULT CHARSET=utf8;
    
-- +migrate Down
DROP TABLE IF EXISTS notes;