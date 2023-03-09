-- +goose Up
START TRANSACTION;

    CREATE TABLE `idn_province` (
        `pkid` TINYINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `slug` VARCHAR(75) NULL DEFAULT NULL,
        `code` VARCHAR(5) NULL DEFAULT NULL,
        `iso_code` VARCHAR(20) NULL DEFAULT NULL,
        `name` VARCHAR(75) NOT NULL,
        `geographic_area` VARCHAR(75) NOT NULL,
        `capital` VARCHAR(75) NULL DEFAULT NULL,
        `image` VARCHAR(255) NULL DEFAULT NULL,
        `created_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `updated_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `deleted_at` DATETIME(6) NULL DEFAULT NULL,
        PRIMARY KEY (`pkid`),
        UNIQUE (`slug`, `code`, `iso_code`)
    );

COMMIT;

-- +goose Down
START TRANSACTION;
    DROP TABLE `idn_province`;
COMMIT;