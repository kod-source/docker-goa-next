SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `email` varchar(50) NOT NULL,
    `password` varchar(64) NOT NULL,
    `created_at` datetime NOT NULL,
    `avatar` text,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `post`;

CREATE TABLE `post` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `title` varchar(50) NOT NULL DEFAULT '',
    `img` longtext,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`),
    KEY `user_id_constraint` (`user_id`),
    CONSTRAINT `user_id_constraint` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
    `id` int NOT NULL AUTO_INCREMENT,
    `post_id` int NOT NULL,
    `text` varchar(50) NOT NULL DEFAULT '',
    `img` longtext,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `user_id` int NOT NULL,
    PRIMARY KEY (`id`),
    KEY `user_constraint` (`user_id`),
    KEY `post_id_constraint` (`post_id`),
    CONSTRAINT `post_id_constraint` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `user_constraint` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `like`;

CREATE TABLE `like` (
    `id` int NOT NULL AUTO_INCREMENT,
    `post_id` int NOT NULL,
    `user_id` int NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `post_user_id_index` (`post_id`,`user_id`),
    KEY `u_id_constraint` (`user_id`),
    CONSTRAINT `p_id_constraint` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `u_id_constraint` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
