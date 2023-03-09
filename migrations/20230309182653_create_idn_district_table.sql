-- +goose Up
START TRANSACTION;

    CREATE TABLE `idn_district` (
        `pkid` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `idn_province_pkid` TINYINT UNSIGNED NOT NULL,
        `slug` VARCHAR(75) NULL DEFAULT NULL,
        `code` VARCHAR(5) NULL DEFAULT NULL,
        `dati_name` VARCHAR(20) NOT NULL,
        `name` VARCHAR(75) NOT NULL,
        `alias` VARCHAR(75) NULL DEFAULT NULL,
        `created_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `updated_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `deleted_at` DATETIME(6) NULL DEFAULT NULL,
        PRIMARY KEY (`pkid`),
        UNIQUE (`slug`, `code`),
        FOREIGN KEY (`idn_province_pkid`) REFERENCES `idn_province`(`pkid`)
    );

COMMIT;

-- +goose Down
START TRANSACTION;
    DROP TABLE `idn_province`;
COMMIT;