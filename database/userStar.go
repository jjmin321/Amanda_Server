package database

// UserStar 유저 평가 테이블
type UserStar struct {
	FkObjectIdx uint `gorm:"primary_key;" json:"fk_object_idx"`
	Star        uint `gorm:"not null" json:"star"`
	FkUserID    uint `gorm:"not null;" json:"fk_user_id"`
}
