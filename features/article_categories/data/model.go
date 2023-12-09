package data

import (
	"FinalProject/features/articles/data"

	"gorm.io/gorm"
)

type ArticleCategory struct {
	*gorm.Model
	Name     string         `gorm:"column:name;type:varchar(255)"`
	Slug     string         `gorm:"column:slug;type:varchar(255)"`
	Articles []data.Article `gorm:"foreignKey:CategoryID"`
}

func (ArticleCategory) TableName() string {
	return "article_categories"
}
