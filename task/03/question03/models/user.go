package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	PostCount int64  `gorm:"not null;default:0;comment:文章数量统计"`
	Posts     []Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // OnDelete:SET NULL 删除时，外键设置为null
}
