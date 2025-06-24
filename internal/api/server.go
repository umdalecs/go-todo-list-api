package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/umdalecs/todo-list-api/internal/auth"
	"github.com/umdalecs/todo-list-api/internal/todos"
)

type APIServer struct {
	addr string
	db   *pgxpool.Pool
}

func NewAPIServer(addr string, db *pgxpool.Pool) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	e := gin.Default()

	v1 := e.Group("/api/v1")

	authRepository := auth.NewAuthRepository(s.db)
	authHandler := auth.NewAuthHandler(authRepository)
	authHandler.RegisterRoutes(v1)

	todosRepository := todos.NewTodoRepository(s.db)
	todosHandler := todos.NewTodosHandler(todosRepository)
	todosHandler.RegisterRoutes(v1)

	return e.Run(s.addr)
}
