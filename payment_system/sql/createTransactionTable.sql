CREATE TABLE `Transaction` (
    `id` bigint(64) unsigned NOT NULL,
    `source_account_id` bigint(64) unsigned NOT NULL,
    `destination_account_id` bigint(64) unsigned NOT NULL,
    `amount` bigint(64) unsigned NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4