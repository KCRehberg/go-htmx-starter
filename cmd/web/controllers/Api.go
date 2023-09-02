package controllers

import (
	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) *gin.RouterGroup {
	api := router.Group("/api")

	{
		InitAccountApi(api)
	}

	return api
}
