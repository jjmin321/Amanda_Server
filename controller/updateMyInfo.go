package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// UpdateMyInfoParam - 파라미터 형식 정의 구조체
type UpdateMyInfoParam struct {
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"description" form:"description" query:"description"`
}

// UpdateMyInfo - 내 정보 변경 메서드
func UpdateMyInfo(c echo.Context) error {
	ID := c.Get("ID")
	u := new(UpdateMyInfoParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	User := &database.User{}
	if u.Name != "" {
		err := database.DB.Model(User).Update("name", u.Name).Where("user_id = ?", ID).Error
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"status":  500,
				"message": "이름 변경 실패",
			})
		}
	}
	if u.Description != "" {
		err := database.DB.Model(User).Update("description", u.Description).Where("user_id = ?", ID).Error
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"status":  500,
				"message": "소개 변경 실패",
			})
		}
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "성공적으로 변경되었습니다.",
	})
}
