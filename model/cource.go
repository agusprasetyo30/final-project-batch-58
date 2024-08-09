package model

type Cource struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
