package service

import (
	"FinalProject/features/articles"
	"FinalProject/features/articles/mocks"
	"errors"
	"mime/multipart"

	mockHelper "FinalProject/helper/mocks"
	mockUtil "FinalProject/utils/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArticles(t *testing.T) {
	data := mocks.NewArticleDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug, cld)
	article := []articles.ArticleInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll", "joko", "depresi", 1, 5).Return(article, nil).Once()

		result, err := service.GetArticles("joko", "depresi", 1, 5)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll", "joko", "depresi", 1, 5).Return(nil, errors.New("Get All Process Failed")).Once()

		result, err := service.GetArticles("joko", "depresi", 1, 5)

		assert.Error(t, err)
		assert.EqualError(t, err, "Get All Process Failed")
		assert.Nil(t, result)
	})
}

func TestGetArticle(t *testing.T) {
	data := mocks.NewArticleDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug, cld)
	article := []articles.ArticleInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByID", 1).Return(article, nil).Once()

		result, err := service.GetArticle(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(nil, errors.New("Get By ID Process Failed")).Once()

		result, err := service.GetArticle(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Get By ID Process Failed")
		assert.Nil(t, result)
	})
}

func TestCreateArticle(t *testing.T) {
	data := mocks.NewArticleDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug, cld)
	var mockFile multipart.File
	article := articles.Article{
		CategoryID:    1,
		Title:         "Depresi",
		Content:       "Mencegah Depresi",
		ThumbnailFile: mockFile,
		ThumbnailUrl:  "https",
		Status:        "Pending",
		Slug:          "123-depresi",
	}

	t.Run("Success Insert", func(t *testing.T) {
		cld.On("UploadImageHelper", article.ThumbnailFile).Return(article.ThumbnailUrl, nil).Once()
		slug.On("GenerateSlug", article.Title).Return(article.Slug).Once()
		data.On("Insert", article).Return(&article, nil).Once()

		result, err := service.CreateArticle(article)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, article.CategoryID, result.CategoryID)
		assert.Equal(t, article.Title, result.Title)
		assert.Equal(t, article.Content, result.Content)
		assert.Equal(t, article.ThumbnailUrl, result.ThumbnailUrl)
		assert.Equal(t, article.Status, result.Status)
		assert.Equal(t, article.Slug, result.Slug)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Insert", article).Return(nil, errors.New("Insert Process Failed")).Once()
		cld.On("UploadImageHelper", article.ThumbnailFile).Return(article.ThumbnailUrl, nil).Once()
		slug.On("GenerateSlug", article.Title).Return(article.Slug).Once()

		result, err := service.CreateArticle(article)

		assert.Error(t, err)
		assert.EqualError(t, err, "Insert Process Failed")
		assert.Nil(t, result)
	})
}

func TestUpdateArticle(t *testing.T) {
	data := mocks.NewArticleDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug, cld)
	var mockFile multipart.File
	article := articles.UpdateArticle{
		Title:         "Depresi",
		Content:       "Mencegah Depresi",
		ThumbnailFile: mockFile,
		Slug:          "123-depresi",
	}

	t.Run("Success Update", func(t *testing.T) {
		data.On("Update", article, 1).Return(true, nil).Once()
		cld.On("UploadImageHelper", article.ThumbnailFile).Return(article.ThumbnailUrl, nil).Once()
		slug.On("GenerateSlug", article.Title).Return(article.Slug).Once()

		result, err := service.UpdateArticle(article, 1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Update", article, 1).Return(false, errors.New("Update Process Failed")).Once()
		cld.On("UploadImageHelper", article.ThumbnailFile).Return(article.ThumbnailUrl, nil).Once()
		slug.On("GenerateSlug", article.Title).Return(article.Slug).Once()

		result, err := service.UpdateArticle(article, 1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Update Process Failed")
		assert.Equal(t, false, result)
	})
}

func TestRejectArticle(t *testing.T) {
	data := mocks.NewArticleDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug, cld)

	t.Run("Success Reject", func(t *testing.T) {
		data.On("Reject", 1).Return(true, nil).Once()

		result, err := service.RejectArticle(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Reject", 1).Return(false, errors.New("Reject Process Failed")).Once()

		result, err := service.RejectArticle(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Reject Process Failed")
		assert.Equal(t, false, result)
	})
}

func TestPublishArticle(t *testing.T) {
	data := mocks.NewArticleDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug, cld)

	t.Run("Success Publish", func(t *testing.T) {
		data.On("Publish", 1).Return(true, nil).Once()

		result, err := service.PublishArticle(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Publish", 1).Return(false, errors.New("Publish Process Failed")).Once()

		result, err := service.PublishArticle(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Publish Process Failed")
		assert.Equal(t, false, result)
	})
}

func TestDashboardArticle(t *testing.T) {
	data := mocks.NewArticleDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug, cld)
	dashboard := articles.ArticleDashboard{
		TotalArticle:        10,
		TotalArticleBaru:    4,
		TotalArticlePending: 6,
	}

	t.Run("Success Get Dashboard", func(t *testing.T) {
		data.On("ArticleDashboard").Return(dashboard, nil).Once()

		result, err := service.ArticleDashboard()

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, dashboard.TotalArticle, result.TotalArticle)
		assert.Equal(t, dashboard.TotalArticleBaru, result.TotalArticleBaru)
		assert.Equal(t, dashboard.TotalArticlePending, result.TotalArticlePending)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("ArticleDashboard").Return(articles.ArticleDashboard{}, errors.New("Process Failed")).Once()

		result, err := service.ArticleDashboard()

		assert.Error(t, err)
		assert.EqualError(t, err, "Process Failed")
		assert.NotNil(t, result)
	})
}

func TestGetArticleByDoctor(t *testing.T) {
	data := mocks.NewArticleDataInterface(t)
	cld := mockUtil.NewCloudinaryInterface(t)
	slug := mockHelper.NewSlugInterface(t)
	service := New(data, slug, cld)
	article := []articles.ArticleInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByIDDoctor", 1).Return(article, nil).Once()

		result, err := service.GetArticleByDoctor(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
		cld.AssertExpectations(t)
		slug.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByIDDoctor", 1).Return(nil, errors.New("Get By ID Doctor Process Failed")).Once()

		result, err := service.GetArticleByDoctor(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Get By ID Doctor Process Failed")
		assert.Nil(t, result)
	})
}
