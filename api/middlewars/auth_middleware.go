package middlewars

import (
	"head_app/pkg/token"

	"github.com/gin-gonic/gin"
)

func VwAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.GetHeader("authorization")

		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "authorization token not provided"})
			ctx.Abort()
			return
		}

		claim, err := token.ParseJWT(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		if claim.UserRole == "viewer" {
			ctx.JSON(401, gin.H{"error": "you role is not a viewer"})
			ctx.Abort()
			return
		}

		ctx.Next()

	}
}


func OwnAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.GetHeader("authorization")

		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "authorization token not provided"})
			ctx.Abort()
			return
		}

		claim, err := token.ParseJWT(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		if claim.UserRole == "owner" {
			ctx.JSON(401, gin.H{"error": "you role is not a owner"})
			ctx.Abort()
			return
		}

		ctx.Next()

	}
}