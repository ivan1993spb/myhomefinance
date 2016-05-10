
CREATE TABLE IF NOT EXISTS `notes` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `title` VARCHAR(300) NOT NULL,
    `datetime` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `text` TEXT
);

CREATE TABLE IF NOT EXISTS `inflow` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `datetime` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `name` VARCHAR(300) NOT NULL,
    `amount` FLOAT NOT NULL,
    `currency` VARCHAR(3) NOT NULL,
    `description` TEXT,
    `source` VARCHAR(300) NOT NULL
);

CREATE TABLE IF NOT EXISTS `outflow` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `datetime` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `name` VARCHAR(300) NOT NULL,
    `amount` FLOAT NOT NULL,
    `currency` VARCHAR(3) NOT NULL,
    `description` TEXT,
    `destination` VARCHAR(300) NOT NULL,
    `target` VARCHAR(300),
    `count` FLOAT,
    `metric_unit` VARCHAR(100),
    `satisfaction` FLOAT
);

CREATE VIEW IF NOT EXISTS `transactions` AS
    SELECT * FROM
        (SELECT `datetime`, `name`, `amount`, `currency` FROM `inflow`
            UNION SELECT `datetime`, `name`, -`amount`, `currency` FROM `outflow`)
        ORDER BY `datetime` DESC;