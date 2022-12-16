SET foreign_key_checks=0;

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(191) NOT NULL,
    `email` VARCHAR(191) NOT NULL,
    `password` VARCHAR(191) NOT NULL,
    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,
    `avatar` LONGTEXT NULL,
    UNIQUE `idx_email` (`email`),
    PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `post`;

CREATE TABLE `post` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `title` VARCHAR(191) NOT NULL,
    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,
    `img` LONGTEXT NULL,
    INDEX `idx_user_id` (`user_id`),
    CONSTRAINT `user_id_constraint` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `post_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `text` VARCHAR(191) NOT NULL,
    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,
    `img` LONGTEXT NULL,
    INDEX `idx_post_id` (`post_id`),
    INDEX `idx_user_id` (`user_id`),
    CONSTRAINT `user_constraint` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `post_id_constraint` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `like`;

CREATE TABLE `like` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `post_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_post_id` (`post_id`),
    UNIQUE `post_user_id_index` (`post_id`, `user_id`),
    CONSTRAINT `u_id_constraint` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `p_id_constraint` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `room`;

CREATE TABLE `room` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(191) NOT NULL,
    `is_group` TINYINT(1) NOT NULL,
    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,
    PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `user_room`;

CREATE TABLE `user_room` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `room_id` BIGINT UNSIGNED NOT NULL,
    `last_read_at` DATETIME(6) NULL,
    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_room_id` (`room_id`),
    UNIQUE `user_room_id_index` (`user_id`, `room_id`),
    CONSTRAINT `user_room_id_constraint` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `room_id_constraint` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `thread`;

CREATE TABLE `thread` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `room_id` BIGINT UNSIGNED NOT NULL,
    `text` VARCHAR(191) NOT NULL,
    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,
    `img` LONGTEXT NULL,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_room_id` (`room_id`),
    UNIQUE `user_room_id_index` (`user_id`, `room_id`),
    CONSTRAINT `user_thread_id_constraint` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `thread_room_id_constraint` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `thread_id` BIGINT UNSIGNED NOT NULL,
    `text` VARCHAR(191) NOT NULL,
    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,
    `img` LONGTEXT NULL,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_thread_id` (`thread_id`),
    UNIQUE `user_thread_id_index` (`user_id`, `thread_id`),
    CONSTRAINT `user_message_id_constraint` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `thread_message_id_constraint` FOREIGN KEY (`thread_id`) REFERENCES `thread` (`id`) ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
);

SET foreign_key_checks=1;
