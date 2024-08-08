package model

type Class struct {
	ID         int    `json:"id" gorm:"primary_key"`
	Number     int    `json:"number"`
	Class_type string `json:"class_type"`
}
