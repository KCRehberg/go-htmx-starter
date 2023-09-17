package middlewares

import (
	"net/http"

	"go-htmx/internal/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(token_type string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c, token_type)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
