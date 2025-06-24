package auth

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header["Authorization"]

		// only accept 1 token
		if len(auth) != 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		authContent := strings.Split(auth[0], " ")

		// The token must be in the format: Authorization: Bearer <token>
		// This uses short-circuit evaluation, so it's safe to access index 0 and 1
		// because we already checked the slice length above
		if len(authContent) != 2 || authContent[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		token := authContent[1]

		// the token must be valid: issued by this server
		claims, err := ValidateToken(token)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// ignore the error as at this point this is a valid token
		sub, _ := strconv.Atoi(claims.Subject)

		ctx.Set("userID", sub)

		ctx.Next()
	}
}
