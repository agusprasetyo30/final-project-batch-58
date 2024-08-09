package repository

import (
	"database/sql"
	"errors"
	"final-project/model"
)

func GetAllCource(db *sql.DB) (result []model.Cource, err error) {
	sql := "SELECT * FROM cources ORDER BY id ASC"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var cource model.Cource

		err = rows.Scan(&cource.ID, &cource.Name)
		if err != nil {
			return
		}

		result = append(result, cource)
	}

	return
}

func GetCource(db *sql.DB, cource model.Cource) (*model.Cource, error) {
	sql := "SELECT * FROM cources WHERE id = $1"

	errs := db.QueryRow(sql, cource.ID).Scan(&cource.ID, &cource.Name)

	if errs != nil {
		return nil, errs
	}

	return &cource, nil
}

func InsertCource(db *sql.DB, cource model.Cource) (*model.Cource, error) {
	sql := "INSERT INTO cources (name) VALUES ($1) RETURNING *"
	errs := db.QueryRow(sql, cource.Name).Scan(&cource.ID, &cource.Name)

	if errs != nil {
		return nil, errs
	}

	return &cource, nil
}

func UpdateCource(db *sql.DB, cource model.Cource) error {
	sql := "UPDATE cources SET name = $1 WHERE id = $2"

	errs := db.QueryRow(sql, cource.Name, cource.ID)

	return errs.Err()
}

func DeleteCource(db *sql.DB, cource model.Cource) error {
	sql := "DELETE FROM cources WHERE id = $1"

	result, err1 := db.Exec(sql, cource.ID)
	if err1 != nil {
		return err1
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("cources not found")
	}

	return nil
}
