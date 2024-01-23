package service

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/domain"
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
	err := service.validate.Struct(ctx)
	helper.PanicIfError(err)

	dbpool := service.db

	nodes, err := service.repository.FindAllNode(ctx, dbpool)
	helper.PanicIfError(err)

	sensors, err := service.repository.FindAllSensor(ctx, dbpool)
	helper.PanicIfError(err)

	webNodes := make([]web.HardwareReadResponse, len(nodes))
	for i, node := range nodes {
		webNodes[i] = web.HardwareReadResponse{
			IdHardware: int64(node.IdHardware),
			Name: node.Name,
			Type: node.Type,
			Description: node.Description,
		}
	}

	webSensors := make([]web.HardwareReadResponse, len(sensors))
	for i, sensor := range sensors {
		webSensors[i] = web.HardwareReadResponse{
			IdHardware: int64(sensor.IdHardware),
			Name: sensor.Name,
			Type: sensor.Type,
			Description: sensor.Description,
		}
	}

	webHardwares := append(webNodes, webSensors...)

	return webHardwares, nil

}

func (service *HardwareServiceImpl) FindById(ctx context.Context,id int64) (web.HardwareReadResponse, error) {
	panic("implement me")
}

func (service *HardwareServiceImpl) Create(ctx context.Context, req web.HardwareCreateRequest) (web.HardwareReadResponse, error) {
	// validate request
	err := service.validate.Struct(req)
	helper.PanicIfError(err)

	// establish db connection
	dbpool := service.db

	// create domain object
	hardware := domain.Hardware {
		Name: req.Name,
		Type: req.Type,
		Description: req.Description,
	}

	// insert to db
	hardware, err = service.repository.Create(ctx, dbpool, hardware)
	helper.PanicIfError(err)

	// convert to web response
	hardwareResponse := web.HardwareReadResponse {
		IdHardware: int64(hardware.IdHardware),
		Name: hardware.Name,
		Type: hardware.Type,
		Description: hardware.Description,
	}

	// return web response
	return hardwareResponse, nil
}

func (service *HardwareServiceImpl) Update(ctx context.Context, req web.HardwareUpdateRequest, id int64) (web.HardwareReadResponse, error) {
	panic("implement me")
}

func (service *HardwareServiceImpl) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}