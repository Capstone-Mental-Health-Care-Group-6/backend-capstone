package service

import (
	articlecategories "FinalProject/features/article_categories"
	"FinalProject/features/article_categories/mocks"
	mockHelper "FinalProject/helper/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArticleCategories(t *testing.T) {
	data := mocks.NewArticleCategoryDataInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug)
	article_category := []articlecategories.ArticleCategory{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll").Return(article_category, nil).Once()

		result, err := service.GetArticleCategories()

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll").Return(nil, errors.New("Get All Process Failed")).Once()

		result, err := service.GetArticleCategories()

		assert.Error(t, err)
		assert.EqualError(t, err, "Get All Process Failed")
		assert.Nil(t, result)
	})
}

func TestGetArticleCategory(t *testing.T) {
	data := mocks.NewArticleCategoryDataInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug)
	article_category := []articlecategories.ArticleCategory{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByID", 1).Return(article_category, nil).Once()

		result, err := service.GetArticleCategory(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(nil, errors.New("Get By ID Process Failed")).Once()

		result, err := service.GetArticleCategory(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Get By ID Process Failed")
		assert.Nil(t, result)
	})
}

func TestCreateArticleCategory(t *testing.T) {
	data := mocks.NewArticleCategoryDataInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug)
	article_categories := articlecategories.ArticleCategory{
		Name: "Kecemasan",
		Slug: "642-Kecemasan",
	}

	t.Run("Success Insert", func(t *testing.T) {
		data.On("Insert", article_categories).Return(&article_categories, nil).Once()
		slug.On("GenerateSlug", article_categories.Name).Return(article_categories.Slug).Once()

		result, err := service.CreateArticleCategory(article_categories)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.Name, article_categories.Name)
		assert.Equal(t, result.Slug, article_categories.Slug)
		data.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Insert", article_categories).Return(nil, errors.New("Insert Process Failed")).Once()
		slug.On("GenerateSlug", article_categories.Name).Return(article_categories.Slug).Once()

		result, err := service.CreateArticleCategory(article_categories)
		assert.Error(t, err)
		assert.EqualError(t, err, "Insert Process Failed")
		assert.Nil(t, result)
	})
}

func TestUpdateArticleCategory(t *testing.T) {
	data := mocks.NewArticleCategoryDataInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug)
	article_categories := articlecategories.UpdateArticleCategory{
		Name: "Kecemasan",
		Slug: "642-Kecemasan",
	}

	t.Run("Success Update", func(t *testing.T) {
		slug.On("GenerateSlug", article_categories.Name).Return(article_categories.Slug).Once()
		data.On("Update", article_categories, 1).Return(true, nil).Once()

		result, err := service.UpdateArticleCategory(article_categories, 1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		slug.On("GenerateSlug", article_categories.Name).Return(article_categories.Slug).Once()
		data.On("Update", article_categories, 1).Return(false, errors.New("Update Process Failed")).Once()

		result, err := service.UpdateArticleCategory(article_categories, 1)
		assert.Error(t, err)
		assert.EqualError(t, err, "Update Process Failed")
		assert.Equal(t, false, result)
	})
}

func TestDeleteArticleCategory(t *testing.T) {
	data := mocks.NewArticleCategoryDataInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug)

	t.Run("Success Delete", func(t *testing.T) {
		data.On("Delete", 1).Return(true, nil).Once()

		result, err := service.DeleteArticleCategory(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Delete", 1).Return(false, errors.New("Delete Process Failed")).Once()

		result, err := service.DeleteArticleCategory(1)
		assert.Error(t, err)
		assert.EqualError(t, err, "Delete Process Failed")
		assert.Equal(t, false, result)
	})
}
