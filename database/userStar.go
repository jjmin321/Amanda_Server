package database

// UserStar 유저 평가 테이블
type UserStar struct {
	Idx         uint `gorm:"primary_key; auto_increment:true" json:"idx"`
	FkObjectIdx uint `gorm:"not null;" json:"fk_object_idx"`
	Star        uint `gorm:"not null" json:"star"`
	FkUserID    uint `gorm:"not null;" json:"fk_user_id"`
}
