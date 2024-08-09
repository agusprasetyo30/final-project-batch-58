-- 003_create_mata_pelajarans_table.sql

-- +migrate Up
CREATE TABLE students (
	id SERIAL PRIMARY KEY,
	class_id INTEGER NOT NULL,
	name VARCHAR(200) NOT NULL,
	gender VARCHAR(20) NOT NULL,
	born_date DATE NOT NULL,
	address text NOT NULL,
	CONSTRAINT fk_class FOREIGN KEY (class_id) REFERENCES classes (id)
);

-- +migrate Down
DROP TABLE students;
