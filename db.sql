
CREATE TABLE `notes` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `title` VARCHAR(300) NOT NULL,
    `datetime` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `text` TEXT
);

CREATE TABLE `inflow` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `datetime` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `name` VARCHAR(300) NOT NULL,
    `amount` FLOAT NOT NULL,
    `currency` VARCHAR(3) NOT NULL,
    `description` TEXT,
    `source` VARCHAR(300) NOT NULL
);

CREATE TABLE `outflow` (
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
