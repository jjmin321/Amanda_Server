package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// ShowMyInfo - 내 정보 가져오는 메서드
func ShowMyInfo(c echo.Context) error {
	ID := c.Get("ID")
	User := &[]database.User{}
	err := database.DB.Where("user_id = ?", ID).Error
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
