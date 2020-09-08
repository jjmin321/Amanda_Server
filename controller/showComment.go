package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// ShowCommentParam - 파라미터 형식 정의 구조체
type ShowCommentParam struct {
	ID string `json:"id" form:"id" query:"id"`
}

// ShowComment - 유저 정보를 보여주는 메서드
func ShowComment(c echo.Context) error {
	u := new(ShowCommentParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	UserComment := &[]database.UserComment{}
	err := database.DB.Where("fk_object_id = ?", u.ID).Find(UserComment).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "멤버 리스트를 읽어오지 못했습니다.",
			"comment": "null",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "멤버 리스트를 읽어오는 데 성공했습니다.",
		"comment": UserComment,
	})
}
