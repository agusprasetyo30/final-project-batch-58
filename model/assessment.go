package model

type Assessment struct {
	ID        int     `json:"id" gorm:"primary_key"`
	Score     int     `json:"score"`
	Grade     string  `json:"grade"`
	StudentID int     `json:"student_id"`
	Student   Student `json:"student" gorm:"foreignKey:StudentID"`
	CourceID  int     `json:"cource_id"`
	Cource    Cource  `json:"cource" gorm:"foreignKey:CourceID"`
}
