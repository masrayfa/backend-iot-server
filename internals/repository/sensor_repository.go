package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/masrayfa/internals/models/domain"
	"github.com/masrayfa/internals/models/web"
)

type SensorRepository interface {
	SensorFieldWithoutId() string
	SensorField() string
	SensorPointer(sensor domain.Sensor) []interface{}
	GetAll(ctx context.Context, tx pgx.Tx, currentUser *domain.User) ([]domain.Sensor, error)
	GetById(ctx context.Context, tx pgx.Tx, id int) (domain.Sensor, error)
	GetHardwareSensor(ctx context.Context, tx pgx.Tx, hardwareId int) ([]domain.Sensor, error)
	GetNodeSensor(ctx context.Context, tx pgx.Tx, nodeId int) ([]domain.Sensor, error)
	GetSensorChannel(ctx context.Context, tx pgx.Tx, sensorId int) ([]domain.Channel, error)
	GetUserWhoOwnsSensorById(ctx context.Context, tx pgx.Tx, sensorId int) (domain.User, error)
	Create(ctx context.Context, tx pgx.Tx, sensor domain.Sensor) (domain.Sensor, error)
	Update(ctx context.Context, tx pgx.Tx, sensor domain.Sensor, payload web.SensorUpdateRequest) (domain.Sensor, error)
	Delete(ctx context.Context, tx pgx.Tx, id int64) error
}
