CREATE TABLE IF NOT EXISTS `users` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    `password` VARCHAR(64) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `avatar` TEXT,
    UNIQUE (`email`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `posts` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `title` VARCHAR(50) NOT NULL DEFAULT "",
    `img` LONGTEXT,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`)
    REFERENCES `users`(`id`)
    ON UPDATE CASCADE ON DELETE CASCADE
)  ENGINE=InnoDB  DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `comments` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `post_id` INT NOT NULL,
    `user_id` INT NOT NULL,
    `text` VARCHAR(50) NOT NULL DEFAULT "",
    `img` LONGTEXT,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),

    FOREIGN KEY (`post_id`)
    REFERENCES `posts`(`id`)
    ON UPDATE CASCADE ON DELETE CASCADE,

    FOREIGN KEY (`user_id`)
    REFERENCES `users`(`id`)
    ON UPDATE CASCADE ON DELETE CASCADE
)  ENGINE=InnoDB  DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `likes` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `post_id` INT NOT NULL,
    `user_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE `post_user_id_index` (`post_id`, `user_id`),

    INDEX (`post_id`),
    INDEX (`user_id`),

    FOREIGN KEY (`post_id`)
    REFERENCES `posts`(`id`)
    ON UPDATE CASCADE ON DELETE CASCADE,

    FOREIGN KEY (`user_id`)
    REFERENCES `users`(`id`)
    ON UPDATE CASCADE ON DELETE CASCADE
)  ENGINE=InnoDB  DEFAULT CHARSET=utf8;
