package controller

import (
	"Amanda_Server/database"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// UpdateProfileImage - 프로필 사진 등록
func UpdateProfileImage(c echo.Context) error {
	ID := c.Get("ID").(string)
	Pw := c.Get("Pw").(string)
	file, err := c.FormFile("image")
	t := time.Now()
	now := strconv.Itoa(int(t.Unix()))
	fileName := ID + now + file.Filename
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "파일을 읽는 데 실패하였습니다. 다시 시도해주세요.",
		})
	}
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "파일을 여는 데 실패하였습니다. 다시 시도해주세요.",
		})
	}
	dst, err := os.Create("ProfileImage/" + fileName)
	defer dst.Close()
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "파일을 생성하는 데 실패하였습니다. 다시 시도해주세요.",
		})
	}
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "파일을 저장하는 데 실패하였습니다. 다시 시도해주세요.",
		})
	}
	// err = model.UpdateImageInTeamMember(Idx, fileName)
	User := &database.User{}
	err = database.DB.Model(User).Where("user_id = ? AND pw = ?", ID, Pw).Update("image", fileName).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "경로를 저장하는 데 실패하였습니다. 다시 시도해주세요.",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "프로필 사진 등록에 성공하셨습니다.",
	})
}
