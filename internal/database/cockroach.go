package database

import (
	"context"
	"log"
	"time"

	appConf "github.com/audryus/crispy-octo-system/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitCockroachConnection(conf appConf.Cockroach) {
	var err error
	Pool, err = pgxpool.NewWithConfig(context.Background(), databaseConfig(conf))
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}
	log.Print("Cockroach connected")
}

func databaseConfig(conf appConf.Cockroach) *pgxpool.Config {
	const defaultMaxConns = int32(2)
	const defaultMinConns = int32(1)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	dbConfig, err := pgxpool.ParseConfig(conf.Url)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	return dbConfig
}
