package routes

import (
	"FinalProject/configs"
	articlecategories "FinalProject/features/article_categories"
	"FinalProject/features/articles"
	"FinalProject/features/doctor"
	"FinalProject/features/users"
	"FinalProject/features/withdraw"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uh users.UserHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/register", uh.Register())
	e.POST("/login", uh.Login())
}

func RouteArticle(e *echo.Echo, ah articles.ArticleHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/articles", ah.GetArticles(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/articles/:id", ah.GetArticle(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/articles", ah.CreateArticle(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/articles/:id", ah.UpdateArticle(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/articles/:id", ah.DeleteArticle(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteArticleCategory(e *echo.Echo, ach articlecategories.ArticleCategoryHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/article/categories", ach.GetArticleCategories(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/article/categories/:id", ach.GetArticleCategory(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/article/categories", ach.CreateArticleCategory(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/article/categories/:id", ach.UpdateArticleCategory(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/article/categories/:id", ach.DeleteArticleCategory(), echojwt.JWT([]byte(cfg.Secret)))
}

// func RoutePatient(e *echo.Echo, ph users.PatientHandlerInterface, cfg configs.ProgrammingConfig) {
// 	e.GET("/patients", ph.GetPatients())
// 	e.GET("/patients/:id", ph.GetPatient())
// 	e.POST("/patients/register", ph.CreatePatient())
// }

func RouteDoctor(e *echo.Echo, ph doctor.DoctorHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/doctor", ph.GetDoctors())
	e.GET("/doctor/:id", ph.GetDoctor())
	e.POST("/doctor/register", ph.CreateDoctor())
}

func RouteWithdraw(e *echo.Echo, wh withdraw.WithdrawHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/withdraw", wh.GetAllWithdraw(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/withdraw", wh.CreateWithdraw(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/withdraw/:id", wh.GetWithdraw(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/withdraw/:id/status", wh.UpdateStatus(), echojwt.JWT([]byte(cfg.Secret)))
}
