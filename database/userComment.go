package database

import "time"

// UserComment 평가하는 댓글 테이블
type UserComment struct {
	Idx         uint      `gorm:"primary_key; auto_increment:true" json:"idx"`
	FkUserID    string    `gorm:"type:varchar(255);not null" json:"fk_user_id"`
	FkObjectIdx uint      `gorm:"not null" json:"fk_object_idx"`
	Comment     string    `gorm:"type:varchar(255);not null" json:"comment"`
	CreatedAt   time.Time `gorm:"not null" sql:"DEFAULT:current_timestamp" json:"created_at"`
}
