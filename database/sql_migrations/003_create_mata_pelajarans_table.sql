-- 003_create_mata_pelajarans_table.sql

-- +migrate Up
CREATE TABLE mata_pelajarans (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL
);

-- +migrate Down
DROP TABLE mata_pelajarans;
