package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {

	secret := os.Getenv(JWT_SECRET_KEY)
	claims := jwt.MapClaims{
		"id": ud.id,
		"email": ud.email,
		"name": ud.name,
		"age": ud.age,
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(
			fmt.Sprintf("Error trying to generate JWT Token, err=%s", err.Error()),
		)
	}

	return tokenString, nil
}

func VerifyTokenMiddleware(c *gin.Context){
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := RemoveBearerPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("Invalid Token. Not parsing")
	})
	if err != nil {
		errRest := rest_err.NewUnauthorizedError("Invalid Token.")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedError("Invalid Token. Claims bad formatting.")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userDomain := userDomain{
		id: claims["id"].(string),
		email: claims["email"].(string),
		name: claims["name"].(string),
		age: int8(claims["age"].(float64)),
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", userDomain))
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer "){
		token = strings.TrimPrefix("Bearer ", token)
	}
	return token
}