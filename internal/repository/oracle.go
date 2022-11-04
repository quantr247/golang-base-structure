package repository

import (
	"context"
	"database/sql"
	"golang-base-structure/internal/helper/db"
	"golang-base-structure/internal/repository/model"
	"strings"
)

type (
	OracleRepository interface {
		ListApplication(ctx context.Context) ([]*model.Application, error)
		GetApplication(ctx context.Context, code string) (model.Application, error)
		SaveApplication(ctx context.Context, app model.Application) error
		UpdateApplication(ctx context.Context, app model.Application) error
		ExecutePackage(ctx context.Context, app model.Application) (string, error)
	}

	oracleRepository struct {
		oracleDB db.DBHelper
	}
)

func NewOracleRepository(oracleDB db.DBHelper) OracleRepository {
	return &oracleRepository{
		oracleDB: oracleDB,
	}
}

func (r *oracleRepository) ListApplication(ctx context.Context) (apps []*model.Application, err error) {
	var statement strings.Builder
	statement.WriteString(`SELECT
					code
					, name
				FROM View_Name
				`)

	db := r.oracleDB.Open()
	rows, err := db.Query(statement.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var app *model.Application
		err = rows.Scan(
			&app.Code,
			&app.Name,
		)

		if err != nil {
			return nil, err
		}

		apps = append(apps, app)
	}

	if len(apps) == 0 {
		return nil, sql.ErrNoRows
	}

	return apps, nil
}

func (r *oracleRepository) GetApplication(ctx context.Context, code string) (app model.Application, err error) {
	var (
		statement strings.Builder
		args      []interface{}
	)
	statement.WriteString(`SELECT 
							name
							, code 
						FROM application 
						WHERE code = :1`)
	args = append(args, code)
	err = r.oracleDB.Open().QueryRow(statement.String(), args...).
		Scan(
			&app.Name,
			&app.Code,
		)
	if err != nil {
		return app, err
	}
	return app, nil
}

func (r *oracleRepository) SaveApplication(ctx context.Context, app model.Application) (err error) {
	var (
		statement = `INSERT INTO application 
			(name,
			code)
		VALUES (:1, :2)`
	)

	tx, errDB := r.oracleDB.Begin()
	defer func() {
		if errDB != nil {
			_ = r.oracleDB.Rollback(tx)
		} else {
			_ = r.oracleDB.Commit(tx)
		}
	}()

	_, errDB = tx.Exec(statement, app.Name, app.Code)

	if errDB != nil {
		return errDB
	}
	return nil
}

func (r *oracleRepository) UpdateApplication(ctx context.Context, app model.Application) (err error) {
	var (
		statement = `UPDATE application 
			SET name=:1
			WHERE code=:2`
	)

	tx, errDB := r.oracleDB.Begin()
	defer func() {
		if errDB != nil {
			_ = r.oracleDB.Rollback(tx)
		} else {
			_ = r.oracleDB.Commit(tx)
		}
	}()

	_, errDB = tx.Exec(statement, app.Name, app.Code)

	if errDB != nil {
		return errDB
	}
	return nil
}

func (r *oracleRepository) ExecutePackage(ctx context.Context, app model.Application) (result string, err error) {
	var (
		statement = `BEGIN PACKAGE.APPLICATION(:1,:2,:3); END;`
	)

	db := r.oracleDB.Open()
	_, err = db.Exec(statement, app.Name, app.Code, sql.Out{Dest: &result})
	if err != nil {
		return "", err
	}

	return result, nil
}
