package repository

import (
	"database/sql"
	"errors"
	"final-project/model"
)

func GetAllAssessment(db *sql.DB) (result []model.Assessment, err error) {
	sql := "SELECT * FROM assessment_reports ORDER BY id ASC"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var assessment model.Assessment

		err = rows.Scan(&assessment.ID, &assessment.StudentID, &assessment.CourceID, &assessment.Score, &assessment.Grade)
		if err != nil {
			return
		}

		// untuk menampilkan students, classes & cources di relasi
		studentQuery := "SELECT * FROM students WHERE id = $1"
		err = db.QueryRow(studentQuery, assessment.StudentID).Scan(&assessment.Student.ID, &assessment.Student.ClassID, &assessment.Student.Name,
			&assessment.Student.Gender, &assessment.Student.BornDate, &assessment.Student.Address)
		if err != nil {
			return
		}

		classQuery := "SELECT * FROM classes WHERE id = $1"
		err = db.QueryRow(classQuery, assessment.Student.ClassID).Scan(&assessment.Student.Class.ID, &assessment.Student.Class.Number, &assessment.Student.Class.Class_type)
		if err != nil {
			return
		}

		courceQuery := "SELECT * FROM cources WHERE id = $1"
		err = db.QueryRow(courceQuery, assessment.CourceID).Scan(&assessment.Cource.ID, &assessment.Cource.Name)
		if err != nil {
			return
		}

		result = append(result, assessment)
	}

	return
}

func GetAssessment(db *sql.DB, assessment model.Assessment) (*model.Assessment, error) {
	sql := "SELECT * FROM assessment_reports WHERE id = $1"

	errs := db.QueryRow(sql, assessment.ID).Scan(&assessment.ID, &assessment.StudentID, &assessment.CourceID, &assessment.Score, &assessment.Grade)
	if errs != nil {
		return nil, errs
	}

	// untuk menampilkan students, classes & cources di relasi
	studentQuery := "SELECT * FROM students WHERE id = $1"
	errs = db.QueryRow(studentQuery, assessment.StudentID).Scan(&assessment.Student.ID, &assessment.Student.ClassID, &assessment.Student.Name,
		&assessment.Student.Gender, &assessment.Student.BornDate, &assessment.Student.Address)
	if errs != nil {
		return nil, errs
	}

	classQuery := "SELECT * FROM classes WHERE id = $1"
	errs = db.QueryRow(classQuery, assessment.Student.ClassID).Scan(&assessment.Student.Class.ID, &assessment.Student.Class.Number, &assessment.Student.Class.Class_type)
	if errs != nil {
		return nil, errs
	}

	courceQuery := "SELECT * FROM cources WHERE id = $1"
	errs = db.QueryRow(courceQuery, assessment.CourceID).Scan(&assessment.Cource.ID, &assessment.Cource.Name)
	if errs != nil {
		return nil, errs
	}

	return &assessment, nil
}

func InsertAssessment(db *sql.DB, assessment model.Assessment) (*model.Assessment, error) {
	sql := "INSERT INTO assessment_reports (student_id,cource_id,score,grade) VALUES ($1,$2,$3,$4) RETURNING id"
	errs := db.QueryRow(sql, assessment.StudentID, assessment.CourceID, assessment.Score, assessment.Grade).Scan(&assessment.ID)

	if errs != nil {
		return nil, errs
	}

	// untuk menampilkan students di relasi
	studentQuery := "SELECT * FROM students WHERE id = $1"
	errs = db.QueryRow(studentQuery, assessment.StudentID).Scan(&assessment.Student.ID, &assessment.Student.ClassID, &assessment.Student.Name,
		&assessment.Student.Gender, &assessment.Student.BornDate, &assessment.Student.Address)
	if errs != nil {
		return nil, errs
	}

	classQuery := "SELECT * FROM classes WHERE id = $1"
	errs = db.QueryRow(classQuery, assessment.Student.ClassID).Scan(&assessment.Student.Class.ID, &assessment.Student.Class.Number, &assessment.Student.Class.Class_type)
	if errs != nil {
		return nil, errs
	}

	courceQuery := "SELECT * FROM cources WHERE id = $1"
	errs = db.QueryRow(courceQuery, assessment.CourceID).Scan(&assessment.Cource.ID, &assessment.Cource.Name)
	if errs != nil {
		return nil, errs
	}

	return &assessment, nil
}

func UpdateAssessment(db *sql.DB, assessment model.Assessment) error {
	sql := "UPDATE assessment_reports SET student_id = $1, cource_id = $2, score = $3, grade = $4 WHERE id = $5"

	errs := db.QueryRow(sql, assessment.StudentID, assessment.CourceID, assessment.Score, assessment.Grade, assessment.ID)

	return errs.Err()
}

func DeleteAssessment(db *sql.DB, assessment model.Assessment) error {
	sql := "DELETE FROM assessment_reports WHERE id = $1"

	result, err := db.Exec(sql, assessment.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("assessment not found")
	}

	return nil
}

// // untuk menampilkan students, classes & cources di relasi
// studentQuery := "SELECT * FROM students WHERE id = $1"
// err = db.QueryRow(studentQuery, assessment.StudentID).Scan(&assessment.Student.ID, &assessment.Student.ClassID, &assessment.Student.Name,
// 	&assessment.Student.Gender, &assessment.Student.BornDate, &assessment.Student.Address)
// if err != nil {
// 	return
// }

// classQuery := "SELECT * FROM classes WHERE id = $1"
// err = db.QueryRow(classQuery, assessment.Student.ClassID).Scan(&assessment.Student.Class.ID, &assessment.Student.Class.Number, &assessment.Student.Class.Class_type)
// if err != nil {
// 	return
// }

// courceQuery := "SELECT * FROM cources WHERE id = $1"
// err = db.QueryRow(courceQuery, assessment.CourceID).Scan(&assessment.Cource.ID, &assessment.Cource.Name)
// if err != nil {
// 	return
// }
