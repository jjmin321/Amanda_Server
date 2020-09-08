package controller

import "github.com/labstack/echo"

// UpdateCommentParam - 파라미터 형식 정의 구조체
type UpdateCommentParam struct {
	Idx     uint   `json:"idx" form:"idx" query:"idx"`
	Comment string `json:"comment" form:"comment" query:"comment"`
}

func UpdateComment(c echo.Context) error {
	u := new(UpdateCommentParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "댓글 달기 성공",
	})
}
