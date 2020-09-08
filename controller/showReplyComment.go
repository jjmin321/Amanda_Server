package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// ShowReplyCommentParam - 파라미터 형식 정의 구조체
type ShowReplyCommentParam struct {
	Idx string `json:"idx" form:"idx" query:"idx"`
}

// ShowReplyComment - 답글을 보여주는 메서드
func ShowReplyComment(c echo.Context) error {
	u := new(ShowReplyCommentParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	UserReplyComment := &[]database.UserReplyComment{}
	err := database.DB.Where("comment_idx = ?", u.Idx).Find(UserReplyComment).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":       500,
			"message":      "답글을 읽어오지 못했습니다.",
			"replyComment": "null",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":       200,
		"message":      "멤버 리스트를 읽어오는 데 성공했습니다.",
		"replyComment": UserReplyComment,
	})
}
