package database

// UserStar 유저 평가 테이블
type UserStar struct {
	FkUserIdx uint `gorm:"primary_key;" json:"fk_user_idx"`
	Star      uint `gorm:"not null" sql:"DEFAULT:0" json:"star"`
	People    uint `gorm:"not null; auto_increment:true" json:"people"`
}
