package service

import (
	"context"

	"github.com/masrayfa/internals/models/domain"
	"github.com/masrayfa/internals/models/web"
)

type HardwareService interface {
	FindAll(ctx context.Context) ([]domain.Hardware, error)
	FindById(ctx context.Context, id int64) (domain.Hardware, error)
	Create(ctx context.Context, req web.HardwareCreateRequest) (domain.Hardware, error)
	Update(ctx context.Context, req web.HardwareUpdateRequest) (domain.Hardware, error)
	Delete(ctx context.Context, id int64) error
}