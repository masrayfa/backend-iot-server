package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/masrayfa/internals/models/domain"
)

type HardwareRepositoryImpl struct {
}

func NewHardwareRepositoryImpl() HardwareRepository {
	return &HardwareRepositoryImpl{}
}

func (r *HardwareRepositoryImpl) FindAllItem(ctx context.Context, tx pgx.Tx, statement string) ([]domain.Hardware, error) {
	var emptyHardware []domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) FindAllHardware(ctx context.Context, tx pgx.Tx) ([]domain.Hardware, error) {
	var emptyHardware []domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) FindAllNode(ctx context.Context, tx pgx.Tx) ([]domain.Hardware, error) {
	var emptyHardware []domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) FindAllSensor(ctx context.Context, tx pgx.Tx) ([]domain.Hardware, error) {
	var emptyHardware []domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) FindById(ctx context.Context, tx pgx.Tx, id int64) (domain.Hardware, error) {
	var emptyHardware domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) Create(ctx context.Context, tx pgx.Tx,hardware domain.Hardware) (domain.Hardware, error) {
	var emptyHardware domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) Update(ctx context.Context, tx pgx.Tx, hardware domain.Hardware) (domain.Hardware, error) {
	var emptyHardware domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) Delete(ctx context.Context, tx pgx.Tx, id int64) error {
	return nil
}