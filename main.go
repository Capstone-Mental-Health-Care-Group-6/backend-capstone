package main

import (
	"FinalProject/configs"
	dataArticle "FinalProject/features/articles/data"
	handlerArticle "FinalProject/features/articles/handler"
	serviceArticle "FinalProject/features/articles/service"
	"fmt"

	dataTransaksi "FinalProject/features/transaction/data"
	handlerTransaksi "FinalProject/features/transaction/handler"
	serviceTransaksi "FinalProject/features/transaction/service"

	dataUser "FinalProject/features/users/data"
	handlerUser "FinalProject/features/users/handler"
	serviceUser "FinalProject/features/users/service"

	dataArticleCategory "FinalProject/features/article_categories/data"
	handlerArticleCategory "FinalProject/features/article_categories/handler"
	serviceArticleCategory "FinalProject/features/article_categories/service"

	"FinalProject/helper"
	"FinalProject/routes"
	"FinalProject/utils/database"
	"FinalProject/utils/midtrans"
	"FinalProject/utils/oauth"

	// "fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = configs.InitConfig()

	db := database.InitDB(*config)
	database.Migrate(db)

	oauth := oauth.NewOauthGoogleConfig(*config)

	midtrans := midtrans.InitMidtrans(*config)
	jwtInterface := helper.New(config.Secret, config.RefSecret)

	userModel := dataUser.New(db)
	userServices := serviceUser.New(userModel, jwtInterface)
	userController := handlerUser.NewHandler(userServices, oauth)

	transaksiModel := dataTransaksi.New(db)
	transaksiServices := serviceTransaksi.New(transaksiModel, midtrans)
	transaksiController := handlerTransaksi.NewTransactionHandler(transaksiServices)

	articleModel := dataArticle.New(db)
	articleServices := serviceArticle.New(articleModel)
	articleController := handlerArticle.NewHandler(articleServices, jwtInterface)

	articleCategoryModel := dataArticleCategory.New(db)
	articleCategoryServices := serviceArticleCategory.New(articleCategoryModel)
	articleCategoryController := handlerArticleCategory.NewHandler(articleCategoryServices)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.RouteUser(e, userController, *config)
	routes.RouteTransaction(e, transaksiController, *config)
	routes.RouteArticle(e, articleController, *config)
	routes.RouteArticleCategory(e, articleCategoryController, *config)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
