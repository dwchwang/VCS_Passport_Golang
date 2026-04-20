package middleware

import (
	"errors"
	"task-management-api/pkg/apperror"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.HTTPCode, appErr)
			return
		}

		c.JSON(500, apperror.ErrInternalServer)
	}
}