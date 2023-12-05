package service

import (
	articlecategories "FinalProject/features/article_categories"
	"FinalProject/features/article_categories/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArticleCategory(t *testing.T) {
	data := mocks.NewArticleCategoryDataInterface(t)
	service := New(data)
	article_category := []articlecategories.ArticleCategory{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll").Return(article_category, nil).Once()

		result, err := service.GetArticleCategories()

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll").Return(nil, errors.New("Get All Process Failed")).Once()

		result, err := service.GetArticleCategories()

		assert.Error(t, err)
		assert.EqualError(t, err, "Get All Process Failed")
		assert.Nil(t, result)
	})
}
