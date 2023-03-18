-- +goose Up
START TRANSACTION;

    CREATE TABLE `idn_subdistrict` (
        `pkid` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `idn_district_pkid` SMALLINT UNSIGNED NOT NULL,
        `slug` VARCHAR(75) NULL DEFAULT NULL,
        `code` VARCHAR(12) NULL DEFAULT NULL,
        `name` VARCHAR(75) NOT NULL,
        `alias` VARCHAR(75) NULL DEFAULT NULL,
        `created_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `updated_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `deleted_at` DATETIME(6) NULL DEFAULT NULL,
        PRIMARY KEY (`pkid`),
        UNIQUE (`slug`, `code`),
        FOREIGN KEY (`idn_district_pkid`) REFERENCES `idn_district`(`pkid`)
    );

COMMIT;

-- +goose Down
START TRANSACTION;
    DROP TABLE `idn_subdistrict`;
COMMIT;