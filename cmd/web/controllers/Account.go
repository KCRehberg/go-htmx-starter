package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAccountApi(r *gin.RouterGroup) {
	account := r.Group("/account")
	{
		account.GET("/", getAccount)
	}
}

func getAccount(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/test")
}
