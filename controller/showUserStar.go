package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// ShowUserStarParam - 파라미터 형식 정의 구조체
type ShowUserStarParam struct {
	Idx uint `json:"idx" form:"idx" query:"idx"`
}

// ShowUserStar - 유저의 평점과 참여자 수를 읽어오는 메소드
func ShowUserStar(c echo.Context) error {
	u := new(ShowUserStarParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	var starCnt uint
	starCnt = 0
	UserStar := &[]database.UserStar{}
	err := database.DB.Where("fk_object_idx = ?", u.Idx).Find(UserStar).Error
	for _, v := range *UserStar {
		starCnt += v.Star
	}

	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "평점 정보를 읽어오지 못했습니다.",
			"평점":      0,
			"참여자 수":   "알 수 없음",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "평점 정보를 읽어오는 데 성공했습니다.",
		"평점":      starCnt,
		"참여자 수":   len(*UserStar),
	})
}
