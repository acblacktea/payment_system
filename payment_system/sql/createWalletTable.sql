CREATE TABLE `Wallet` (
   `id` bigint(64) unsigned NOT NULL,
   `balance` bigint(64) unsigned NOT NULL,
   `currency` varchar(128) DEFAULT NULL,
   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   `deleted_at` timestamp NULL DEFAULT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4