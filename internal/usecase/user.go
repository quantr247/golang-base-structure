package usecase

import (
	"context"
	"golang-base-structure/config"
	"golang-base-structure/internal/dto"
	"golang-base-structure/internal/repository"
)

type (
	UserUseCase interface {
		RegisterUser(ctx context.Context, req *dto.RegisterUserRequestDTO) (*dto.RegisterUserResponseDTO, error)
		GetUserByID(ctx context.Context, id string) (*dto.UserResponseDTO, error)
	}

	userUseCase struct {
		cfg                *config.Config
		postgresRepository repository.PostgresRepository
	}
)

func NewUserUseCase(
	cfg *config.Config,
	postgresRepository repository.PostgresRepository,
) UserUseCase {
	return &userUseCase{
		cfg:                cfg,
		postgresRepository: postgresRepository,
	}
}

func (u *userUseCase) RegisterUser(ctx context.Context, req *dto.RegisterUserRequestDTO) (res *dto.RegisterUserResponseDTO, err error) {
	res = &dto.RegisterUserResponseDTO{}
	return res, nil
}

func (u *userUseCase) GetUserByID(ctx context.Context, id string) (res *dto.UserResponseDTO, err error) {
	res = &dto.UserResponseDTO{
		ID: "1",
		UserName: "ahihi",
		PhoneNumber: "0909000999",
	}
	return res, nil
}