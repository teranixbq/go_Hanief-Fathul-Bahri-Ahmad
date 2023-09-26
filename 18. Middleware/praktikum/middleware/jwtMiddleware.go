package middleware

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	echojwt"github.com/labstack/echo-jwt/v4"
	"praktikum/rest/constants"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(constants.JWT_SECRET),
		SigningMethod: "HS256",
	})
}

func CreateToken (user_id int, name string)(string,error){
	claims := jwt.MapClaims{}
	claims["user_id"] = user_id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_SECRET))
}