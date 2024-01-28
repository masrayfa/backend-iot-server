package service

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/models/web"
	"github.com/masrayfa/internals/repository"
)

type ChannelServiceImpl struct {
	db *pgxpool.Pool
	repository repository.ChannelRepository
	nodeRepository repository.NodeRepository
	validator *validator.Validate
}

func NewChannelService(db *pgxpool.Pool, repository repository.ChannelRepository, nodeRepository repository.NodeRepository, validator *validator.Validate) ChannelService {
	return &ChannelServiceImpl{
		db: db,
		repository: repository,
		nodeRepository: nodeRepository,
		validator: validator,
	}
}

func (service *ChannelServiceImpl) GetNodeChannel(ctx context.Context, nodeId int64, limit int64) ([]web.ChannelReadResponse, error) {
	return nil, nil
}

func (service *ChannelServiceImpl) GetNodeChannelMultiple(ctx context.Context, nodes []web.NodeRequest, limit int64) ([]web.NodeWithFeedResponse, error) {
	return nil, nil
}

func (service *ChannelServiceImpl) Create(ctx context.Context, req web.ChannelCreateRequest) (web.ChannelReadResponse, error) {
	var empty web.ChannelReadResponse
	return empty, nil
}