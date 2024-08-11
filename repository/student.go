package repository

import (
	"database/sql"
	"errors"
	"final-project/model"
)

func GetAllStudents(db *sql.DB) (result []model.Student, err error) {
	sql := "SELECT * FROM students ORDER BY class_id ASC"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var student model.Student

		err = rows.Scan(&student.ID, &student.ClassID, &student.Name, &student.Gender, &student.BornDate, &student.Address)
		if err != nil {
			return
		}

		// untuk menampilkan class di relasi
		classQuery := "SELECT * FROM classes WHERE id = $1"
		err = db.QueryRow(classQuery, student.ClassID).Scan(&student.Class.ID, &student.Class.Number, &student.Class.Class_type)
		if err != nil {
			return
		}

		result = append(result, student)
	}

	return
}

func GetStudent(db *sql.DB, student model.Student) (*model.Student, error) {
	sql := "SELECT * FROM students WHERE id = $1"

	errs := db.QueryRow(sql, student.ID).Scan(&student.ID, &student.ClassID, &student.Name, &student.Gender, &student.BornDate, &student.Address)
	if errs != nil {
		return nil, errs
	}

	// untuk menampilkan class di relasi
	classQuery := "SELECT * FROM classes WHERE id = $1"
	errs = db.QueryRow(classQuery, student.ClassID).Scan(&student.Class.ID, &student.Class.Number, &student.Class.Class_type)
	if errs != nil {
		return nil, errs
	}

	return &student, nil
}

func InsertStudent(db *sql.DB, student model.Student) (*model.Student, error) {
	sql := "INSERT INTO students (class_id,name,gender,born_date,address) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	errs := db.QueryRow(sql, student.ClassID, student.Name, student.Gender, student.BornDate, student.Address).Scan(&student.ID)

	if errs != nil {
		return nil, errs
	}

	// untuk menampilkan class di relasi
	classQuery := "SELECT * FROM classes WHERE id = $1"
	errs = db.QueryRow(classQuery, student.ClassID).Scan(&student.Class.ID, &student.Class.Number, &student.Class.Class_type)

	if errs != nil {
		return nil, errs
	}

	return &student, nil
}

func UpdateStudent(db *sql.DB, student model.Student) error {
	sql := "UPDATE students SET class_id = $1, name = $2, gender = $3, born_date = $4, address = $5 WHERE id = $6"

	errs := db.QueryRow(sql, student.ClassID, student.Name, student.Gender, student.BornDate, student.Address, student.ID)

	return errs.Err()
}

func DeleteStudent(db *sql.DB, student model.Student) error {
	sql := "DELETE FROM students WHERE id = $1"

	result, err := db.Exec(sql, student.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("student not found")
	}

	return nil
}
