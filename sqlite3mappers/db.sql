
CREATE TABLE IF NOT EXISTS `notes` (
    `id`            INTEGER PRIMARY KEY AUTOINCREMENT,
    `name`          VARCHAR(300) NOT NULL,
    `unixtimestamp` INT8 DEFAULT (strftime('%s', 'now')) NOT NULL,
    `text`          TEXT
);

CREATE TABLE IF NOT EXISTS `inflow` (
    `id`            INTEGER PRIMARY KEY AUTOINCREMENT,
    `document_guid` VARCHAR(36) NOT NULL,
    `unixtimestamp` INT8 DEFAULT (strftime('%s', 'now')) NOT NULL,
    `name`          VARCHAR(300) NOT NULL,
    `amount`        DOUBLE NOT NULL,
    `description`   TEXT,
    `source`        VARCHAR(300) NOT NULL
);

CREATE TABLE IF NOT EXISTS `outflow` (
    `id`            INTEGER PRIMARY KEY AUTOINCREMENT,
    `document_guid` VARCHAR(36) NOT NULL,
    `unixtimestamp` INT8 DEFAULT (strftime('%s', 'now')) NOT NULL,
    `name`          VARCHAR(300) NOT NULL,
    `amount`        DOUBLE NOT NULL,
    `description`   TEXT,
    `destination`   VARCHAR(300) NOT NULL,
    `target`        VARCHAR(300),
    `count`         DOUBLE,
    `metric_unit`   VARCHAR(100),
    `satisfaction`  FLOAT
);

CREATE VIEW IF NOT EXISTS `transactions` AS
    SELECT * FROM (
        SELECT `document_guid`, `unixtimestamp`, `name`, `amount`, `description` FROM `inflow`
        UNION
        SELECT `document_guid`, `unixtimestamp`, `name`, -`amount` AS `amount`, `description` FROM `outflow`
    ) `result_union`
    ORDER BY (`result_union`.`unixtimestamp`) DESC;