package service

import (
	"FinalProject/features/articles"
	"FinalProject/helper/slug"
	"FinalProject/utils/cloudinary"
	"errors"
)

type ArticleService struct {
	d     articles.ArticleDataInterface
	slug  slug.SlugInterface
	cloud cloudinary.CloudinaryInterface
}

func New(data articles.ArticleDataInterface, slug slug.SlugInterface, cld cloudinary.CloudinaryInterface) articles.ArticleServiceInterface {
	return &ArticleService{
		d:     data,
		slug:  slug,
		cloud: cld,
	}
}

func (as *ArticleService) GetArticles(name, kategori string, timePublication, limit int) ([]articles.ArticleInfo, error) {
	result, err := as.d.GetAll(name, kategori, timePublication, limit)
	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}
	return result, nil
}

func (as *ArticleService) GetArticle(id int) ([]articles.ArticleInfo, error) {
	result, err := as.d.GetByID(id)
	if err != nil {
		return nil, errors.New("Get By ID Process Failed")
	}
	return result, nil
}

func (as *ArticleService) CreateArticle(newData articles.Article) (*articles.Article, error) {
	slug := as.slug.GenerateSlug(newData.Title)
	uploadUrlThumbnail, err := as.cloud.UploadImageHelper(newData.ThumbnailFile)

	newData.Slug = slug
	newData.ThumbnailUrl = uploadUrlThumbnail
	result, err := as.d.Insert(newData)
	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}
	return result, nil
}

func (as *ArticleService) UpdateArticle(newData articles.UpdateArticle, id int) (bool, error) {
	slug := as.slug.GenerateSlug(newData.Title)
	uploadUrlThumbnail, err := as.cloud.UploadImageHelper(newData.ThumbnailFile)

	newData.Slug = slug
	newData.ThumbnailUrl = uploadUrlThumbnail
	result, err := as.d.Update(newData, id)

	if err != nil {
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}

func (as *ArticleService) RejectArticle(id int) (bool, error) {
	result, err := as.d.Reject(id)

	if err != nil {
		return false, errors.New("Reject Process Failed")
	}

	return result, nil
}

func (as *ArticleService) PublishArticle(id int) (bool, error) {
	result, err := as.d.Publish(id)

	if err != nil {
		return false, errors.New("Publish Process Failed")
	}

	return result, nil
}

func (as *ArticleService) ArticleDashboard() (articles.ArticleDashboard, error) {
	res, err := as.d.ArticleDashboard()

	if err != nil {
		return res, errors.New("Process Failed")
	}

	return res, nil
}

func (as *ArticleService) GetArticleByDoctor(id int) ([]articles.ArticleInfo, error) {
	result, err := as.d.GetByIDDoctor(id)
	if err != nil {
		return nil, errors.New("Get By ID Doctor Process Failed")
	}
	return result, nil
}
