
INSERT INTO `inflow`(`unixtimestamp`, `document_guid`, `name`, `amount`, `currency`, `source`) VALUES
    (5, "deadc457-a6b5-4071-911e-b79f29328262", "name1", 2.0, "RUB", "job1"),
    (2, "8e6ab86d-c53c-4b39-8c7b-5e672c6c887d", "name2", 3.0, "RUB", "job2"),
    (6, "098a399c-3e97-4ebb-be5e-a35a7cff0e6a", "name3", 4.1, "RUB", "job1"),
    (8, "de3228eb-9594-4a8e-b97b-8092a2353ba3", "name4", 5.3, "RUB", "job2");

INSERT INTO `outflow`(`unixtimestamp`, `document_guid`, `name`, `amount`, `currency`, `destination`) VALUES
    (1, "b3493958-b49e-45eb-9d53-a03ce67c25ec", "name5", 3.4, "RUB", "shop1"),
    (3, "89fb189b-03e1-40bd-b265-feecb8296005", "name6", 4.5, "RUB", "shop1"),
    (7, "b91f2dcf-ba1a-4e9a-a251-ff49ecebcdae", "name7", 5.1, "RUB", "shop2"),
    (4, "b7acba0d-7e2b-462d-a460-15bd82918085", "name8", 6.8, "RUB", "shop2");
