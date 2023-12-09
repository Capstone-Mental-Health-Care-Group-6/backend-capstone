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

	dataDoctor "FinalProject/features/doctor/data"
	handlerDoctor "FinalProject/features/doctor/handler"
	serviceDoctor "FinalProject/features/doctor/service"

	dataArticleCategory "FinalProject/features/article_categories/data"
	handlerArticleCategory "FinalProject/features/article_categories/handler"
	serviceArticleCategory "FinalProject/features/article_categories/service"

	dataWithdraw "FinalProject/features/withdraw/data"
	handlerWithdraw "FinalProject/features/withdraw/handler"
	serviceWithdraw "FinalProject/features/withdraw/service"

	dataPatient "FinalProject/features/patients/data"
	handlerPatient "FinalProject/features/patients/handler"
	servicePatient "FinalProject/features/patients/service"

	dataBundle "FinalProject/features/bundle_counseling/data"
	handlerBundle "FinalProject/features/bundle_counseling/handler"
	serviceBundle "FinalProject/features/bundle_counseling/service"

	dataCounseling "FinalProject/features/counseling_session/data"
	handlerCounseling "FinalProject/features/counseling_session/handler"
	serviceCounseling "FinalProject/features/counseling_session/service"

	dataChat "FinalProject/features/chats/data"
	handlerChat "FinalProject/features/chats/handler"
	serviceChat "FinalProject/features/chats/service"

	dataChatbot "FinalProject/features/chatbot/data"
	handlerChatbot "FinalProject/features/chatbot/handler"
	serviceChatbot "FinalProject/features/chatbot/service"

	dataChatbotCs "FinalProject/features/chatbotcs/data"
	handlerChatbotCs "FinalProject/features/chatbotcs/handler"
	serviceChatbotCs "FinalProject/features/chatbotcs/service"

	dataMessage "FinalProject/features/chat_messages/data"
	handlerMessage "FinalProject/features/chat_messages/handler"
	serviceMessage "FinalProject/features/chat_messages/service"

	"FinalProject/helper"
	"FinalProject/routes"
	"FinalProject/utils/cloudinary"
	"FinalProject/utils/database"
	"FinalProject/utils/midtrans"
	"FinalProject/utils/oauth"
	"FinalProject/utils/openai"
	"FinalProject/utils/websocket"

	// "fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = configs.InitConfig()

	var cld = cloudinary.InitCloud(*config)
	var midtrans = midtrans.InitMidtrans(*config)
	var openai = openai.InitOpenAI(*config)
	db, err := database.InitDB(*config)
	if err != nil {
		e.Logger.Fatal("cannot run database, ", err.Error())
	}

	mongo, err := database.InitMongoDb(*config)
	if err != nil {
		e.Logger.Fatal("cannot run mongo database, ", err.Error())
	}

	database.Migrate(db)
	oauth := oauth.NewOauthGoogleConfig(*config)
	jwtInterface := helper.New(config.Secret, config.RefSecret)

	userModel := dataUser.New(db)
	userServices := serviceUser.New(userModel, jwtInterface, *config)
	userController := handlerUser.NewHandler(userServices, oauth, jwtInterface)

	transaksiModel := dataTransaksi.New(db)
	transaksiServices := serviceTransaksi.New(transaksiModel, cld, midtrans)
	transaksiController := handlerTransaksi.NewTransactionHandler(transaksiServices, jwtInterface)

	articleModel := dataArticle.New(db)
	articleServices := serviceArticle.New(articleModel)
	articleController := handlerArticle.NewHandler(articleServices, jwtInterface)

	articleCategoryModel := dataArticleCategory.New(db)
	articleCategoryServices := serviceArticleCategory.New(articleCategoryModel)
	articleCategoryController := handlerArticleCategory.NewHandler(articleCategoryServices, jwtInterface)

	patientModel := dataPatient.New(db)
	patientServices := servicePatient.NewPatient(patientModel, cld, jwtInterface)
	patientController := handlerPatient.NewHandlerPatient(patientServices, jwtInterface)

	doctorModel := dataDoctor.NewDoctor(db)
	doctorServices := serviceDoctor.NewDoctor(doctorModel, cld, *config)
	doctorController := handlerDoctor.NewHandlerDoctor(doctorServices, jwtInterface)

	withdrawModel := dataWithdraw.New(db)
	withdrawServices := serviceWithdraw.New(withdrawModel)
	withdrawController := handlerWithdraw.New(withdrawServices, jwtInterface)

	bundleModel := dataBundle.New(db)
	bundleServices := serviceBundle.New(bundleModel, cld)
	bundleController := handlerBundle.New(bundleServices, jwtInterface)

	counselingModel := dataCounseling.New(db)
	counselingServices := serviceCounseling.New(counselingModel, cld)
	counselingController := handlerCounseling.New(counselingServices, jwtInterface)

	chatData := dataChat.New(db)
	chatServices := serviceChat.New(chatData, websocket.NewServer())
	chatController := handlerChat.New(chatServices)

	messageModel := dataMessage.New(db)
	messageServices := serviceMessage.New(messageModel)
	messageController := handlerMessage.New(messageServices)

	chatbotModel := dataChatbot.New(mongo)
	chatbotService := serviceChatbot.New(chatbotModel, openai)
	chatbotController := handlerChatbot.New(chatbotService, jwtInterface)

	chatbotCsModel := dataChatbotCs.New(map[string]string{})
	chatbotCsService := serviceChatbotCs.New(chatbotCsModel)
	chatbotCsHandler := handlerChatbotCs.New(chatbotCsService, jwtInterface)

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
	routes.RoutePatient(e, patientController, *config)
	routes.RouteDoctor(e, doctorController, *config)
	routes.RouteWithdraw(e, withdrawController, *config)
	routes.RouteBundle(e, bundleController, *config)
	routes.RouteCounseling(e, counselingController, *config)

	routes.RouteChat(e, chatController, *config)
	routes.RouteMessage(e, messageController, *config)
	routes.RouteChatBot(e, chatbotController, *config)
	routes.RouteChatBotCS(e, chatbotCsHandler, *config)

	e.Logger.Debug(db)

	//DEVMODE

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
