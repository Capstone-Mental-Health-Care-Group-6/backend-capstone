package data

import (
	"FinalProject/features/articles"
	"errors"
	"time"

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

func (ad *ArticleData) GetAll(name, kategori string, timePublication, limit int) ([]articles.ArticleInfo, error) {
	var listArticle = []articles.ArticleInfo{}
	var now time.Time
	var qry = ad.db.Table("articles").Select("articles.*,users.name as user_name,article_categories.name as category_name").
		Joins("JOIN users on users.id = articles.user_id").
		Joins("JOIN article_categories ON article_categories.id = articles.category_id").
		Order("articles.created_at DESC")

	now = time.Now()

	if name != "" {
		qry.Where("articles.name LIKE ?", "%"+name+"%").Or("users.name LIKE ?", "%"+name+"%")
	}

	if kategori != "" {
		qry.Where("category_name = ?", kategori)
	}

	switch timePublication {
	case 1:
		before := now.AddDate(0, 0, -7)
		qry.Where("articles.created_at BETWEEN ? and ?", before, now)
		break
	case 2:
		before := now.AddDate(0, 0, -30)
		qry.Where("articles.created_at BETWEEN ? and ?", before, now)
		break
	}

	if limit != 0 {
		qry.Limit(limit)
	}

	if err := qry.Scan(&listArticle).Error; err != nil {
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
	dbData.Thumbnail = newData.ThumbnailUrl
	dbData.Status = newData.Status
	dbData.Slug = newData.Slug

	if err := ad.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (ad *ArticleData) Update(newData articles.UpdateArticle, id int) (bool, error) {
	var qry = ad.db.Table("articles").Where("id = ?", id).Updates(Article{Title: newData.Title, Content: newData.Content, Thumbnail: newData.ThumbnailUrl, Slug: newData.Slug})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}

func (ad *ArticleData) Reject(id int) (bool, error) {
	var qry = ad.db.Table("articles").Where("id = ?", id).Updates(Article{Status: "Reject"})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}

func (ad *ArticleData) Publish(id int) (bool, error) {
	var qry = ad.db.Table("articles").Where("id = ?", id).Updates(Article{Status: "Publish"})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}

func (ad *ArticleData) ArticleDashboard() (articles.ArticleDashboard, error) {
	var dashboardArticle articles.ArticleDashboard

	tArticle, tArticleBaru, tArticlePending := ad.getTotalArticle()

	dashboardArticle.TotalArticle = tArticle
	dashboardArticle.TotalArticleBaru = tArticleBaru
	dashboardArticle.TotalArticlePending = tArticlePending

	return dashboardArticle, nil
}

func (ad *ArticleData) getTotalArticle() (int, int, int) {
	var totalArticle int64
	var totalArticleBaru int64
	var totalArticlePending int64

	var now = time.Now()
	var before = now.AddDate(0, 0, -30)

	var _ = ad.db.Table("articles").Count(&totalArticle)
	var _ = ad.db.Table("articles").Where("created_at BETWEEN ? and ?", before, now).Count(&totalArticleBaru)
	var _ = ad.db.Table("articles").Where("status = ?", "Pending").Count(&totalArticlePending)

	totalArticleInt := int(totalArticle)
	totalArticleBaruInt := int(totalArticleBaru)
	totalArticlePendingInt := int(totalArticlePending)

	return totalArticleInt, totalArticleBaruInt, totalArticlePendingInt
}

func (ad *ArticleData) GetByIDDoctor(id int) ([]articles.ArticleInfo, error) {
	var listArticle = []articles.ArticleInfo{}
	var qry = ad.db.Table("articles").Select("articles.*,users.name as user_name,article_categories.name as category_name").
		Joins("LEFT JOIN users on users.id = articles.user_id").Joins("LEFT JOIN article_categories ON article_categories.id = articles.category_id").
		Where("users.id = ?", id).
		Where("articles.deleted_at is null").
		Order("articles.created_at DESC")

	if err := qry.Scan(&listArticle).Error; err != nil {
		logrus.Info("DB error : ", err.Error())
		return nil, err
	}
	return listArticle, nil
}
