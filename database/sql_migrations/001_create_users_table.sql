-- 001_create_users_table.sql

-- +migrate Up
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	role VARCHAR(20) NULL
);

-- +migrate Down
DROP TABLE users;
