package middlewares

import (
	"gin-blog-app/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Not Authenticated!"})
		ctx.Abort()
		return
	}

	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Not Authenticated!"})
		ctx.Abort()
		return
	}

	claims, err := auth.VerifyToken(parts[1])
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Not Authenticated!"})
		ctx.Abort()
		return
	}

	ctx.Set("username", claims.Username)
	ctx.Next()
}
