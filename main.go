<<<<<<< HEAD
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
	"FinalProject/utils/cloudinary"
	"FinalProject/utils/database"
	"FinalProject/utils/midtrans"

	// "fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = configs.InitConfig()

	db := database.InitDB(*config)
	database.Migrate(db)

	midtrans := midtrans.InitMidtrans(*config)
	cld := cloudinary.InitCloud(*config)

	userModel := dataUser.New(db)
	jwtInterface := helper.New(config.Secret, config.RefSecret)
	userServices := serviceUser.New(userModel, jwtInterface)
	userController := handlerUser.NewHandler(userServices)

	transaksiModel := dataTransaksi.New(db)
	transaksiServices := serviceTransaksi.New(transaksiModel, cld, midtrans)
	transaksiController := handlerTransaksi.NewTransactionHandler(transaksiServices)

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
	routes.RouteTransaction(e, transaksiController, *config)
	routes.RouteArticle(e, articleController, *config)
	routes.RouteArticleCategory(e, articleCategoryController, *config)

	config.ServerPort = 8080

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
=======
package main

import (
	"FinalProject/configs"
	dataArticle "FinalProject/features/articles/data"
	handlerArticle "FinalProject/features/articles/handler"
	serviceArticle "FinalProject/features/articles/service"
	"FinalProject/utils/cloudinary"

	dataUser "FinalProject/features/users/data"
	handlerUser "FinalProject/features/users/handler"
	serviceUser "FinalProject/features/users/service"

	dataPatient "FinalProject/features/users/data"
	handlerPatient "FinalProject/features/users/handler"
	servicePatient "FinalProject/features/users/service"

	dataDoctor "FinalProject/features/doctor/data"
	handlerDoctor "FinalProject/features/doctor/handler"
	serviceDoctor "FinalProject/features/doctor/service"

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
	config := configs.InitConfig()

	var db = database.InitDB(*config)
	database.Migrate(db)

	var cld = cloudinary.InitCloud(*config)

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

	patientModel := dataPatient.NewPatient(db)
	patientServices := servicePatient.NewPatient(patientModel, cld)
	patientController := handlerPatient.NewHandlerPatient(patientServices)

	doctorModel := dataDoctor.NewDoctor(db)
	doctorServices := serviceDoctor.NewDoctor(doctorModel, cld)
	doctorController := handlerDoctor.NewHandlerDoctor(doctorServices)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.RouteUser(e, userController, *config)
	routes.RouteArticle(e, articleController, *config)
	routes.RouteArticleCategory(e, articleCategoryController, *config)
	routes.RoutePatient(e, patientController, *config)
	routes.RouteDoctor(e, doctorController, *config)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
>>>>>>> b13016e4c03d4e417f83b95a33ada4d1c328b308
