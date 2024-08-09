package model

type Mapel struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
