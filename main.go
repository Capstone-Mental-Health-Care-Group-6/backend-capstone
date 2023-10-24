package main

import (
	"FinalProject/configs"
	dataArticle "FinalProject/features/articles/data"
	handlerArticle "FinalProject/features/articles/handler"
	serviceArticle "FinalProject/features/articles/service"

	dataUser "FinalProject/features/users/data"
	handlerUser "FinalProject/features/users/handler"
	serviceUser "FinalProject/features/users/service"

	dataArticleCategory "FinalProject/features/article_categories/data"
	handlerArticleCategory "FinalProject/features/article_categories/handler"
	serviceArticleCategory "FinalProject/features/article_categories/service"

	"FinalProject/helper"
	"FinalProject/routes"
	"FinalProject/utils/database"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = configs.InitConfig()

	db := database.InitDB(*config)
	database.Migrate(db)

	userModel := dataUser.New(db)
	jwtInterface := helper.New(config.Secret, config.RefSecret)
	userServices := serviceUser.New(userModel, jwtInterface)
	userController := handlerUser.NewHandler(userServices)

	articleModel := dataArticle.New(db)
	articleServices := serviceArticle.New(articleModel)
	articleController := handlerArticle.NewHandler(articleServices)

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
	routes.RouteArticle(e, articleController, *config)
	routes.RouteArticleCategory(e, articleCategoryController, *config)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
