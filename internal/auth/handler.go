package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	repo *AuthRepository
}

func NewAuthHandler(repo *AuthRepository) *AuthHandler {
	return &AuthHandler{
		repo: repo,
	}
}

func (h *AuthHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/register", h.handleRegister)
	r.POST("/login", h.handleLogin)
}

func (h *AuthHandler) handleRegister(ctx *gin.Context) {
	var payload RegisterDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	validate := validator.New()

	if err := validate.Struct(payload); err != nil {
		errors := &[]string{}
		for _, e := range err.(validator.ValidationErrors) {
			var message string

			if e.Param() == "" {
				message = fmt.Sprintf("field %s has error: %s", e.Field(), e.Tag())
			} else {
				message = fmt.Sprintf("field %s has error: %s=%s", e.Field(), e.Tag(), e.Param())
			}

			*errors = append(*errors, message)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	var user User
	err := h.repo.CreateUser(&payload, &user)
	if err != nil {
		if err == ErrDuplicatedEmail {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error storing user"})
		return
	}

	token, err := GenerateToken(int64(user.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error generating access token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) handleLogin(ctx *gin.Context) {
	var payload LoginDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	var user User
	err := h.repo.GetUserByEmail(payload.Email, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error searching user"})
		return
	}

	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	if !ComparePassword(user.Password, payload.Password) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	token, err := GenerateToken(int64(user.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error generating access token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
