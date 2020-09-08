package controller

import (
	"Amanda_Server/database"
	"Amanda_Server/library/jwt"

	"github.com/labstack/echo"
)

// SignInParam - 파라미터 형식 정의 구조체
type SignInParam struct {
	ID string `json:"id" form:"id" query:"id"`
	Pw string `json:"pw" form:"pw" query:"pw"`
}

func SignIn(c echo.Context) error {
	u := new(SignInParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	User := &database.User{}
	err := database.DB.Where("user_id = ? AND pw = ?", u.ID, u.Pw).Error
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":       400,
			"message":      "일치하는 회원이 없습니다",
			"refreshToken": "null",
			"accessToken":  "null",
		})
	}
	refreshToken, err := jwt.CreateRefreshToken(User.UserID, User.Pw)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":       500,
			"message":      "refreshToken 생성 중 에러",
			"refreshToken": "null",
			"accessToken":  "null",
		})
	}
	accessToken, err := jwt.CreateAccessToken(User.UserID, User.Pw, User.IsManager)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":       500,
			"message":      "accessToken 생성 중 에러",
			"refreshToken": refreshToken,
			"accessToken":  "null",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":       200,
		"message":      "토큰 발급 완료",
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})
}
