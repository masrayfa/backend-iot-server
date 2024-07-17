package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/configs"
)
var databaseUrl string

func init() {
	config := configs.GetConfig()
	databaseConfig := config.Database
	databaseUrl = fmt.Sprintf("postgres://%s:%s@%s/%s", databaseConfig.Username, databaseConfig.Password, databaseConfig.Host, databaseConfig.Name)
}

func NewDBPool() *pgxpool.Pool {
	fmt.Println("ini database url di new database.go", databaseUrl)

	config, err := pgxpool.ParseConfig(databaseUrl)
	config.MinConns = 10
	config.MaxConns = 50
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 10

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil
	}

	// Check connection
	err = dbpool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		dbpool.Close()
		return nil
	}

	log.Println("Database connected")
	return dbpool
}