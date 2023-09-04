package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "Authorization failed",
		"code":  http.StatusUnauthorized,
	})
}

func NotFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error": "Not found",
		"code":  http.StatusNotFound,
	})
}

func BadRequest(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": "Invalid request, check your parameters",
		"code":  http.StatusBadRequest,
	})
}

func BadRequestWithMsg(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": msg,
		"error":   "Invalid request, check your parameters",
		"code":    http.StatusBadRequest,
	})
}

func SuccessWithData[T any](c *gin.Context, json T) {
	c.JSON(http.StatusOK, json)
}

func EmptySuccessRequest(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
