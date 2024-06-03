package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
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

func NewNodeService(repository repository.NodeRepository, hardwareRepository repository.HardwareRepository,  channelRepository repository.ChannelRepository, userRepository repository.UserRepository, db *pgxpool.Pool, validator *validator.Validate) NodeService {
	return &NodeServiceImpl{
		db: db,
		repository: repository,
		hardwareRepository: hardwareRepository,
		channelRepository: channelRepository,
		userRepository: userRepository,
		validator: validator,
	}
}

// need user authentication middleware
func (service *NodeServiceImpl) FindAll(ctx context.Context, limit int64) ([]domain.NodeWithFeed, error) {
	err := service.validator.Struct(ctx)
	if err != nil {
		return []domain.NodeWithFeed{}, errors.New("error when validate context")
	}

	dbpool := service.db

	currentUser, ok := ctx.Value("currentUser").(domain.UserRead)
	if !ok {
		return []domain.NodeWithFeed{}, errors.New("user not found")
	}

	log.Println("currentUser: ", currentUser)

	nodes, err := service.repository.FindAll(ctx, dbpool, &currentUser)
	if err != nil {
		return []domain.NodeWithFeed{}, err
	}

	// get all channel
	nodeChannels, err := service.channelRepository.GetNodeChannelMultiple(ctx, dbpool, nodes, limit)
	if err != nil {
		return []domain.NodeWithFeed{}, err
	}

	log.Println("node channels dari node all service: ", nodeChannels)

	return nodeChannels, nil
}

// need user authentication middleware
func (service *NodeServiceImpl) FindById(ctx context.Context, id int64, limit int64) (domain.NodeWithFeed, error) {
	err := service.validator.Struct(ctx)
	if err != nil {
		return domain.NodeWithFeed{}, errors.New("error when validate context")
	}

	dbpool := service.db

	nodeResponseChannel := make(chan web.NodeResponse)

	go func() {
		node, err := service.repository.FindById(ctx, dbpool, id)
		nodeResponseChannel <- web.NodeResponse{Node: node, Err: err}
	}()

	
	currentUser, ok := ctx.Value("currentUser").(domain.UserRead)
	if !ok {
		return domain.NodeWithFeed{}, errors.New("user not found")
	}
		
	nodeResponse := <- nodeResponseChannel
	node := nodeResponse.Node
	err = nodeResponse.Err
	if err != nil {
		return domain.NodeWithFeed{}, err
	}

	// fmt.Println("node id user: ", node.IdUser)
	// fmt.Println("current id user: ", currentUser.IdUser)

	if node.IdUser != currentUser.IdUser && !currentUser.IsAdmin {
		return domain.NodeWithFeed{}, errors.New("user is not authorized")
	}

	feed, err := service.channelRepository.GetNodeChannel(ctx, dbpool, id, limit)
	if err != nil {
		return domain.NodeWithFeed{}, errors.New("error when get node channel")
	}

	nodeWithFeed := domain.NodeWithFeed{
		Node: node,
		Feed: feed,
	}
	// hardware, err := service.hardwareRepository.FindById(ctx, dbpool, node.IdHardwareNode)
	// helper.PanicIfError(err)

	// var fieldSensor []domain.Hardware
	// for _, idHardwareSensor := range node.IdHardwareSensor {
	// 	hardware, err := service.hardwareRepository.FindById(ctx, dbpool, idHardwareSensor)
	// 	helper.PanicIfError(err)
	// 	fieldSensor = append(fieldSensor, hardware)
	// }

	return nodeWithFeed, nil
}

