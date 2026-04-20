package middleware

import (
	"strings"
	"task-management-api/pkg/apperror"
	pkgjwt "task-management-api/pkg/jwt"

	"github.com/gin-gonic/gin"
)

var JWTSecret string

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, apperror.ErrUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := pkgjwt.Parse(tokenStr, JWTSecret)
		if err != nil {
			c.AbortWithStatusJSON(401, apperror.ErrUnauthorized)
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
