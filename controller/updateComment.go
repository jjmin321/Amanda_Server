package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// UpdateCommentParam - 파라미터 형식 정의 구조체
type UpdateCommentParam struct {
	Idx     uint   `json:"idx" form:"idx" query:"idx"`
	Comment string `json:"comment" form:"comment" query:"comment"`
}

// UpdateComment - 댓글 수정 메서드
func UpdateComment(c echo.Context) error {
	u := new(UpdateCommentParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	ID := c.Get("ID").(string)
	UserComment := &database.UserComment{}
	err := database.DB.Model(UserComment).Update("comment", u.Comment).Where("idx = ?", u.Idx).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "댓글 수정 실패",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "댓글 수정 성공",
	})
}