// need user authentication middleware
func (service *NodeServiceImpl) Create(ctx context.Context, req web.NodeCreateRequest) (nodeCreateRes web.NodeCreateResponse, err error) {
	err = service.validator.Struct(req)
	if err != nil {
		return nodeCreateRes, errors.New("error when validate request")
	}

	log.Println("req: ", req)
	// hardware validation for node 
	hardwareType, err := service.hardwareRepository.FindHardwareTypeById(ctx, service.db, req.IdHardwareNode)
	if err != nil {
		log.Println("err hardwareType: ", err)
		return nodeCreateRes, errors.New("hardware type is not found") 
	}
	hardwareType = strings.ToLower(hardwareType)
	if hardwareType != "microcontroller unit" && hardwareType != "single-board computer" {
		return nodeCreateRes, errors.New("hardware type is not valid")
	}

	// validate sensor hardware id length with sensor field
	if len(req.IdHardwareSensor) != len(req.FieldSensor) {
		return nodeCreateRes, errors.New("sensor hardware id length is not valid")
	}

	currentUser, ok := ctx.Value("currentUser").(domain.UserRead)
	if !ok {
		return nodeCreateRes, errors.New("user not found")
	}
	log.Println("currentUser: ", currentUser)

	sensorHardwareIdLength := len(req.IdHardwareSensor)
	// validate sensor hardware id
	for i := 0; i < sensorHardwareIdLength; i++ {
		log.Println("req.IdHardwareSensor[i]: ", req.IdHardwareSensor[i])
		// validate sensor hardware id
		hardwareTypeSensor, err := service.hardwareRepository.FindHardwareTypeById(ctx, service.db, req.IdHardwareSensor[i])
		if err != nil {
			return nodeCreateRes, err
		}

		hardwareTypeSensor = strings.ToLower(hardwareTypeSensor)
		log.Println("hardwareTypeSensor: ", hardwareTypeSensor)
		
		if hardwareTypeSensor != "sensor" {
			return nodeCreateRes, errors.New("sensor hardware type is not valid")
		}
	}

	// create node object
	node := domain.Node{
		Name: req.Name,
		Location: req.Location,
		FieldSensor: req.FieldSensor,
		IdUser: currentUser.IdUser,
		IdHardwareNode: req.IdHardwareNode,
		IdHardwareSensor: req.IdHardwareSensor,
		IsPublic: req.IsPublic,
	}

	// create node in database
	node, err = service.repository.Create(ctx, service.db, node, currentUser.IdUser)
	if err != nil {
		return nodeCreateRes, err
	}

	nodeCreateRes = web.NodeCreateResponse{
		Name: node.Name,
		Location: node.Location,
		FieldSensor: node.FieldSensor,
		IdHardwareNode: node.IdHardwareNode,
		IdHardwareSensor: node.IdHardwareSensor,
		IsPublic: node.IsPublic,
	}

	return nodeCreateRes, nil	
}

// need user authentication middleware
func (service *NodeServiceImpl) Update(ctx context.Context, req web.NodeUpdateRequest, id int64) error {
	// validate request
	err := service.validator.Struct(ctx)
	if err != nil {
		return errors.New("error when validate request")
	}

	// establish db connection
	dbpool := service.db

	// get node
	node, err := service.repository.FindById(ctx, dbpool, id)
	if err != nil {
		return errors.New("node not found")
	}

	// setup node payload
	nodePayload := web.NodeUpdateRequest(req)
	nodePayload.ChangeSettedField(&node)

	currentUser, ok := ctx.Value("currentUser").(domain.UserRead)
	if !ok {
		return errors.New("user not found")
	}

	if node.IdUser != currentUser.IdUser && !currentUser.IsAdmin {
		return errors.New("user is not authorized")
	}
	fmt.Println("current user: ", currentUser)

	// update node
	_, err = service.repository.Update(ctx, dbpool, &node, &nodePayload)
	if err != nil {
		return errors.New("error when update node")
	}

	return nil
}

// need user authentication middleware
func (service *NodeServiceImpl) Delete(ctx context.Context,id int64) error {
	// validate request
	err := service.validator.Struct(ctx)
	if err != nil {
		return errors.New("error when validate request")
	}

	// establish db connection
	dbpool := service.db

	// get node
	node, err := service.repository.FindById(ctx, dbpool, id)
	if err != nil {
		return errors.New("node not found")
	}

	// get user
	currentUser, ok := ctx.Value("currentUser").(domain.UserRead)
	if !ok {
		return errors.New("user not found")
	}

	// check if user is authorized
	if node.IdUser != currentUser.IdUser && !currentUser.IsAdmin {
		return errors.New("user is not authorized")
	}

	// delete node
	err = service.repository.Delete(ctx, dbpool, id)
	if err != nil {
		return errors.New("error when delete node")
	}

	return nil
}

func (service *NodeServiceImpl) FindHardwareNode(ctx context.Context, id int64) (web.NodeByHardwareResponse, error) {
	err := service.validator.Struct(ctx)
	if err != nil {
		return web.NodeByHardwareResponse{}, err
	}

	dbpool := service.db

	currentUser, ok := ctx.Value("currentUser").(domain.UserRead)
	if !ok {
		return web.NodeByHardwareResponse{}, errors.New("user not found")
	}
	
	hardware, err := service.hardwareRepository.FindById(ctx, dbpool, id)
	if err != nil {
		return web.NodeByHardwareResponse{}, errors.New("hardware not found")
	}
	
	// get hardware node
	nodes, err := service.repository.FindHardwareNode(ctx, dbpool, currentUser.IdUser, id)
	if err != nil {
		return web.NodeByHardwareResponse{}, errors.New("node not found")
	}

	nodeHardware := make([]web.NodeByHardware,0)
	for _, node := range nodes {
		nodeHardware = append(nodeHardware, web.NodeByHardware{
			IdNode: node.IdNode,
			Name: node.Name,
			Location: node.Location,
		})
	}

	nodeHardwareRes := web.NodeByHardwareResponse{
		Hardware: hardware,
		Node: nodeHardware,
	}

	return nodeHardwareRes, nil
}