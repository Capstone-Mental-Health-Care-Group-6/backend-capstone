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

func (as *ArticleService) GetArticles() ([]articles.ArticleInfo, error) {
	result, err := as.d.GetAll()
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

func (as *ArticleService) DeleteArticle(id int) (bool, error) {
	result, err := as.d.Delete(id)

	if err != nil {
		return false, errors.New("Delete Process Failed")
	}

	return result, nil
}
