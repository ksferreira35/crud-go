package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/ksferreira35/crud-go/src/config/logger"
	rest_err "github.com/ksferreira35/crud-go/src/config/res_err"
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
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(
		fmt.Sprintf("error to trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBetterPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil 
		}

		return nil, rest_err.NewBadRequestError("Invalid token")
	})
	if err != nil {
		return nil, rest_err.NewUnauthorizedRequestError("Invalid token")		
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedRequestError("Invalid token")	
	}

	return &userDomain{
		id: 	   claims["id"].(string),
		email:     claims["email"].(string),
		name:      claims["name"].(string),
		age:       int8(claims["age"].(float64)),
	}, nil
}


func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := RemoveBetterPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(RemoveBetterPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil 
		}


		return nil, rest_err.NewBadRequestError("Invalid token")
	})
	if err != nil {
		errRest := rest_err.NewUnauthorizedRequestError("Invalid token") 
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}
	

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedRequestError("Invalid token") 
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userDomain := userDomain{
		id: 	   claims["id"].(string),
		email:     claims["email"].(string),
		name:      claims["name"].(string),
		age:       int8(claims["age"].(float64)),
	}
	logger.Info(fmt.Sprintf("User athenticated: %#v", userDomain))
}

func RemoveBetterPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer "){
		token = strings.TrimPrefix("Bearer ", token)
	}

	return token
}
