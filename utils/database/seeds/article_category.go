package seeds

import (
	articlecategories "FinalProject/features/article_categories"
	helperSlug "FinalProject/helper/slug"
	"time"

	"gorm.io/gorm"
)

func CreateArticleCategory(db *gorm.DB, name string) error {
	slug := helperSlug.New()
	slugDb := slug.GenerateSlug(name)
	return db.Create(&articlecategories.ArticleCategory{Name: name, Slug: slugDb, CreatedAt: time.Now(), UpdatedAt: time.Now()}).Error
}
