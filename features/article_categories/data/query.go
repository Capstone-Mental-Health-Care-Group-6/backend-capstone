package data

import (
	articlecategories "FinalProject/features/article_categories"
	"errors"

	"gorm.io/gorm"
)

type ArticleCategoryData struct {
	db *gorm.DB
}

func New(db *gorm.DB) articlecategories.ArticleCategoryDataInterface {
	return &ArticleCategoryData{
		db: db,
	}
}

func (acd *ArticleCategoryData) GetAll() ([]articlecategories.ArticleCategory, error) {
	var listArticleCategory = []articlecategories.ArticleCategory{}

	if err := acd.db.Find(&listArticleCategory).Error; err != nil {
		return nil, err
	}

	return listArticleCategory, nil
}

func (acd *ArticleCategoryData) Insert(newData articlecategories.ArticleCategory) (*articlecategories.ArticleCategory, error) {
	var dbData = new(ArticleCategory)
	dbData.Name = newData.Name
	dbData.Slug = newData.Slug

	if err := acd.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (acd *ArticleCategoryData) GetByID(id int) ([]articlecategories.ArticleCategory, error) {
	var listArticleCategory = []articlecategories.ArticleCategory{}

	if err := acd.db.Where("id = ?", id).Find(&listArticleCategory).Error; err != nil {
		return nil, err
	}

	return listArticleCategory, nil
}

func (acd *ArticleCategoryData) Update(newData articlecategories.UpdateArticleCategory, id int) (bool, error) {
	var qry = acd.db.Where("id = ?", id).Updates(ArticleCategory{Name: newData.Name, Slug: newData.Slug})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}

func (acd *ArticleCategoryData) Delete(id int) (bool, error) {
	var deleteData = new(ArticleCategory)

	if err := acd.db.Where("id = ?", id).Delete(&deleteData).Error; err != nil {
		return false, err
	}
	return true, nil
}
