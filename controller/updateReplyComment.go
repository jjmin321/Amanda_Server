package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// UpdateReplyCommentParam - 파라미터 형식 정의 구조체
type UpdateReplyCommentParam struct {
	Idx     uint   `json:"idx" form:"idx" query:"idx"`
	Comment string `json:"comment" form:"comment" query:"comment"`
}

// UpdateReplyComment - 댓글 수정 메서드
func UpdateReplyComment(c echo.Context) error {
	u := new(UpdateCommentParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	ID := c.Get("ID").(string)
	UserReplyComment := &database.UserReplyComment{}
	err := database.DB.Model(UserReplyComment).Where("idx = ? AND fk_user_id = ?", u.Idx, ID).Update("comment", u.Comment).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "답글 수정 실패",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "답글 수정 성공",
	})
}
