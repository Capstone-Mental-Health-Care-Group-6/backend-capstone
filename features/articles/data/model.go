package data

import (
	"gorm.io/gorm"
)

type Article struct {
	*gorm.Model
	CategoryID uint   `gorm:"column:category_id"`
	UserID     uint   `gorm:"column:user_id"`
	Title      string `gorm:"column:title;type:varchar(255)"`
	Content    string `gorm:"column:content;type:text"`
	Thumbnail  string `gorm:"column:thumbnail;type:varchar(255)"`
	Status     string `gorm:"column:status;type:enum('Publish','Pending','Reject')"`
	Slug       string `gorm:"column:slug;type:varchar(255)"`
}
