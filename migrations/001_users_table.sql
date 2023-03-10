-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(36) NOT NULL,
  `email` varchar(64) NOT NULL,
  `password` varchar(32) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
--- SQL in this section is executed when the migration is rolled back.
