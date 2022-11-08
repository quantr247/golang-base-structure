package dto

// swagger:parameters postTransaction
type TransactionRequestBody struct {
	// in: body
	Body *TransactionRequestDTO
}

type TransactionRequestDTO struct {
	Amount      int64
	Description string
}

// swagger:model TransactionResponseDTO
type TransactionResponseDTO struct {
	ResultCode int64
	Message    string
}
