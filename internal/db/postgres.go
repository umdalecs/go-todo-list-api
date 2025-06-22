package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Ctx = context.Background()

func InitPostgresDB(cfg *pgxpool.Config) *pgxpool.Pool {
	conn, err := pgxpool.NewWithConfig(Ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
