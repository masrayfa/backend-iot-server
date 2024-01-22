package service

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/models/web"
	"github.com/masrayfa/internals/repository"
)

type HardwareServiceImpl struct {
	repository repository.HardwareRepository
	db *pgxpool.Pool
	validate *validator.Validate
}

func NewHardwareService(repository repository.HardwareRepository, db *pgxpool.Pool, validate *validator.Validate) HardwareService {
	return &HardwareServiceImpl{
		repository: repository,
		db: db,
		validate: validate,
	}
}

func (service *HardwareServiceImpl) FindAll(ctx context.Context) ([]web.HardwareReadResponse, error) {
	panic("implement me")
}

func (service *HardwareServiceImpl) FindById(ctx context.Context,id int64) (web.HardwareReadResponse, error) {
	panic("implement me")
}

func (service *HardwareServiceImpl) Create(ctx context.Context, req web.HardwareCreateRequest) (web.HardwareReadResponse, error) {
	panic("implement me")
}

func (service *HardwareServiceImpl) Update(ctx context.Context, req web.HardwareUpdateRequest, id int64) (web.HardwareReadResponse, error) {
	panic("implement me")
}

func (service *HardwareServiceImpl) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}