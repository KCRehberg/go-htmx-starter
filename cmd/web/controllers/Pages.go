package controllers

import (
	"encoding/json"
	"go-htmx/cmd/web/middlewares"
	"go-htmx/internal/database"
	"go-htmx/internal/database/models"
	"go-htmx/internal/utils/token"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PageProps struct {
	Account   models.Account
	Products  []models.Product
	CartItems []models.CartItem
	CartCount uint
}

func Pages(router *gin.Engine) *gin.RouterGroup {
	pages := router.Group("/")

	{
		pages.GET("/", homePage)
		pages.GET("/sign-in", signInPage)
		pages.GET("/sign-up", signUpPage)

		accountPages := pages.Group("/account")
		{
			accountPages.Use(middlewares.JwtAuthMiddleware("auth"))
			accountPages.GET("/", accountDashboard)
		}

		adminPages := pages.Group("/admin")
		{
			adminPages.Use(middlewares.JwtAuthMiddleware("admin_auth"))
			adminPages.GET("/", adminDashboard)
		}
	}

	return pages
}

func homePage(c *gin.Context) {
	pageProps := getPageProps(c)

	c.HTML(http.StatusOK, "index.html", pageProps)
}

func signInPage(c *gin.Context) {
	pageProps := getPageProps(c)

	c.HTML(http.StatusOK, "sign-in.html", pageProps)
}

func signUpPage(c *gin.Context) {
	pageProps := getPageProps(c)

	c.HTML(http.StatusOK, "sign-up.html", pageProps)
}

func accountDashboard(c *gin.Context) {
	pageProps := getPageProps(c)

	c.HTML(http.StatusOK, "account-dashboard.html", pageProps)
}

func adminDashboard(c *gin.Context) {
	pageProps := getPageProps(c)

	c.HTML(http.StatusOK, "admin.html", pageProps)
}

func getPageProps(c *gin.Context) PageProps {
	account := checkUserSession(c)

	pageProps := PageProps{
		Account:   account,
		CartItems: []models.CartItem{},
		CartCount: 0,
	}

	// Get products
	products := []models.Product{}

	query :=
		`SELECT
		*
	FROM products
	ORDER BY id
	LIMIT 10`
	data := database.GetQuery(query, map[string]interface{}{})

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(jsonData, &products)

	pageProps.Products = products

	// Get cart items and count
	id, err := token.ExtractSessionID(c)
	if err != nil {
		log.Println(err)
	}

	if id != uuid.Nil {
		cartItems := []models.CartItem{}

		query :=
			`SELECT
		*
		FROM cart_items
		WHERE session_id = @session_id`
		data := database.GetQuery(query, map[string]interface{}{
			"session_id": id.String(),
		})

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}

		json.Unmarshal(jsonData, &cartItems)

		var counter uint = 0
		for _, cart_item := range cartItems {
			counter += cart_item.Quantity
		}
		pageProps.CartItems = cartItems
		pageProps.CartCount = counter
	}

	return pageProps
}

func checkUserSession(c *gin.Context) models.Account {
	account := models.Account{}

	id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return account
	}

	if id != 0 {
		query :=
			`SELECT
				id,
				email
			FROM accounts
			WHERE id = @id`
		data := database.GetQueryFirst(query, map[string]interface{}{
			"id": id,
		})

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}

		json.Unmarshal(jsonData, &account)
	}

	return account
}
