package dto

type TransactionRequestDTO struct {
	Amount      int64
	Description string
}

type TransactionResponseDTO struct {
	ResultCode int64
	Message    string
}
