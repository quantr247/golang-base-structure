package usecase

import (
	"context"
	"golang-base-structure/config"
	"golang-base-structure/internal/dto"
	"golang-base-structure/internal/repository"
)

type (
	TransactionUseCase interface {
		Payment(ctx context.Context, req *dto.TransactionRequestDTO) (*dto.TransactionResponseDTO, error)
	}

	transactionUseCase struct {
		cfg                 *config.Config
		sqlServerRepository repository.SQLServerRepository
	}
)

func NewTransactionUseCase(
	cfg *config.Config,
	sqlServerRepository repository.SQLServerRepository,
) TransactionUseCase {
	return &transactionUseCase{
		cfg:                 cfg,
		sqlServerRepository: sqlServerRepository,
	}
}

func (u *transactionUseCase) Payment(ctx context.Context, req *dto.TransactionRequestDTO) (res *dto.TransactionResponseDTO, err error) {
	res = &dto.TransactionResponseDTO{}
	return res, nil
}
