package todos

import "github.com/gin-gonic/gin"

type TodoHandler struct {
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

func (h *TodoHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/posts", h.handleCreate)
	r.PUT("/posts/:id", h.handleUpdate)
	r.DELETE("/posts/:id", h.handleDelete)
	r.GET("/posts", h.handleGetAll)
}

func (h *TodoHandler) handleCreate(ctx *gin.Context) {

}

func (h *TodoHandler) handleUpdate(ctx *gin.Context) {

}

func (h *TodoHandler) handleDelete(ctx *gin.Context) {

}

func (h *TodoHandler) handleGetAll(ctx *gin.Context) {

}
