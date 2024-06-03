package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/domain"
)

type HardwareRepositoryImpl struct {
}

func NewHardwareRepository() HardwareRepository {
	return &HardwareRepositoryImpl{}
}

func (r *HardwareRepositoryImpl) FindAllItem(ctx context.Context, pool *pgxpool.Pool, statement string) ([]domain.Hardware, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return nil, errors.New("error when begin transaction")
	}
	defer helper.CommitOrRollback(ctx, tx)

	rows, err := tx.Query(ctx, statement)
	if err != nil {
		return nil, errors.New("error when query rows")
	}
	defer rows.Close()

	var hardwares []domain.Hardware
	for rows.Next() {
		var hardware domain.Hardware
		err = rows.Scan(&hardware.IdHardware, &hardware.Name, &hardware.Type, &hardware.Description)
		if err != nil {
			return nil, errors.New("error when scan row")
		}

		hardwares = append(hardwares, hardware)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("error when scan rows")
	}

	return hardwares, nil
}

func (r *HardwareRepositoryImpl) FindAllHardware(ctx context.Context, pool *pgxpool.Pool ) ([]domain.Hardware, error) {
	sqlStatement := `SELECT * FROM hardware`
	return r.FindAllItem(ctx, pool, sqlStatement)
}

func (r *HardwareRepositoryImpl) FindAllNode(ctx context.Context, pool *pgxpool.Pool ) ([]domain.Hardware, error) {
	sqlStatement := `SELECT * FROM hardware WHERE lower(type) = 'single-board computer' or lower(type) = 'microcontroller unit'`
	return r.FindAllItem(ctx, pool, sqlStatement)
}

func (r *HardwareRepositoryImpl) FindAllSensor(ctx context.Context, pool *pgxpool.Pool ) ([]domain.Hardware, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return nil, errors.New("error when begin transaction")
	}
	defer helper.CommitOrRollback(ctx, tx)

	script := "SELECT * FROM hardware WHERE lower(type) = 'sensor'"
	rows, err := tx.Query(ctx, script)
	if err != nil {
		return nil, errors.New("error when query rows")
	}

	var hardwares []domain.Hardware
	for rows.Next() {
		var hardware domain.Hardware
		err = rows.Scan(&hardware.IdHardware, &hardware.Name, &hardware.Type, &hardware.Description)
		if err != nil {
			return nil, errors.New("error when scan row")
		}
		
		hardwares = append(hardwares, hardware)
	}

	return hardwares, nil
}

func (r *HardwareRepositoryImpl) FindById(ctx context.Context, pool *pgxpool.Pool , id int64) (domain.Hardware, error) {
	script := "SELECT * FROM hardware WHERE id_hardware = $1"

	tx, err := pool.Begin(ctx)
	if err != nil {
		return domain.Hardware{}, errors.New("error when begin transaction")
	}
	defer helper.CommitOrRollback(ctx, tx)

	var hardware domain.Hardware
	err = tx.QueryRow(ctx, script, id).Scan(&hardware.IdHardware, &hardware.Name, &hardware.Type, &hardware.Description)
	if err != nil {
		return domain.Hardware{}, errors.New("error when query row")
	}

	return hardware, nil
}

func (r *HardwareRepositoryImpl) FindHardwareTypeById(ctx context.Context, pool *pgxpool.Pool , id int64) (string, error) {
	script := "SELECT type FROM hardware WHERE id_hardware = $1"
	
	tx, err := pool.Begin(ctx)
	if err != nil {
		return "", errors.New("error when begin transaction")
	}
	defer helper.CommitOrRollback(ctx, tx)

	var hardwareType string
	err = tx.QueryRow(ctx, script, id).Scan(&hardwareType)
	if err != nil {
		return "", err
	}

	return hardwareType, nil
}

func (r *HardwareRepositoryImpl) Create(ctx context.Context, pool *pgxpool.Pool, hardware domain.Hardware) (domain.Hardware, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return domain.Hardware{}, errors.New("error when begin transaction")
	}
	defer helper.CommitOrRollback(ctx, tx)

	script := "INSERT INTO hardware (name, type, description) VALUES ($1, $2, $3) RETURNING id_hardware"
	err = tx.QueryRow(ctx, script, hardware.Name, hardware.Type, hardware.Description).Scan(&hardware.IdHardware)
	if err != nil {
		return domain.Hardware{}, errors.New("error when insert hardware")
	}

	return hardware, nil
}

func (r *HardwareRepositoryImpl) Update(ctx context.Context, pool *pgxpool.Pool, hardware domain.Hardware) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return errors.New("error when begin transaction")
	}
	defer helper.CommitOrRollback(ctx, tx)

	script := "UPDATE hardware SET name = $1, type = $2, description = $3 WHERE id_hardware = $4"
	_, err = tx.Exec(ctx, script, hardware.Name, hardware.Type, hardware.Description, hardware.IdHardware)
	if err != nil {
		return errors.New("error when update hardware")
	}

	log.Println("Hardware with id: ", hardware.IdHardware, " has been updated")

	return nil
}

func (r *HardwareRepositoryImpl) Delete(ctx context.Context, pool *pgxpool.Pool , id int64) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return errors.New("error when begin transaction ")
	}
	defer helper.CommitOrRollback(ctx, tx)

	script := "DELETE FROM hardware WHERE id_hardware = $1"
	res, err := tx.Exec(ctx, script, id)
	if err != nil {
		return errors.New("error when delete hardware")
	}

	if res.RowsAffected() != 1 {
		log.Println("No row affected on delete hardware with id: ", id)
		return err
	}

	log.Println("Hardware with id: ", id, " has been deleted")

	return nil
}

func (r *HardwareRepositoryImpl) FindHardwareSensor(ctx context.Context, pool *pgxpool.Pool , id int64) (domain.Hardware, error) {
	script := "SELECT id_hardware, name, type FROM hardware WHERE id_hardware = $1"
	
	tx, err := pool.Begin(ctx)
	if err != nil {
		return domain.Hardware{}, errors.New("error when begin transaction")
	}
	defer helper.CommitOrRollback(ctx, tx)

	var hardware domain.Hardware
	err = tx.QueryRow(ctx, script, id).Scan(&hardware.IdHardware, &hardware.Name, &hardware.Type)
	if err != nil {
		return domain.Hardware{}, errors.New("error when query row")
	}

	return hardware, nil
}