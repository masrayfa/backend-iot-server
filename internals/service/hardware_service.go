package service

import (
	"context"

	"github.com/masrayfa/internals/models/web"
)

type HardwareService interface {
	FindAll(ctx context.Context) ([]web.HardwareReadResponse, error)
	FindById(ctx context.Context, id int64) (web.HardwareReadResponse, error)
	Create(ctx context.Context, req web.HardwareCreateRequest) (web.HardwareReadResponse, error)
	Update(ctx context.Context, req web.HardwareUpdateRequest, id int64) (web.HardwareReadResponse, error)
	Delete(ctx context.Context, id int64) error
}