package main

import (
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/umdalecs/todo-list-api/internal"
	"github.com/umdalecs/todo-list-api/internal/db"
)

func main() {
	cfg, err := pgxpool.ParseConfig("")
	if err != nil {
		log.Fatal(err)
	}

	cfg.ConnConfig.Host = internal.Envs.DbAddr
	cfg.ConnConfig.Port = internal.Envs.DbPort
	cfg.ConnConfig.User = internal.Envs.DbUser
	cfg.ConnConfig.Password = internal.Envs.DbPass
	cfg.ConnConfig.Database = internal.Envs.DbName
	cfg.ConnConfig.TLSConfig = nil
	cfg.MaxConns = 10
	cfg.MaxConnLifetime = time.Hour

	db := db.InitPostgresDB(cfg)

	s := NewAPIServer(":8080", db)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
