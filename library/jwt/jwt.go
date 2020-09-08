package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type jwtMethod interface {
	CreateaccessTokenInLocal()
	CreateRefreshTokenInTeam()
	VerifyRefreshTokenInLocal()
	VerifyRefreshTokenInTeam()
}

// CreateAccessTokenInLocal : Middleware that create AccessToken of Ownproject
func CreateAccessTokenInLocal(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// CreateRefreshTokenInTeam : Middleware that create RefreshToken of Teamproject
func CreateRefreshTokenInTeam(TeamName, Email string) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["TeamName"] = TeamName
	claims["Email"] = Email
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix()

	t, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// CreateAccessTokenInTeam : Middleware that create AccessToken of Teamproject
func CreateAccessTokenInTeam(Idx uint, TeamName, Email string, IsCreator bool) (string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["Idx"] = Idx
	claims["TeamName"] = TeamName
	claims["Email"] = Email
	claims["IsCreator"] = IsCreator
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	t, err := accessToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// VerifyAccessTokenInLocal : Middleware that verify AccessToken of Ownproject
func VerifyAccessTokenInLocal(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		c.Set("id", id)
		return next(c)
	}
}

// VerifyRefreshTokenInTeam : Middleware that verify RefreshToken of Teamproject
func VerifyRefreshTokenInTeam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		TeamName := claims["TeamName"].(string)
		Email := claims["Email"].(string)
		c.Set("TeamName", TeamName)
		c.Set("Email", Email)
		return next(c)
	}
}

// VerifyAccessTokenInTeam : Middleware that verify AccessToken of Teamproject
func VerifyAccessTokenInTeam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		Idx := claims["Idx"].(float64)
		TeamName := claims["TeamName"].(string)
		Email := claims["Email"].(string)
		IsCreator := claims["IsCreator"].(bool)
		c.Set("Idx", Idx)
		c.Set("TeamName", TeamName)
		c.Set("Email", Email)
		c.Set("IsCreator", IsCreator)
		return next(c)
	}
}
