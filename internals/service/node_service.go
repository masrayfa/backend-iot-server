package service

import (
	"context"

	"github.com/masrayfa/internals/models/domain"
	"github.com/masrayfa/internals/models/web"
)

type NodeService interface {
	FindAll(ctx context.Context, limit int64) ([]domain.NodeWithFeed, error)
	FindById(ctx context.Context, id int64, limit int64) (domain.NodeWithFeed, error)
	FindHardwareNode(ctx context.Context, id int64) (web.NodeByHardwareResponse, error)
	Create(ctx context.Context, req web.NodeCreateRequest) (web.NodeCreateResponse, error)
	Update(ctx context.Context, req web.NodeUpdateRequest, id int64) error
	Delete(ctx context.Context, id int64) error
}