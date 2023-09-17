package controllers

import (
	"encoding/json"
	"go-htmx/internal/database"
	"go-htmx/internal/database/models"
	"go-htmx/internal/helpers"
	"go-htmx/internal/utils/token"
	"net/http"
	"os"

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

	account := models.Account{}

	jsonData, _ := json.Marshal(data)

	json.Unmarshal(jsonData, &account)

	err := verifyPassword(request.Password, account.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		helpers.BadRequestWithMsg(c, "Incorrent Password")
		return
	}

	if account.IsAdmin {
		token, err := token.GenerateToken(account.Id, true)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		c.SetCookie("admin_auth", token, 86400, "/", "localhost", false, true)
	}

	token, err := token.GenerateToken(account.Id, false)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.SetCookie("auth", token, 86400, "/", "localhost", false, true)
	helpers.SuccessWithData(c, gin.H{
		"message": "success",
	})
}

func signOut(c *gin.Context) {
	c.SetCookie("auth", "", -3600, "/", "localhost", false, true)
	c.SetCookie("admin_auth", "", -3600, "/", "localhost", false, true)

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

	account := models.Account{
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	is_admin := account.Email == os.Getenv("ADMIN_EMAIL")
	if is_admin {
		account.IsAdmin = true
	}

	tx := database.Create(&account)

	if tx.Error != nil {
		helpers.BadRequestWithMsg(c, "Email is already in use.")
		return
	}

	if is_admin {
		token, err := token.GenerateToken(account.Id, true)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		c.SetCookie("admin_auth", token, 86400, "/", "localhost", false, true)
	}

	token, err := token.GenerateToken(account.Id, false)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.SetCookie("auth", token, 86400, "/", "localhost", false, true)
	helpers.SuccessWithData(c, gin.H{
		"message": "success",
	})
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
