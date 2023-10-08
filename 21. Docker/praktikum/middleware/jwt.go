package middleware

import (
	"clean/constant"
	
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {

	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(constant.JWT_SECRET),
		SigningMethod: "HS256",
	})
}

func CreateToken (id int)(string,error){
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.JWT_SECRET))
}

func ExtractToken(e echo.Context) (int) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		Id := int(claims["id"].(float64))
		return Id
	}
	return 0
}