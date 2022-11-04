package usecase

import (
	"context"
	"golang-base-structure/config"
	"golang-base-structure/internal/dto"

	goCache "github.com/patrickmn/go-cache"
)

type (
	ApplicationUseCase interface {
		GetApplication(ctx context.Context, req *dto.GetApplicationRequestDTO) (*dto.GetApplicationResponseDTO, error)
	}

	applicationUseCase struct {
		cfg      *config.Config
		memCache *goCache.Cache
		//oracleRepository repository.OracleRepository
	}
)

func NewApplicationUseCase(
	cfg *config.Config,
	memCache *goCache.Cache,
	//oracleRepository repository.OracleRepository,
) ApplicationUseCase {
	return &applicationUseCase{
		cfg:      cfg,
		memCache: memCache,
		//oracleRepository: oracleRepository,
	}
}

func (u *applicationUseCase) GetApplication(ctx context.Context, req *dto.GetApplicationRequestDTO) (
	res *dto.GetApplicationResponseDTO, err error) {
	res = &dto.GetApplicationResponseDTO{}

	return res, nil
}
