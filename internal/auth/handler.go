package auth

import "github.com/gin-gonic/gin"

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/register", h.handleRegister)
	r.POST("/login", h.handleLogin)
}

func (h *AuthHandler) handleRegister(ctx *gin.Context) {

}

func (h *AuthHandler) handleLogin(ctx *gin.Context) {

}
