package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// CreateReplyCommentParam - 파라미터 형식 정의 구조체
type CreateReplyCommentParam struct {
	Idx     uint   `json:"idx" form:"idx" query:"idx"`
	Comment string `json:"comment" form:"comment" query:"comment"`
}

// CreateReplyComment - 답글을 작성하는 메서드
func CreateReplyComment(c echo.Context) error {
	ID := c.Get("ID").(string)
	u := new(CreateReplyCommentParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	UserReplyComment := &database.UserReplyComment{CommentIdx: u.Idx, FkUserID: ID, Comment: u.Comment}
	err := database.DB.Create(UserReplyComment).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "답글 달기 실패",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "답글 달기 성공",
	})
}
