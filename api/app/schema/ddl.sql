CREATE TABLE IF NOT EXISTS `user` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    `password` VARCHAR(64) NOT NULL,
    `created_at` DATETIME NOT NULL
);

-- INSERT INTO `user` (name, email, password, created_at) VALUES (
--     "佐藤　太郎",
--     "test@exmaple.com",
--     "Test-1234",
--     "2019-03-28 01:23:45"
-- );
