package service

import (
	articlecategories "FinalProject/features/article_categories"
	"errors"
)

type ArticleCategoryService struct {
	d articlecategories.ArticleCategoryDataInterface
}

func New(data articlecategories.ArticleCategoryDataInterface) articlecategories.ArticleCategoryServiceInterface {
	return &ArticleCategoryService{
		d: data,
	}
}

func (acs *ArticleCategoryService) GetArticleCategories() ([]articlecategories.ArticleCategory, error) {
	result, err := acs.d.GetAll()

	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}

	return result, nil
}

func (acs *ArticleCategoryService) CreateArticleCategory(newData articlecategories.ArticleCategory) (*articlecategories.ArticleCategory, error) {
	result, err := acs.d.Insert(newData)

	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}

	return result, nil
}

func (acs *ArticleCategoryService) GetArticleCategory(id int) ([]articlecategories.ArticleCategory, error) {
	result, err := acs.d.GetByID(id)

	if err != nil {
		return nil, errors.New("Get By Id Process Failed")
	}

	return result, nil
}

func (acs *ArticleCategoryService) UpdateArticleCategory(newData articlecategories.UpdateArticleCategory, id int) (bool, error) {
	result, err := acs.d.Update(newData, id)

	if err != nil {
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}

func (acs *ArticleCategoryService) DeleteArticleCategory(id int) (bool, error) {
	result, err := acs.d.Delete(id)

	if err != nil {
		return false, errors.New("Delete Process Failed")
	}

	return result, nil
}
