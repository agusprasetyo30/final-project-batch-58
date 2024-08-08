package repository

import (
	"database/sql"
	"final-project/model"
)

type UserRepository interface {
	GetUserByUsernameAndPassword(username string) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByUsernameAndPassword(username string) (*model.User, error) {
	var user model.User
	sql := "SELECT * FROM users WHERE username = $1"
	err := r.db.QueryRow(sql, username).Scan(&user.ID,
		&user.Username,
		&user.Password,
		&user.Role)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(db *sql.DB, user model.User) (*model.User, error) {
	sql := "INSERT INTO users (username, password, role) VALUES ($1,$2,$3) RETURNING *"
	errs := db.QueryRow(sql, user.Username, user.Password, user.Role).Scan(&user.ID, &user.Username, &user.Password, &user.Role)

	if errs != nil {
		return nil, errs
	}

	return &user, nil
}
