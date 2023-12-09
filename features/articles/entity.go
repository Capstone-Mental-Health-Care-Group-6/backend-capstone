package articles

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Article struct {
	ID            uint           `json:"id"`
	CategoryID    uint           `json:"category_id"`
	UserID        uint           `json:"user_id"`
	Title         string         `json:"title"`
	Content       string         `json:"content"`
	ThumbnailUrl  string         `json:"thumbnail_url"`
	ThumbnailFile multipart.File `json:"thumbnail"`
	Status        string         `json:"status"`
	Slug          string         `json:"slug"`
}

type ArticleInfo struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	UserName     string `json:"user_name"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Thumbnail    string `json:"thumbnail"`
	Status       string `json:"status"`
	Slug         string `json:"slug"`
}

type UpdateArticle struct {
	Title         string         `json:"title"`
	Content       string         `json:"content"`
	ThumbnailUrl  string         `json:"thumbnail_url"`
	ThumbnailFile multipart.File `json:"thumbnail"`
	Slug          string         `json:"slug"`
}

type ThumbnailDataModel struct {
	ThumbnailPhoto multipart.File `json:"thumbnail"`
}

type ArticleDashboard struct {
	TotalArticle        int `json:"total_article"`
	TotalArticleBaru    int `json:"total_article_baru"`
	TotalArticlePending int `json:"total_article_pending"`
}

type ArticleHandlerInterface interface {
	GetArticles() echo.HandlerFunc
	GetArticle() echo.HandlerFunc
	CreateArticle() echo.HandlerFunc
	UpdateArticle() echo.HandlerFunc
	ArticleDashboard() echo.HandlerFunc
	DenyArticle() echo.HandlerFunc
	ApproveArticle() echo.HandlerFunc
}

type ArticleServiceInterface interface {
	GetArticles(name, kategori string, timePublication, limit int) ([]ArticleInfo, error)
	GetArticle(id int) ([]ArticleInfo, error)
	CreateArticle(newData Article) (*Article, error)
	UpdateArticle(newData UpdateArticle, id int) (bool, error)
	ArticleDashboard() (ArticleDashboard, error)
	DenyArticle(id int) (bool, error)
	ApproveArticle(id int) (bool, error)
}

type ArticleDataInterface interface {
	GetAll(name, kategori string, timePublication, limit int) ([]ArticleInfo, error)
	GetByID(id int) ([]ArticleInfo, error)
	Insert(newData Article) (*Article, error)
	Update(newData UpdateArticle, id int) (bool, error)
	ArticleDashboard() (ArticleDashboard, error)
	Deny(id int) (bool, error)
	Approve(id int) (bool, error)
}
