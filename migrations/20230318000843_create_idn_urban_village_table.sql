-- +goose Up
START TRANSACTION;

    CREATE TABLE `idn_urban_village` (
        `pkid` INT UNSIGNED NOT NULL AUTO_INCREMENT,
        `idn_subdistrict_pkid` SMALLINT UNSIGNED NOT NULL,
        `slug` VARCHAR(75) NULL DEFAULT NULL,
        `code` VARCHAR(15) NULL DEFAULT NULL,
        `name` VARCHAR(75) NOT NULL,
        `alias` VARCHAR(75) NULL DEFAULT NULL,
        `postal_code` VARCHAR(10) NULL DEFAULT NULL,
        `created_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `updated_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `deleted_at` DATETIME(6) NULL DEFAULT NULL,
        PRIMARY KEY (`pkid`),
        UNIQUE (`slug`, `code`),
        FOREIGN KEY (`idn_subdistrict_pkid`) REFERENCES `idn_subdistrict`(`pkid`)
    );

COMMIT;

-- +goose Down
START TRANSACTION;
    DROP TABLE `idn_urban_village`;
COMMIT;