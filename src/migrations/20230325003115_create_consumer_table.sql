-- +goose Up
START TRANSACTION;

    CREATE TABLE `consumer` (
        `pkid` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `hostname` VARCHAR(255) NOT NULL,
        `maintener_email` VARCHAR(255) NOT NULL,
        `is_active` BOOLEAN NOT NULL DEFAULT FALSE,
        `is_validated` BOOLEAN NOT NULL DEFAULT FALSE,
        `secret_key` TEXT NULL DEFAULT NULL,
        `created_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `updated_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `deleted_at` DATETIME(6) NULL DEFAULT NULL,
        PRIMARY KEY (`pkid`),
        UNIQUE (`hostname`)
    );

COMMIT;

-- +goose Down
START TRANSACTION;
    DROP TABLE `consumer`;
COMMIT;