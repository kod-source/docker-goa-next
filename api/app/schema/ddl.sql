CREATE TABLE IF NOT EXISTS `users` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    `password` VARCHAR(64) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `avatar` TEXT,
    UNIQUE (`email`)
);


CREATE TABLE IF NOT EXISTS `posts` (
    `id` MEDIUMINT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `title` VARCHAR(50) NOT NULL DEFAULT "",
    `img` LONGTEXT,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS `comments` (
    `id` MEDIUMINT NOT NULL AUTO_INCREMENT,
    `post_id` INT NOT NULL,
    `text` VARCHAR(50) NOT NULL DEFAULT "",
    `img` LONGTEXT,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS `likes` (
    `id` MEDIUMINT NOT NULL AUTO_INCREMENT,
    `post_id` INT NOT NULL,
    `user_id` INT NOT NULL,
    PRIMARY KEY (id),
    UNIQUE post_user_id_index (post_id, user_id)
);
