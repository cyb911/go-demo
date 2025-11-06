package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text; not null"`
	UserID  uint   `gorm:"index; not null"`
	PostID  uint   `gorm:"index; not null"`
	User    *User
	Post    *Post
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	if c.PostID == 0 && c.ID != 0 {
		var tmp struct{ PostID uint } // 声明一个临时的匿名结构体，用于接受查询数据
		err := tx.Model(&Comment{}).Select("post_id").Where("id = ?", c.ID).Take(&tmp).Error
		if err != nil {
			return err
		}
		c.PostID = tmp.PostID
	}
	return nil
}

// AfterDelete 删除后，如果该 Post 没有评论了，把评论状态改为“无评论”
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).
		Count(&count).Error

	if err != nil {
		return nil
	}

	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).
			UpdateColumn("comment_status", gorm.Expr("comment_status = ?", "0")).Error
	}
	return nil
}

// AfterCreate 新增评论，把Post评论状态改为“有评论”
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&Post{}).Where("id = ?", c.PostID).
		UpdateColumn("comment_status", gorm.Expr("comment_status = ?", "1")).Error
}
