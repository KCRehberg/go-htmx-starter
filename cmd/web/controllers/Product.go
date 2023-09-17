package controllers

import (
	"go-htmx/cmd/web/middlewares"
	"go-htmx/internal/database"
	"go-htmx/internal/database/models"
	"go-htmx/internal/helpers"

	"github.com/gin-gonic/gin"
)

func InitProductApi(r *gin.RouterGroup) {
	product := r.Group("/product")
	{
		product.Use(middlewares.JwtAuthMiddleware("admin_auth"))
		product.POST("/create", createProduct)
	}
}

func createProduct(c *gin.Context) {
	product := models.Product{}
	if err := c.BindJSON(&product); err != nil {
		helpers.BadRequestWithMsg(c, err.Error())
		return
	}

	tx := database.Create(&product)
	if tx.Error != nil {
		helpers.BadRequestWithMsg(c, "Error creating product try again.")
		return
	}

	helpers.SuccessWithData(c, gin.H{
		"message": "success",
	})
}
