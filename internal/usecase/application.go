package usecase

import (
	"context"
	"errors"
	"golang-base-structure/config"
	"golang-base-structure/internal/common"
	"golang-base-structure/internal/dto"
	"strings"

	goCache "github.com/patrickmn/go-cache"
	"go.uber.org/zap"
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

	res = &dto.GetApplicationResponseDTO{
		Name: "Ahihi",
		Code: "Do ngok",
	}
	if req != nil && strings.EqualFold(req.ApplicationID, "") {
		zap.S().Warnf("Application id not found")
		return nil, errors.New(common.ReasonApplicationNotFound.Code())
	}
	return res, nil
}
