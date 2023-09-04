package controllers

import (
	"encoding/json"
	"go-htmx/internal/database"
	"go-htmx/internal/database/models"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Pages(router *gin.Engine) *gin.RouterGroup {
	pages := router.Group("/")

	{
		pages.GET("/", homePage)
		pages.GET("/sign-in", signInPage)
		pages.GET("/sign-up", signUpPage)
	}

	return pages
}

func homePage(c *gin.Context) {
	cookie, err := c.Cookie("user_email")

	if err != nil {
		log.Println(err)
	}

	account := auth(cookie)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"user_email": account.Email,
	})
}

func signInPage(c *gin.Context) {
	cookie, err := c.Cookie("user_email")

	if err != nil {
		log.Println(err)
	}

	account := auth(cookie)

	c.HTML(http.StatusOK, "sign-in.html", gin.H{
		"user_email": account.Email,
	})
}

func signUpPage(c *gin.Context) {
	cookie, err := c.Cookie("user_email")

	if err != nil {
		log.Println(err)
	}

	account := auth(cookie)

	c.HTML(http.StatusOK, "sign-up.html", gin.H{
		"user_email": account.Email,
	})
}

func auth(cookie string) models.Accounts {
	account := models.Accounts{}

	if cookie != "" {
		user_email, err := url.QueryUnescape(cookie)
		if err != nil {
			log.Println(err)
		}

		query :=
			`SELECT
				id,
				email
			FROM accounts
			WHERE email = @email`
		data := database.GetQueryFirst(query, map[string]interface{}{
			"email": user_email,
		})

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}

		json.Unmarshal(jsonData, &account)
	}

	return account
}
