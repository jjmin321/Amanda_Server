package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type jwtMethod interface {
	CreateRefreshToken()
	CreateAccessToken()
	VerifyRefreshToken()
	VerifyAccessToken()
}

// CreateRefreshToken : Middleware that create RefreshToken
func CreateRefreshToken(Id, Pw string) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["ID"] = Id
	claims["Pw"] = Pw
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix()

	t, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// CreateAccessToken : Middleware that create AccessToken
func CreateAccessToken(Id, Pw string, IsManager bool) (string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["ID"] = Id
	claims["Pw"] = Pw
	claims["IsManager"] = IsManager
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	t, err := accessToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// VerifyRefreshToken : Middleware that verify RefreshToken
func VerifyRefreshToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		ID := claims["ID"].(string)
		Pw := claims["Pw"].(string)
		c.Set("ID", ID)
		c.Set("Pw", Pw)
		return next(c)
	}
}

// VerifyAccessToken : Middleware that verify AccessToken
func VerifyAccessToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		ID := claims["ID"].(string)
		Pw := claims["Pw"].(string)
		IsManager := claims["IsManager"].(bool)
		c.Set("ID", ID)
		c.Set("Pw", Pw)
		c.Set("IsManager", IsManager)
		return next(c)
	}
}
