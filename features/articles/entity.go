package articles

import "github.com/labstack/echo/v4"

type Article struct {
	ID         uint   `json:"id"`
	CategoryID uint   `json:"category_id"`
	UserID     uint   `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Thumbnail  string `json:"thumbnail"`
	Status     string `json:"status"`
	Slug       string `json:"slug"`
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
	Title     string `json:"title"`
	Content   string `json:"content"`
	Thumbnail string `json:"thumbnail"`
	Slug      string `json:"slug"`
}

type ArticleHandlerInterface interface {
	GetArticles() echo.HandlerFunc
	GetArticle() echo.HandlerFunc
	CreateArticle() echo.HandlerFunc
	UpdateArticle() echo.HandlerFunc
	DeleteArticle() echo.HandlerFunc
}

type ArticleServiceInterface interface {
	GetArticles() ([]ArticleInfo, error)
	GetArticle(id int) ([]ArticleInfo, error)
	CreateArticle(newData Article) (*Article, error)
	UpdateArticle(newData UpdateArticle, id int) (bool, error)
	DeleteArticle(id int) (bool, error)
}

type ArticleDataInterface interface {
	GetAll() ([]ArticleInfo, error)
	GetByID(id int) ([]ArticleInfo, error)
	Insert(newData Article) (*Article, error)
	Update(newData UpdateArticle, id int) (bool, error)
	Delete(id int) (bool, error)
}
