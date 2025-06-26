package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umdalecs/todo-list-api/internal/auth"
)

type TodoHandler struct {
	repo *TodoRepository
}

func NewTodosHandler(repo *TodoRepository) *TodoHandler {
	return &TodoHandler{
		repo: repo,
	}
}

func (h *TodoHandler) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/todos")

	g.Use(auth.AuthMiddleware())

	g.POST("/", h.handleCreate)
	g.PUT("/:id", h.handleUpdate)
	g.DELETE("/:id", h.handleDelete)
	g.GET("/", h.handleGetAll)
}

func (h *TodoHandler) handleCreate(ctx *gin.Context) {

}

func (h *TodoHandler) handleUpdate(ctx *gin.Context) {

}

func (h *TodoHandler) handleDelete(ctx *gin.Context) {

}

func (h *TodoHandler) handleGetAll(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	ctx.JSON(http.StatusOK, gin.H{"user": userID})
}
