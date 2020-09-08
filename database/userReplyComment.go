package database

import "time"

// UserReplyComment 평가하는 댓글의 답글 테이블
type UserReplyComment struct {
	Idx         uint      `gorm:"primary_key; auto_increment:true" json:"idx"`
	CommentIdx  uint      `gorm:"not null" json:"comment_idx`
	FkUserIdx   uint      `gorm:"not null" json:"fk_user_idx"`
	FkObjectIdx uint      `gorm:"not null" json:"fk_object_idx"`
	Comment     string    `gorm:"not null" json:"comment"`
	CreatedAt   time.Time `gorm:"not null" sql:"DEFAULT:current_timestamp" json:"created_at"`
}
