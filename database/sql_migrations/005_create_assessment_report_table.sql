-- 004_create_assessment_report_table.sql

-- +migrate Up
CREATE TABLE assessment_reports (
	id SERIAL PRIMARY KEY,
	student_id INTEGER NOT NULL,
	cource_id INTEGER NOT NULL,
	score INTEGER NOT NULL,
	grade VARCHAR(10) NOT NULL,
	CONSTRAINT fk_student FOREIGN KEY (student_id) REFERENCES students (id),
	CONSTRAINT fk_cource FOREIGN KEY (cource_id) REFERENCES cources (id)
);

-- +migrate Down
DROP TABLE assessment_reports;
