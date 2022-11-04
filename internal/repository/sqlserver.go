package repository

import (
	"context"
	"database/sql"
	"golang-base-structure/internal/helper/db"
	"golang-base-structure/internal/repository/model"
	"strings"
)

type (
	SQLServerRepository interface {
		ListTransaction(ctx context.Context, userID string) ([]model.Transaction, error)
		GetTransaction(ctx context.Context, transactionID string) (model.Transaction, error)
		SaveTransaction(ctx context.Context, trans model.Transaction) (int64, error)
		UpdateTransaction(ctx context.Context, trans model.Transaction) error
	}

	sqlServerRepository struct {
		sqlServerDB db.DBHelper
	}
)

func NewSQLServerRepository(sqlServerDB db.DBHelper) SQLServerRepository {
	return &sqlServerRepository{
		sqlServerDB: sqlServerDB,
	}
}

func (r *sqlServerRepository) ListTransaction(ctx context.Context, userID string) (
	transactions []model.Transaction, err error) {
	var (
		statement strings.Builder
	)
	statement.WriteString(`SELECT
							id 
							, amount
							, transaction_type
							, user_id
						FROM transaction 
						WHERE user_id = @userID
							AND state = 'A'`)
	db := r.sqlServerDB.Open()
	rows, err := db.Query(statement.String(), sql.Named("userID", userID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tran model.Transaction
		err = rows.Scan(
			&tran.ID,
			&tran.Amount,
			&tran.TransactionType,
			&tran.UserID,
		)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, tran)
	}

	if len(transactions) == 0 {
		return nil, sql.ErrNoRows
	}

	return transactions, nil
}

func (r *sqlServerRepository) GetTransaction(ctx context.Context, transactionID string) (tran model.Transaction, err error) {
	var (
		statement strings.Builder
	)
	statement.WriteString(`SELECT 
							amount
							, transaction_type
							, user_id
						FROM transaction
						WHERE id = @tranID
							`)
	tran = model.Transaction{}
	err = r.sqlServerDB.Open().QueryRow(statement.String(), sql.Named("tranID", transactionID)).
		Scan(
			&tran.Amount,
			&tran.TransactionType,
			&tran.UserID,
		)
	if err != nil {
		return tran, err
	}
	return tran, nil
}

func (r *sqlServerRepository) SaveTransaction(ctx context.Context, tran model.Transaction) (id int64, err error) {
	var (
		statement strings.Builder
	)
	statement.WriteString(`INSERT INTO transaction(
								amount
								, transaction_type
								, user_id
							)
							VALUES (
								@am
								, @transtype
								, @userID
							);
							SELECT SCOPE_IDENTITY()`)
	tx, errDB := r.sqlServerDB.Begin()
	defer func() {
		if errDB != nil {
			_ = r.sqlServerDB.Rollback(tx)
		} else {
			_ = r.sqlServerDB.Commit(tx)
		}
	}()

	errDB = tx.QueryRow(statement.String(), sql.Named("am", tran.Amount),
		sql.Named("transtype", tran.TransactionType), sql.Named("userID", tran.UserID)).Scan(&id)

	if errDB != nil {
		return 0, errDB
	}
	return id, nil
}

func (r *sqlServerRepository) UpdateTransaction(ctx context.Context, tran model.Transaction) (err error) {
	var (
		statement strings.Builder
	)
	statement.WriteString(`UPDATE transaction 
							SET amount = @pAmount
							, transaction_type = @pTranType
							, user_id = @pUserID
							WHERE id = @pID`)
	tx, errDB := r.sqlServerDB.Begin()
	defer func() {
		if errDB != nil {
			_ = r.sqlServerDB.Rollback(tx)
		} else {
			_ = r.sqlServerDB.Commit(tx)
		}
	}()

	_, errDB = tx.Exec(statement.String(), sql.Named("pAmount", tran.Amount), sql.Named("pTranType", tran.TransactionType),
		sql.Named("pUserID", tran.UserID), sql.Named("pID", tran.ID))

	if errDB != nil {
		return errDB
	}
	return nil
}
