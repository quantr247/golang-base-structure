package repository

import (
	"context"
	"database/sql"
	"golang-base-structure/internal/helper/db"
	"golang-base-structure/internal/repository/model"
	"strconv"
	"strings"
)

type (
	PostgresRepository interface {
		ListUser(ctx context.Context) ([]model.User, error)
		GetUser(ctx context.Context, userName string) (model.User, error)
		SaveUser(ctx context.Context, user model.User) (string, error)
		UpdateUser(ctx context.Context, user model.User) error
	}

	postgresRepository struct {
		postgresDB db.DBHelper
	}
)

func NewPostgresRepository(postgresDB db.DBHelper) PostgresRepository {
	return &postgresRepository{
		postgresDB: postgresDB,
	}
}

func (r *postgresRepository) ListUser(ctx context.Context) (users []model.User, err error) {
	var statement strings.Builder
	statement.WriteString(`SELECT
							id
							, username
							, phonenumber
							FROM user
							`)

	db := r.postgresDB.Open()
	rows, err := db.Query(statement.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err = rows.Scan(
			&user.ID,
			&user.UserName,
			&user.PhoneNumber,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, sql.ErrNoRows
	}

	return users, nil
}

func (r *postgresRepository) GetUser(ctx context.Context, userName string) (user model.User, err error) {
	var (
		statement strings.Builder
		args      []interface{}
	)
	statement.WriteString(`SELECT 
							id
							, username
							, phonenumber
						FROM user 
						WHERE username = $1`)
	args = append(args, userName)
	user = model.User{}
	err = r.postgresDB.Open().QueryRow(statement.String(), args...).
		Scan(
			&user.ID,
			&user.UserName,
			&user.PhoneNumber,
		)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *postgresRepository) SaveUser(ctx context.Context, user model.User) (returnID string, err error) {
	var (
		args      []interface{}
		statement = `INSERT INTO user 
						(username,
						phonenumber)
					VALUES ($1, $2)
					RETURNING id`
	)

	tx, errDB := r.postgresDB.Begin()
	defer func() {
		if errDB != nil {
			_ = r.postgresDB.Rollback(tx)
		} else {
			_ = r.postgresDB.Commit(tx)
		}
	}()

	args = append(args, user.UserName, user.PhoneNumber)

	id := 0
	errDB = tx.QueryRow(statement, args...).Scan(&id)

	if errDB != nil {
		return "", errDB
	}
	return strconv.Itoa(id), nil
}

func (r *postgresRepository) UpdateUser(ctx context.Context, user model.User) (err error) {
	var (
		statement = `UPDATE user 
			SET username=$1
			WHERE phonenumber=$2`
	)

	tx, errDB := r.postgresDB.Begin()
	defer func() {
		if errDB != nil {
			_ = r.postgresDB.Rollback(tx)
		} else {
			_ = r.postgresDB.Commit(tx)
		}
	}()

	_, errDB = tx.Exec(statement, user.UserName, user.PhoneNumber)

	if errDB != nil {
		return errDB
	}
	return nil
}
