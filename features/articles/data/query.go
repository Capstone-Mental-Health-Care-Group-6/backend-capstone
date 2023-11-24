package data

import (
	"FinalProject/features/articles"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleData struct {
	db *gorm.DB
}

func New(db *gorm.DB) articles.ArticleDataInterface {
	return &ArticleData{
		db: db,
	}
}

func (ad *ArticleData) GetAll() ([]articles.ArticleInfo, error) {
	var listArticle = []articles.ArticleInfo{}
	var qry = ad.db.Table("articles").Select("articles.*,users.name as user_name,article_categories.name as category_name").
		Joins("JOIN users on users.id = articles.user_id").
		Joins("JOIN article_categories ON article_categories.id = articles.category_id").
		Where("articles.deleted_at is null").
		Scan(&listArticle)

	if err := qry.Error; err != nil {
		logrus.Info("DB error : ", err.Error())
		return nil, err
	}

	return listArticle, nil
}

func (ad *ArticleData) GetByID(id int) ([]articles.ArticleInfo, error) {
	var listArticle = []articles.ArticleInfo{}
	var qry = ad.db.Table("articles").Select("articles.*,users.name as user_name,article_categories.name as category_name").
		Joins("LEFT JOIN users on users.id = articles.user_id").Joins("LEFT JOIN article_categories ON article_categories.id = articles.category_id").
		Where("articles.id = ?", id).
		Where("articles.deleted_at is null").
		Scan(&listArticle)

	if err := qry.Error; err != nil {
		logrus.Info("DB error : ", err.Error())
		return nil, err
	}
	return listArticle, nil
}

func (ad *ArticleData) Insert(newData articles.Article) (*articles.Article, error) {
	var dbData = new(Article)
	dbData.CategoryID = newData.CategoryID
	dbData.UserID = newData.UserID
	dbData.Title = newData.Title
	dbData.Content = newData.Content
	dbData.Thumbnail = newData.Thumbnail
	dbData.Status = newData.Status
	dbData.Slug = newData.Slug

	if err := ad.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (ad *ArticleData) Update(newData articles.UpdateArticle, id int) (bool, error) {
	var qry = ad.db.Table("articles").Where("id = ?", id).Updates(Article{Title: newData.Title, Content: newData.Content, Thumbnail: newData.Thumbnail, Slug: newData.Slug})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}

func (ad *ArticleData) Delete(id int) (bool, error) {
	var deleteData = new(Article)

	if err := ad.db.Where("id = ?", id).Delete(&deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}
