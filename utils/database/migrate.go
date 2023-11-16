package database

import (
	DataArticleCategory "FinalProject/features/article_categories/data"
	DataArticle "FinalProject/features/articles/data"
	DataUser "FinalProject/features/users/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(DataUser.User{})
	db.AutoMigrate(DataArticleCategory.ArticleCategory{})
	db.AutoMigrate(DataArticle.Article{})
	db.AutoMigrate(DataUser.Patient{})
}
