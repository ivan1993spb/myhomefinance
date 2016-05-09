
CREATE TABLE `notes` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `title` VARCHAR(300) NOT NULL,
    `datetime` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `text` TEXT
);
