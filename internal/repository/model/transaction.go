package model

type Transaction struct {
	ID              string
	Amount          int64
	TransactionType string
	UserID          string
}
