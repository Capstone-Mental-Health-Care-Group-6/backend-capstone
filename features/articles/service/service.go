package service

import (
	"FinalProject/features/articles"
	"errors"
)

type ArticleService struct {
	d articles.ArticleDataInterface
}

func New(data articles.ArticleDataInterface) articles.ArticleServiceInterface {
	return &ArticleService{
		d: data,
	}
}

func (as *ArticleService) GetArticles(name, kategori string, timePublication int) ([]articles.ArticleInfo, error) {
	result, err := as.d.GetAll(name, kategori, timePublication)
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
	result, err := as.d.Insert(newData)
	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}
	return result, nil
}

func (as *ArticleService) UpdateArticle(newData articles.UpdateArticle, id int) (bool, error) {
	result, err := as.d.Update(newData, id)

	if err != nil {
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}

func (as *ArticleService) DenyArticle(id int) (bool, error) {
	result, err := as.d.Deny(id)

	if err != nil {
		return false, errors.New("Denny Process Failed")
	}

	return result, nil
}

func (as *ArticleService) ApproveArticle(id int) (bool, error) {
	result, err := as.d.Approve(id)

	if err != nil {
		return false, errors.New("Approve Process Failed")
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
