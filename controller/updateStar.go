package controller

import (
	"Amanda_Server/database"
	"log"

	"github.com/labstack/echo"
)

// UpdateStarParam - 파라미터 형식 정의 구조체
type UpdateStarParam struct {
	Idx  uint `json:"idx" form:"idx" query:"idx"`
	Star uint `json:"star" form:"star" query:"star"`
}

// UpdateStar - 유저 평점 등록 메서드
func UpdateStar(c echo.Context) error {
	u := new(UpdateStarParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	if u.Idx == 0 || u.Star == 0 {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "모든 값을 입력해주세요",
		})
	}
	ID := c.Get("ID").(string)
	log.Print(ID)
	UserStar := &database.UserStar{FkObjectIdx: u.Idx, Star: u.Star, FkUserID: ID}
	err := database.DB.Create(UserStar).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "평점 등록 실패",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "평점 등록 완료",
	})
}
