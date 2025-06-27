package todos

import (
	"net/http"
	"strconv"

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

	// verifies jwt and puts userID in the context
	g.Use(auth.AuthMiddleware())

	g.POST("/", h.handleCreate)
	g.PUT("/:id", h.handleUpdate)
	g.DELETE("/:id", h.handleDelete)
	g.GET("/", h.handleGetAll)
}

func (h *TodoHandler) handleCreate(ctx *gin.Context) {
	userID := getUserIDFromCtx(ctx)

	// this should never happen
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var payload TodoDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var todo Todo
	if err := h.repo.CreateTodo(userID, payload, &todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error saving todo"})
		return
	}

	ctx.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) handleUpdate(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "todo id must be an integer"})
		return
	}

	userID := getUserIDFromCtx(ctx)

	// this should never happen
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var payload TodoDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var todo Todo
	// treats other users todos as non existent
	if err := h.repo.UpdateTodo(todoID, userID, payload, &todo); err != nil {
		if err == ErrTodoNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error saving todo"})
		return
	}

	ctx.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) handleDelete(ctx *gin.Context) {
	userID := getUserIDFromCtx(ctx)

	// this should never happen
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	todoID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "todo id must be an integer"})
		return
	}

	// treats other users todos as non existent
	if err = h.repo.DeleteTodo(userID, todoID); err != nil {
		if err == ErrTodoNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting todo"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *TodoHandler) handleGetAll(ctx *gin.Context) {
	userID := getUserIDFromCtx(ctx)

	// this should never happen
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var (
		page  int
		limit int
		err   error
	)

	pageStr := ctx.Query("page")
	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "page must be an integer"})
		return
	}

	limitStr := ctx.Query("limit")
	if limitStr == "" {
		limit = 5
	} else {
		limit, err = strconv.Atoi(ctx.Query("limit"))
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "limit must be an integer"})
		return
	}

	filter := ctx.Query("filter")

	todos := &[]Todo{}
	var totalPages int
	if err := h.repo.GetAllTodos(userID, page, limit, filter, todos, &totalPages); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving todos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  todos,
		"page":  page,
		"limit": limit,
		"total": totalPages,
	})
}

func getUserIDFromCtx(ctx *gin.Context) int {
	if id, exists := ctx.Get("userID"); exists {
		if intId, ok := id.(int); ok {
			return intId
		}
	}

	return 0
}
