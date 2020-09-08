package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// CreateCommentParam - 파라미터 형식 정의 구조체
type CreateCommentParam struct {
	ID      string `json:"id" form:"id" query:"id"`
	Comment string `json:"comment" form:"comment" query:"comment"`
}

// CreateComment - 댓글을 작성하는 메서드
func CreateComment(c echo.Context) error {
	ID := c.Get("ID").(string)
	u := new(CreateCommentParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	UserComment := &database.UserComment{FkUserID: ID, FkObjectID: u.ID, Comment: u.Comment}
	err := database.DB.Create(UserComment).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "댓글 달기 실패",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "댓글 달기 성공",
	})
}
