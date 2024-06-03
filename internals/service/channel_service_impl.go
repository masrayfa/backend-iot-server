package service

import (
	"context"
	"errors"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/models/domain"
	"github.com/masrayfa/internals/models/web"
	"github.com/masrayfa/internals/repository"
)

type ChannelServiceImpl struct {
	db *pgxpool.Pool
	repository repository.ChannelRepository
	nodeRepository repository.NodeRepository
	validator *validator.Validate
}

func NewChannelService( repository repository.ChannelRepository, nodeRepository repository.NodeRepository, db *pgxpool.Pool, validator *validator.Validate) ChannelService {
	return &ChannelServiceImpl{
		db: db,
		repository: repository,
		nodeRepository: nodeRepository,
		validator: validator,
	}
}

func (service *ChannelServiceImpl) Create(ctx context.Context, req web.ChannelCreateRequest) (web.ChannelReadResponse, error) {
	err := service.validator.Struct(req)
	if err != nil {
		return web.ChannelReadResponse{}, errors.New("error when validate request")
	}

	// log.Println("req dari channel service: ", req)

	currentUser, ok := ctx.Value("currentUser").(domain.UserRead)
	if !ok {
		return web.ChannelReadResponse{}, errors.New("user not found")
	}

	node, err := service.nodeRepository.FindById(ctx, service.db, req.IdNode)
	if err != nil {
		return web.ChannelReadResponse{}, errors.New("node not found")
	}

	if currentUser.IdUser != node.IdUser {
		return web.ChannelReadResponse{}, errors.New("current user is not the owner of the node")
	}

	// create channel
	channel := domain.Channel{
		Value: req.Value,
		IdNode: req.IdNode,
	}

	channel, err = service.repository.Create(ctx, service.db, channel)
	if err != nil {
		return web.ChannelReadResponse{}, errors.New("error when create channel")
	}

	// convert to web response
	channelResponse := web.ChannelReadResponse{
		Value: channel.Value,
		IdNode: channel.IdNode,
	}

	// return web response
	return channelResponse, nil
}
