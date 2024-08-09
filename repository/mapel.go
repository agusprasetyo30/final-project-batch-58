package repository

import (
	"database/sql"
	"errors"
	"final-project/model"
)

func GetAllMapel(db *sql.DB) (result []model.Mapel, err error) {
	sql := "SELECT * FROM mata_pelajarans ORDER BY id ASC"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var mapel model.Mapel

		err = rows.Scan(&mapel.ID, &mapel.Name)
		if err != nil {
			return
		}

		result = append(result, mapel)
	}

	return
}

func GetMapel(db *sql.DB, mapel model.Mapel) (*model.Mapel, error) {
	sql := "SELECT * FROM mata_pelajarans WHERE id = $1"

	errs := db.QueryRow(sql, mapel.ID).Scan(&mapel.ID, &mapel.Name)

	if errs != nil {
		return nil, errs
	}

	return &mapel, nil
}

func InsertMapel(db *sql.DB, mapel model.Mapel) (*model.Mapel, error) {
	sql := "INSERT INTO mata_pelajarans (name) VALUES ($1) RETURNING *"
	errs := db.QueryRow(sql, mapel.Name).Scan(&mapel.ID, &mapel.Name)

	if errs != nil {
		return nil, errs
	}

	return &mapel, nil
}

func UpdateMapel(db *sql.DB, mapel model.Mapel) error {
	sql := "UPDATE mata_pelajarans SET name = $1 WHERE id = $2"

	errs := db.QueryRow(sql, mapel.Name, mapel.ID)

	return errs.Err()
}

func DeleteMapel(db *sql.DB, mapel model.Mapel) error {
	sql := "DELETE FROM mata_pelajarans WHERE id = $1"

	result, err1 := db.Exec(sql, mapel.ID)
	if err1 != nil {
		return err1
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("mata pelajaran not found")
	}

	return nil
}
