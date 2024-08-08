package repository

import (
	"database/sql"
	"errors"
	"final-project/model"
)

func GetAllClasses(db *sql.DB) (result []model.Class, err error) {
	sql := "SELECT * FROM classes ORDER BY id ASC"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var class model.Class

		err = rows.Scan(&class.ID, &class.Number, &class.Class_type)
		if err != nil {
			return
		}

		result = append(result, class)
	}

	return
}

func GetClass(db *sql.DB, class model.Class) (*model.Class, error) {
	sql := "SELECT * FROM classes WHERE id = $1"

	errs := db.QueryRow(sql, class.ID).Scan(&class.ID,
		&class.Number, &class.Class_type)

	if errs != nil {
		return nil, errs
	}

	return &class, nil
}

func InsertClass(db *sql.DB, class model.Class) (*model.Class, error) {
	sql := "INSERT INTO classes (number, class_type) VALUES ($1,$2) RETURNING *"
	errs := db.QueryRow(sql, class.Number, class.Class_type).Scan(&class.ID, &class.Number, &class.Class_type)

	if errs != nil {
		return nil, errs
	}

	return &class, nil
}

func UpdateClass(db *sql.DB, class model.Class) error {
	sql := "UPDATE classes SET number = $1, class_type = $2 WHERE id = $3"

	errs := db.QueryRow(sql, class.Number, class.Class_type, class.ID)

	return errs.Err()
}

func DeleteClass(db *sql.DB, class model.Class) error {
	sql := "DELETE FROM classes WHERE id = $1"

	result, err1 := db.Exec(sql, class.ID)
	if err1 != nil {
		return err1
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("classes not found")
	}

	return nil
}
