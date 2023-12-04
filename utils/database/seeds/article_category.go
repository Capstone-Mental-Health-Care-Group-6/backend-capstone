package seeds

import (
	articlecategories "FinalProject/features/article_categories"

	"gorm.io/gorm"
)

func CreateArticleCategory(db *gorm.DB, name, slug string) error {
	return db.Create(articlecategories.ArticleCategory{Name: name, Slug: slug}).Error
}
