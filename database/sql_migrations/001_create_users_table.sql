-- 001_create_users_table.sql

-- +migrate Up
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	role VARCHAR(20) NULL
);

INSERT INTO users (username, password, role) VALUES
    ('admin', 'jGl25bVBBBW96Qi9Te4V37Fnqchz/Eu4qB9vKrRIqRg=', 'ADMIN'), -- Password : admin
    ('user', 'CgQblGLKpKMbrDVn4Lbm/ZEAeH2yq0M9lvbReMq/zpA=', 'USER'); -- Password : user

-- +migrate Down
DROP TABLE users;
