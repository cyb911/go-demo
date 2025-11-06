package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint      `gorm:"index"`
	User          *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments      []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommentCount  int64     `gorm:"<-:false;column:comment_count;-:migration"` // 非持久化字段,-:migration：跳过迁移创建列
	CommentStatus string    `gorm:"size:1;not null;default:'0';comment:0:无评论/1:有评论"`
}

// AfterCreate hook 方法，新增Post后，自动更新用户文章数量
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + 1")).Error
}
