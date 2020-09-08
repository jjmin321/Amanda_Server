package controller

import (
	"Amanda_Server/database"

	"github.com/labstack/echo"
)

// SignUpParam - 파라미터 형식 정의 구조체
type SignUpParam struct {
	ID        string `json:"id" form:"id" query:"id"`
	Pw        string `json:"pw" form:"pw" query:"pw"`
	Name      string `json:"name" form:"name" query:"name"`
	Image     string `json:"image" form:"image" query:"image"`
	IsManager bool   `json:"manager" form:"mamager" query:"manager"`
}

// SignUp - 회원가입
func SignUp(c echo.Context) error {
	u := new(SignUpParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	if u.ID == "" || u.Pw == "" || u.Name == "" {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "모든 값을 입력해주세요",
		})
	}
	User := &database.User{UserID: u.ID, Pw: u.Pw, Name: u.Name, IsManager: u.IsManager}
	err := database.DB.Create(User).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "회원가입 실패",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "회원가입이 완료되었습니다",
	})
}
