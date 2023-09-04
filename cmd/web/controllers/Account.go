package controllers

import (
	"encoding/json"
	"go-htmx/internal/database"
	"go-htmx/internal/database/models"
	"go-htmx/internal/helpers"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func InitAccountApi(r *gin.RouterGroup) {
	account := r.Group("/account")
	{
		account.POST("/sign-in", signIn)
		account.GET("/sign-out", signOut)
		account.POST("/sign-up", createAccount)
	}
}

func signIn(c *gin.Context) {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	request := Request{}
	if err := c.BindJSON(&request); err != nil {
		helpers.BadRequestWithMsg(c, err.Error())
		return
	}

	query :=
		`SELECT
		*
	FROM accounts
	WHERE email = @email`
	data := database.GetQueryFirst(query, map[string]interface{}{
		"email": request.Email,
	})

	if data == nil {
		helpers.BadRequestWithMsg(c, "No user found with the provided email")
		return
	}

	account := models.Accounts{}

	jsonData, _ := json.Marshal(data)

	json.Unmarshal(jsonData, &account)

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password))
	if err != nil {
		helpers.BadRequestWithMsg(c, "Incorrent Password")
		return
	}

	c.SetCookie("user_email", account.Email, 3600, "/", "localhost", false, true)
	helpers.SuccessWithData(c, gin.H{
		"message": "success",
	})
}

func signOut(c *gin.Context) {
	c.SetCookie("user_email", "", -3600, "/", "localhost", false, true)

	helpers.EmptySuccessRequest(c)
}

func createAccount(c *gin.Context) {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	request := Request{}
	if err := c.BindJSON(&request); err != nil {
		helpers.BadRequestWithMsg(c, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	account := models.Accounts{
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	database.Create(&account)

	helpers.SuccessWithData(c, gin.H{
		"message": "success",
	})
}
