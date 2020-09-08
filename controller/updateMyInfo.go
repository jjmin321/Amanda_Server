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

	if u.Name != "" {
		User := &database.User{}
		err := database.DB.Model(User).Where("user_id = ?", ID).Update("name", u.Name).Error
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"status":  500,
				"message": "이름 변경 실패",
			})
		}
	}
	if u.Description != "" {
		User := &database.User{}
		err := database.DB.Model(User).Where("user_id = ?", ID).Update("description", u.Description).Error
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
