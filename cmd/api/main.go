package main

import (
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/umdalecs/todo-list-api/config"
	"github.com/umdalecs/todo-list-api/internal/api"
	"github.com/umdalecs/todo-list-api/internal/db"
)

func main() {
	cfg, err := pgxpool.ParseConfig("")
	if err != nil {
		log.Fatal(err)
	}

	cfg.ConnConfig.Host = config.Envs.DbAddr
	cfg.ConnConfig.Port = config.Envs.DbPort
	cfg.ConnConfig.User = config.Envs.DbUser
	cfg.ConnConfig.Password = config.Envs.DbPass
	cfg.ConnConfig.Database = config.Envs.DbName
	cfg.ConnConfig.TLSConfig = nil
	cfg.MaxConns = 10
	cfg.MaxConnLifetime = time.Hour

	db := db.InitPostgresDB(cfg)

	s := api.NewAPIServer(":8080", db)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
