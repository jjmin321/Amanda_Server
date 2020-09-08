package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// CheckidParam - 파라미터 형식 정의 구조체
type CheckidParam struct {
	ID string `json:"id" form:"id" query:"id"`
}

// Checkid - 아이디 중복 체크
func Checkid(c echo.Context) error {
	u := new(ShowUserInfoParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	User := &database.User{}
	err := database.DB.Where("user_id = ?", u.ID).Find(User).Error
	if err != nil {
		return c.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "사용할 수 있는 아이디입니다",
		})
	}
	return c.JSON(400, map[string]interface{}{
		"status":  400,
		"message": "이미 사용중인 아이디입니다.",
	})
}
