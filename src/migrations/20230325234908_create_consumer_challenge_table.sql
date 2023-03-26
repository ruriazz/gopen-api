-- +goose Up
START TRANSACTION;

    CREATE TABLE `consumer_challenge` (
        `pkid` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `consumer_pkid` BIGINT UNSIGNED NOT NULL,
        `is_active` BOOLEAN NOT NULL DEFAULT FALSE,
        `is_verified` BOOLEAN NOT NULL DEFAULT FALSE,
        `challenge_type` VARCHAR(30) NOT NULL DEFAULT 'dns',
        `challenge_value` TEXT NULL DEFAULT NULL,
        `expired_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `created_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `updated_at` DATETIME(6) NOT NULL DEFAULT NOW(6),
        `deleted_at` DATETIME(6) NULL DEFAULT NULL,
        PRIMARY KEY (`pkid`),
        FOREIGN KEY (`consumer_pkid`) REFERENCES `consumer`(`pkid`)
    );

COMMIT;

-- +goose Down
START TRANSACTION;
    DROP TABLE `consumer_challenge`;
COMMIT;