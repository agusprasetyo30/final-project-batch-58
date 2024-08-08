-- 002_create_classes_table.sql

-- +migrate Up
CREATE TABLE classes (
	id SERIAL PRIMARY KEY,
	number INT NOT NULL,
	class_type VARCHAR(100) NOT NULL
);

-- +migrate Down
DROP TABLE classes;
