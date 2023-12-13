package articlecategories

import (
	"time"

	"github.com/labstack/echo/v4"
)

type ArticleCategory struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateArticleCategory struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ArticleCategoryHandlerInterface interface {
	GetArticleCategories() echo.HandlerFunc
	GetArticleCategory() echo.HandlerFunc
	CreateArticleCategory() echo.HandlerFunc
	UpdateArticleCategory() echo.HandlerFunc
	DeleteArticleCategory() echo.HandlerFunc
}

type ArticleCategoryServiceInterface interface {
	GetArticleCategories() ([]ArticleCategory, error)
	GetArticleCategory(id int) ([]ArticleCategory, error)
	CreateArticleCategory(newData ArticleCategory) (*ArticleCategory, error)
	UpdateArticleCategory(newData UpdateArticleCategory, id int) (bool, error)
	DeleteArticleCategory(id int) (bool, error)
}

type ArticleCategoryDataInterface interface {
	GetAll() ([]ArticleCategory, error)
	GetByID(id int) ([]ArticleCategory, error)
	Insert(newData ArticleCategory) (*ArticleCategory, error)
	Update(newData UpdateArticleCategory, id int) (bool, error)
	Delete(id int) (bool, error)
}
