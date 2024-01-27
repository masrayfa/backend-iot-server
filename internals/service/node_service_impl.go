package service

import (
	"context"
	"errors"
	"strings"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/domain"
	"github.com/masrayfa/internals/models/web"
	"github.com/masrayfa/internals/repository"
)

type NodeServiceImpl struct {
	db *pgxpool.Pool
	repository repository.NodeRepository
	hardwareRepository repository.HardwareRepository
	channelRepository repository.ChannelRepository
	userRepository repository.UserRepository
	validator *validator.Validate 
}

func NewNodeService(db *pgxpool.Pool, repository repository.NodeRepository, hardwareRepository repository.HardwareRepository, channelRepository repository.ChannelRepository, validator *validator.Validate) NodeService {
	return &NodeServiceImpl{
		db: db,
		repository: repository,
		hardwareRepository: hardwareRepository,
		channelRepository: channelRepository,
		validator: validator,
	}
}

func (service *NodeServiceImpl) FindAll(ctx context.Context, limit int64, idUser int64) ([]domain.NodeWithFeed, error) {
	err := service.validator.Struct(ctx)
	helper.PanicIfError(err)

	dbpool := service.db

	currentUser, err := service.userRepository.FindById(ctx, dbpool, idUser)
	helper.PanicIfError(err)

	nodes, err := service.repository.FindAll(ctx, dbpool, &currentUser)
	helper.PanicIfError(err)

	// get all channel
	nodeChannels, err := service.channelRepository.GetNodeChannelMultiple(ctx, dbpool, nodes, limit)
	helper.PanicIfError(err)

	return nodeChannels, nil
}

func (service *NodeServiceImpl) FindById(ctx context.Context, id int64, limit int64) (domain.NodeWithFeed, error) {
	err := service.validator.Struct(ctx)
	helper.PanicIfError(err)

	dbpool := service.db

	nodeResponseChannel := make(chan web.NodeResponse)

	go func() {
		node, err := service.repository.FindById(ctx, dbpool, id)
		nodeResponseChannel <- web.NodeResponse{Node: node, Err: err}
	}()

	currentUser, err := service.userRepository.FindById(ctx, dbpool, id)
	helper.PanicIfError(err)

	nodeResponse := <- nodeResponseChannel
	node := nodeResponse.Node
	err = nodeResponse.Err
	helper.PanicIfError(err)

	if node.IdUser != currentUser.IdUser && !currentUser.IsAdmin {
		return domain.NodeWithFeed{}, errors.New("user is not authorized")
	}

	feed, err := service.channelRepository.GetNodeChannel(ctx, dbpool, id, limit)
	helper.PanicIfError(err)

	nodeWithFeed := domain.NodeWithFeed{
		Node: node,
		Feed: feed,
	}

	return nodeWithFeed, nil
}

func (service *NodeServiceImpl) Create(ctx context.Context, req web.NodeCreateRequest, idUser int64) (nodeCreateRes web.NodeCreateResponse, err error) {
	nodeRequest := web.NodeCreateRequest{}
	parseChannel := make(chan error)

	// parse req body in goroutine
	go func() {
		nodeRequest = req
		err = service.validator.Struct(&nodeRequest)
		parseChannel <- err
	}() 

	err = <- parseChannel
	if err != nil {
		return nodeCreateRes, err
	}

	// hardware validation for node asynchronusly
	validateNodeHardwareChannel := make(chan error)
	go func() {
		hardwareType, err := service.hardwareRepository.FindHardwareTypeById(ctx, service.db, nodeRequest.IdHardwareNode)
		if err != nil {
			validateNodeHardwareChannel <- err
		}

		hardwareType = strings.ToLower(hardwareType)
		if hardwareType != "microcontroller-unit" && hardwareType != "single-board-computer" {
			validateNodeHardwareChannel <- errors.New("hardware type is not valid")
			return
		}

		validateNodeHardwareChannel <- nil
	}()

	// validate sensor hardware id length with sensor field
	if len(nodeRequest.IdHardwareSensor) != len(nodeRequest.FieldSensor) {
		return nodeCreateRes, errors.New("sensor hardware id length is not valid")
	}

	// hardware validation for sensor asynchronusly
	sensorHardwareIdLength := len(nodeRequest.IdHardwareSensor)
	validateSensorHardwareChannel := make(chan error, sensorHardwareIdLength)
	 for _, id := range nodeRequest.IdHardwareSensor {
		go func(id int64) {
			hardwareType, err := service.hardwareRepository.FindHardwareTypeById(ctx, service.db, id)
			if err != nil {
				validateSensorHardwareChannel <- err
			}

			hardwareType = strings.ToLower(hardwareType)
			if hardwareType != "sensor" {
				validateSensorHardwareChannel <- errors.New("hardware type is not valid")
				return
			}

			validateSensorHardwareChannel <- nil
		}(id)
	}

	currentUser, err := service.userRepository.FindById(ctx, service.db, idUser)
	helper.PanicIfError(err)

	err = <- validateNodeHardwareChannel
	if err != nil {
		return nodeCreateRes, err
	}

	for i := 0; i < sensorHardwareIdLength; i++ {
		err = <- validateSensorHardwareChannel
		if err != nil {
			return nodeCreateRes, err
		}
	}

	// create node object
	node := domain.Node{
		Name: nodeRequest.Name,
		Location: nodeRequest.Location,
		IdUser: currentUser.IdUser,
		IdHardwareNode: nodeRequest.IdHardwareNode,
	}

	// create node in database
	node, err = service.repository.Create(ctx, service.db, node, &currentUser)
	helper.PanicIfError(err)

	nodeCreateRes = web.NodeCreateResponse{
		Name: node.Name,
		Location: node.Location,
		IdHardwareNode: node.IdHardwareNode,
		FieldSensor: node.FieldSensor,
		IdHardwareSensor: node.IdHardwareSensor,
		IsPublic: node.IsPublic,
	}

	return nodeCreateRes, nil	
}

func (service *NodeServiceImpl) Update(ctx context.Context, req web.NodeUpdateRequest, id int64) error {
	// validate request
	err := service.validator.Struct(ctx)
	helper.PanicIfError(err)

	// establish db connection
	dbpool := service.db

	// get node
	node, err := service.repository.FindById(ctx, dbpool, id)
	helper.PanicIfError(err)

	// setup node payload
	nodePayload := web.NodeUpdateRequest(req)
	nodePayload.ChangeSettedField(&node)

	// update node
	_, err = service.repository.Update(ctx, dbpool, node, &nodePayload)
	helper.PanicIfError(err)

	return nil
}

func (service *NodeServiceImpl) Delete(ctx context.Context,id int64) error {
	panic("implement me")
}