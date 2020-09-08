package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// ShowUserInfoParam - 파라미터 형식 정의 구조체
type ShowUserInfoParam struct {
	ID   string `json:"id" form:"id" query:"id"`
	Name string `json:"name" form:"name" query:"name"`
}

// ShowUserInfo - 유저 정보를 보여주는 메서드
func ShowUserInfo(c echo.Context) error {
	u := new(ShowUserInfoParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	User := &[]database.User{}
	err := database.DB.Where("user_id = ? OR name = ?", u.ID, u.Name).Find(User).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "멤버 리스트를 읽어오지 못했습니다.",
			"data":    "null",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "멤버 리스트를 읽어오는 데 성공했습니다.",
		"data":    User,
	})
}
