-- 003_create_cources_table.sql

-- +migrate Up
CREATE TABLE cources (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL
);

-- +migrate Down
DROP TABLE cources;
