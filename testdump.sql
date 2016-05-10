
INSERT INTO `inflow`(`datetime`, `name`, `amount`, `currency`, `source`) VALUES
    ("2016-05-7 15:30:46", "name1", 2.0, "RUB", "job1"),
    ("2016-05-8 15:30:46", "name2", 3.0, "RUB", "job2"),
    ("2016-05-9 15:30:46", "name3", 4.1, "RUB", "job1"),
    ("2016-05-10 15:30:46", "name4", 5.3, "RUB", "job2");

INSERT INTO `outflow`(`datetime`, `name`, `amount`, `currency`, `destination`) VALUES
    ("2016-05-6 15:20:46", "name5", 3.4, "RUB", "shop1"),
    ("2016-05-7 15:20:46", "name6", 4.5, "RUB", "shop1"),
    ("2016-05-9 15:20:46", "name7", 5.1, "RUB", "shop2"),
    ("2016-05-10 15:20:46", "name8", 6.8, "RUB", "shop2");
