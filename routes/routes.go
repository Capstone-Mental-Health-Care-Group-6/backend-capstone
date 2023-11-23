package routes

import (
	"FinalProject/configs"
	articlecategories "FinalProject/features/article_categories"
	"FinalProject/features/articles"
	"FinalProject/features/patients"
	transaction "FinalProject/features/transaction"
	"FinalProject/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uh users.UserHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/register", uh.Register())
	e.POST("/login", uh.Login())
}

func RouteTransaction(e *echo.Echo, th transaction.TransactionHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/transaksi", th.CreateTransaction())
	e.POST("/transaksi/notif", th.NotifTransaction())
	e.GET("/transaksi/:id", th.GetTransaction())
	e.GET("/transaksi", th.GetTransactions())
	e.DELETE("/transaksi/:id", th.DeleteTransaction())
	e.GET("/transaksi/check/:id", th.GetTransactionByMidtransID())
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

func RoutePatientAccount(e *echo.Echo, Patient patients.PatientHandlerInterface, cfg configs.ProgrammingConfig) {
	patient := e.Group("/patient")
	patient.POST("/register", Patient.CreatePatient())
	patient.POST("/login", Patient.LoginPatient())
	patient.GET("/account", Patient.GetPatients(), echojwt.JWT([]byte(cfg.Secret)))
	patient.PUT("/account/update", Patient.UpdatePatient(), echojwt.JWT([]byte(cfg.Secret)))
	patient.PUT("/account/update/password", Patient.UpdatePassword(), echojwt.JWT([]byte(cfg.Secret)))
	//patient.DELETE("/account/delete", Patient.DeletePatient(), echojwt.JWT([]byte(cfg.Secret)))
}
