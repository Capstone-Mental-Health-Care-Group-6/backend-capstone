package main

import (
	"FinalProject/configs"
	dataArticle "FinalProject/features/articles/data"
	handlerArticle "FinalProject/features/articles/handler"
	serviceArticle "FinalProject/features/articles/service"
	"FinalProject/utils/cloudinary"
	"FinalProject/utils/midtrans"

	dataTransaksi "FinalProject/features/transaction/data"
	handlerTransaksi "FinalProject/features/transaction/handler"
	serviceTransaksi "FinalProject/features/transaction/service"

	dataUser "FinalProject/features/users/data"
	handlerUser "FinalProject/features/users/handler"
	serviceUser "FinalProject/features/users/service"

	dataDoctor "FinalProject/features/doctor/data"
	handlerDoctor "FinalProject/features/doctor/handler"
	serviceDoctor "FinalProject/features/doctor/service"

	dataArticleCategory "FinalProject/features/article_categories/data"
	handlerArticleCategory "FinalProject/features/article_categories/handler"
	serviceArticleCategory "FinalProject/features/article_categories/service"

	dataWithdraw "FinalProject/features/withdraw/data"
	handlerWithdraw "FinalProject/features/withdraw/handler"
	serviceWithdraw "FinalProject/features/withdraw/service"

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

	midtrans := midtrans.InitMidtrans(*config)

	userModel := dataUser.New(db)
	jwtInterface := helper.New(config.Secret, config.RefSecret)
	userServices := serviceUser.New(userModel, jwtInterface)
	userController := handlerUser.NewHandler(userServices)

	articleModel := dataArticle.New(db)
	articleServices := serviceArticle.New(articleModel)
	articleController := handlerArticle.NewHandler(articleServices)

	transaksiModel := dataTransaksi.New(db)
	transaksiServices := serviceTransaksi.New(transaksiModel, cld, midtrans)
	transaksiController := handlerTransaksi.NewTransactionHandler(transaksiServices)

	articleCategoryModel := dataArticleCategory.New(db)
	articleCategoryServices := serviceArticleCategory.New(articleCategoryModel)
	articleCategoryController := handlerArticleCategory.NewHandler(articleCategoryServices)

	// patientModel := dataPatient.NewPatient(db)
	// patientServices := servicePatient.NewPatient(patientModel, cld)
	// patientController := handlerPatient.NewHandlerPatient(patientServices)

	doctorModel := dataDoctor.NewDoctor(db)
	doctorServices := serviceDoctor.NewDoctor(doctorModel, cld)
	doctorController := handlerDoctor.NewHandlerDoctor(doctorServices)

	withdrawModel := dataWithdraw.New(db)
	withdrawServices := serviceWithdraw.New(withdrawModel)
	withdrawController := handlerWithdraw.New(withdrawServices, jwtInterface)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.RouteUser(e, userController, *config)
	routes.RouteArticle(e, articleController, *config)
	routes.RouteArticleCategory(e, articleCategoryController, *config)
	// routes.RoutePatient(e, patientController, *config)
	routes.RouteTransaction(e, transaksiController, *config)
	routes.RouteDoctor(e, doctorController, *config)
	routes.RouteWithdraw(e, withdrawController, *config)

	config.ServerPort = 8080

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
