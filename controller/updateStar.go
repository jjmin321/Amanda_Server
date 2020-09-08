package controller

import (
	"Amanda_Server/database"

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
	UserStar := &database.UserStar{}
	err := database.DB.Where("fk_object_idx = ? AND fk_user_id = ?", u.Idx, ID).Find(UserStar).Error
	if err != nil {
		CreateUserStar := &database.UserStar{FkObjectIdx: u.Idx, Star: u.Star, FkUserID: ID}
		err := database.DB.Create(CreateUserStar).Error
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"status":  500,
				"message": "평점 등록 실패",
			})
		}
	}
	err = database.DB.Model(UserStar).Where("fk_object_idx = ? AND fk_user_id = ?", u.Idx, ID).Update("star", u.Star).Error
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
