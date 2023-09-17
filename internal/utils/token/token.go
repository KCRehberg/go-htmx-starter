package token

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenerateToken(user_id uint, is_admin bool) (string, error) {
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if is_admin {
		return token.SignedString([]byte(os.Getenv("ADMIN_API_SECRET")))
	}
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func GenerateSessionToken() (string, uuid.UUID, error) {
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	session_id := uuid.New()
	if err != nil {
		return "", uuid.UUID{}, err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["session_id"] = session_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed_token, err := token.SignedString([]byte(os.Getenv("SESSION_API_SECRET")))
	return signed_token, session_id, err
}

func TokenValid(c *gin.Context, token_type string) error {
	tokenString := ExtractToken(c, token_type)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method:")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		if token_type == "session_token" {
			return []byte(os.Getenv("SESSION_API_SECRET")), nil
		}
		if token_type == "admin_auth" {
			return []byte(os.Getenv("ADMIN_API_SECRET")), nil
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context, token_type string) string {
	token, err := c.Cookie(token_type)
	if err != nil {
		log.Println(err)
	}

	if token != "" {
		return token
	}

	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c, "auth")
	if tokenString != "" {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})
		if err != nil {
			return 0, err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
			if err != nil {
				return 0, err
			}
			return uint(uid), nil
		}
	}
	return 0, nil
}

func ExtractSessionID(c *gin.Context) (uuid.UUID, error) {
	tokenString := ExtractToken(c, "session_token")
	if tokenString != "" {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SESSION_API_SECRET")), nil
		})
		if err != nil {
			return uuid.Nil, err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			uid, err := uuid.Parse(fmt.Sprintf("%v", claims["session_id"]))
			if err != nil {
				return uuid.Nil, err
			}
			return uid, nil
		}
	}
	return uuid.Nil, nil
}
