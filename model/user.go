package model

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
