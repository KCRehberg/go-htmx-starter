package controllers

import (
	"go-htmx/internal/database"
	"go-htmx/internal/database/models"
	"go-htmx/internal/helpers"
	"go-htmx/internal/utils/token"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitCartApi(r *gin.RouterGroup) {
	cart := r.Group("/cart")
	{
		cart.POST("/add/:id", addItemToCart)
	}
}

func addItemToCart(c *gin.Context) {
	product_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
	}
	session_uuid := checkSessionToken(c)
	uuid_string := session_uuid.String()
	log.Println(uuid_string)

	query :=
		`SELECT
			*
		FROM cart_items
		WHERE session_id = @session_id AND product_id = @product_id`
	data := database.GetQueryFirst(query, map[string]interface{}{
		"session_id": uuid_string,
		"product_id": product_id,
	})

	if data != nil {
		query :=
			`UPDATE cart_items
		SET quantity = quantity + 1
		WHERE session_id = @session_id AND product_id = @product_id`
		database.DB.Exec(query, map[string]interface{}{
			"session_id": uuid_string,
			"product_id": product_id,
		})
	} else {
		cart_item := models.CartItem{
			SessionId: uuid_string,
			ProductId: product_id,
			Quantity:  1,
		}

		tx := database.Create(&cart_item)

		if tx.Error != nil {
			helpers.BadRequestWithMsg(c, "Email is already in use.")
			return
		}
	}

	helpers.EmptySuccessRequest(c)
}

func checkSessionToken(c *gin.Context) uuid.UUID {
	id, err := token.ExtractSessionID(c)
	if err != nil {
		log.Println(err)
	}

	if id != uuid.Nil {
		return id
	}

	token, session_id, err := token.GenerateSessionToken()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.SetCookie("session_token", token, 86400, "/", "localhost", false, true)
	return session_id
}
