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
	parseChannel := make(chan error)
	go func() {
		err := service.validator.Struct(req)
		parseChannel <- err
	}()

	type currentUserResult struct {
		res domain.UserRead
		err error
	}
	currentUserChannel := make(chan currentUserResult)
	go func() {
		currentUser, ok := ctx.Value("currentUser").(domain.UserRead)
		if !ok {
			currentUserChannel <- currentUserResult{
				err: nil,
			}
		}
		currentUserChannel <- currentUserResult{
			res: currentUser,
			err: nil,
		}
	}()

	// parsing 
	err := <- parseChannel
	if err != nil {
		return web.ChannelReadResponse{}, err
	}

	// validate node owner async
	nodeOwnerErrorChannel := make(chan error)
	go func() {
		node, err := service.nodeRepository.FindById(ctx, service.db, req.IdNode)
		if err != nil {
			nodeOwnerErrorChannel <- err
			return
		}
		currentUserRes := <- currentUserChannel
		err = currentUserRes.err
		if err != nil {
			nodeOwnerErrorChannel <- err
			return
		}
		currentUser := currentUserRes.res

		if currentUser.IdUser != node.IdUser {
			nodeOwnerErrorChannel <- errors.New("node owner not match")
			return
		}

		nodeOwnerErrorChannel <- nil
	}()

	err = <- nodeOwnerErrorChannel
	if err != nil {
		return web.ChannelReadResponse{}, err
	}

	// create channel
	channel := domain.Channel{
		Value: req.Value,
		IdNode: req.IdNode,
	}

	channel, err = service.repository.Create(ctx, service.db, channel)
	if err != nil {
		return web.ChannelReadResponse{}, err
	}

	// convert to web response
	channelResponse := web.ChannelReadResponse{
		Value: channel.Value,
		IdNode: channel.IdNode,
	}

	// return web response
	return channelResponse, nil
}
