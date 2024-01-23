package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/domain"
)

type HardwareRepositoryImpl struct {
}

func NewHardwareRepositoryImpl() HardwareRepository {
	return &HardwareRepositoryImpl{}
}

func (r *HardwareRepositoryImpl) FindAllItem(ctx context.Context, pool *pgxpool.Pool, statement string) ([]domain.Hardware, error) {
	var emptyHardware []domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) FindAllHardware(ctx context.Context, pool *pgxpool.Pool ) ([]domain.Hardware, error) {
	var emptyHardware []domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) FindAllNode(ctx context.Context, pool *pgxpool.Pool ) ([]domain.Hardware, error) {
	tx, err := pool.Begin(ctx)
	defer helper.CommitOrRollback(ctx, tx)

	script := "SELECT * FROM hardware WHERE lower(type) = 'single-board computer' or lower(type) = 'microcontroller'"

	rows, err := tx.Query(ctx, script)
	helper.PanicIfError(err)

	var hardwares []domain.Hardware
	for rows.Next() {
		var hardware domain.Hardware
		err = rows.Scan(&hardware.IdHardware, &hardware.Name, &hardware.Type, &hardware.Description)
		helper.PanicIfError(err)

		hardwares = append(hardwares, hardware)
	}

	return hardwares, nil
}

func (r *HardwareRepositoryImpl) FindAllSensor(ctx context.Context, pool *pgxpool.Pool ) ([]domain.Hardware, error) {
	tx, err := pool.Begin(ctx)
	defer helper.CommitOrRollback(ctx, tx)

	script := "SELECT * FROM hardware WHERE lower(type) = 'sensor'"
	rows, err := tx.Query(ctx, script)
	helper.PanicIfError(err)

	var hardwares []domain.Hardware
	for rows.Next() {
		var hardware domain.Hardware
		err = rows.Scan(&hardware.IdHardware, &hardware.Name, &hardware.Type, &hardware.Description)
		helper.PanicIfError(err)
		
		hardwares = append(hardwares, hardware)
	}

	return hardwares, nil
}

func (r *HardwareRepositoryImpl) FindById(ctx context.Context, pool *pgxpool.Pool , id int64) (domain.Hardware, error) {
	var emptyHardware domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) Create(ctx context.Context, pool *pgxpool.Pool ,hardware domain.Hardware) (domain.Hardware, error) {
	var emptyHardware domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) Update(ctx context.Context, pool *pgxpool.Pool , hardware domain.Hardware) (domain.Hardware, error) {
	var emptyHardware domain.Hardware
	return emptyHardware, nil
}

func (r *HardwareRepositoryImpl) Delete(ctx context.Context, pool *pgxpool.Pool , id int64) error {
	return nil
}