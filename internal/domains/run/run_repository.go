package run

import (
	"fmt"
	"go-mini/internal/response/failure"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RunRepository struct {
	DB     *sqlx.DB
	Logger *logrus.Entry
}

func NewRunRepository(db *sqlx.DB, logger *logrus.Entry) *RunRepository {
	return &RunRepository{DB: db, Logger: logger}
}

func (r *RunRepository) Create(load Run) (err error) {
	query := `INSERT INTO run (id,user_id,run_time,run_kilometers,run_city,created_at,updated_at,created_by,updated_by)
    VALUES (:id,:user_id,:run_time,:run_kilometers,:run_city,:created_at,:updated_at,:created_by,:updated_by)`

	stmt, err := r.DB.PrepareNamed(query)

	if err != nil {
		r.Logger.Error(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(load)
	if err != nil {
		r.Logger.Error(err)
		return
	}

	return
}

func (r *RunRepository) GetAll(limit, offset int, sort, field string, location string) (res []Run, err error) {
	query := `SELECT * FROM run `

	if location != "" {
		query += fmt.Sprintf(`WHERE run_city = '%s' `, location)
	}

	query += fmt.Sprintf(" ORDER BY %s %s LIMIT %d OFFSET %d", field, sort, limit, offset)
	err = r.DB.Select(&res, query)
	if err != nil {
		println(err.Error())
		println(query)
		err = failure.InternalError(err)
		r.Logger.Error(err)
	}
	return
}

func (r *RunRepository) Update(load Run) (err error) {
	query := `UPDATE run
    SET
        run_time = :run_time,
        run_kilometers = :run_kilometers,
        run_city = :run_city,
        updated_at = :updated_at,
        updated_by = :updated_by,
        deleted_at = :deleted_at,
        deleted_by = :deleted_by
    WHERE id = :id
    `

	stmt, err := r.DB.PrepareNamed(query)

	if err != nil {
		r.Logger.Error(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(load)
	if err != nil {
		r.Logger.Error(err)
		return
	}

	return
}

func (r *RunRepository) GetByID(id string) (res Run, err error) {
	err = r.DB.Get(&res, "SELECT * FROM run WHERE id = ?", id)
	if err != nil {
		err = failure.InternalError(err)
		r.Logger.Error(err)
		return
	}
	return
}

func (r *RunRepository) Delete(id string) (err error) {
	_, err = r.DB.Exec("DELETE FROM run WHERE id = ?", id)
	return
}
