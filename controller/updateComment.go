package controller

import "github.com/labstack/echo"

// UpdateCommentParam - 파라미터 형식 정의 구조체
type UpdateCommentParam struct {
	Idx  uint `json:"idx" form:"idx" query:"idx"`
	Star uint `json:"star" form:"star" query:"star"`
}

func UpdateComment(c echo.Context) {
	u := new(UpdateStarParam)
	if err := c.Bind(u); err != nil {
		return err
	}
}
