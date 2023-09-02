package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pages(router *gin.Engine) *gin.RouterGroup {
	pages := router.Group("/")

	{
		pages.GET("/", homePage)
	}

	return pages
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home",
	})
}
