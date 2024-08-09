package model

type Student struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	BornDate string `json:"born_date"`
	Address  string `json:"address"`
	ClassID  int    `json:"class_id"`
	Class    Class  `json:"class" gorm:"foreignKey:ClassID"`
}
