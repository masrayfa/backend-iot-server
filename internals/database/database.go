package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/configs"
)
var databaseUrl string

func init() {
	config := configs.GetConfig()
	databaseConfig := config.Database
	databaseUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", databaseConfig.Username, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Name)
}

func NewDBPool() *pgxpool.Pool {
	fmt.Println("ini database url", databaseUrl)

	config, err := pgxpool.ParseConfig(databaseUrl)
	config.MinConns = 5
	config.MaxConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 10

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	return dbpool
}