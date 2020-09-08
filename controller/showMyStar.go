package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// ShowMyStarInfo - 내거 먀간 푱좀 = 가져오는 메서드
func ShowMyStarInfo(c echo.Context) error {
	ID := c.Get("ID")
	UserStar := &[]database.UserStar{}
	err := database.DB.Where("fk_user_id = ?", ID).Find(UserStar).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "내 정보를 읽어오지 못했습니다.",
			"data":    "null",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "내 정보를 읽어오는 데 성공했습니다.",
		"data":    UserStar,
	})
}
