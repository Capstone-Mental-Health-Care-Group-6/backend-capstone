package main

import (
	"FinalProject/configs"

	dataArticle "FinalProject/features/articles/data"
	handlerArticle "FinalProject/features/articles/handler"
	serviceArticle "FinalProject/features/articles/service"

	"FinalProject/helper/email"
	"FinalProject/helper/enkrip"
	"FinalProject/helper/slug"

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

	dataCounselingMethod "FinalProject/features/counseling_methods/data"
	handlerCounselingMethod "FinalProject/features/counseling_methods/handler"
	serviceCounselingMethod "FinalProject/features/counseling_methods/service"

	dataCounselingDuration "FinalProject/features/counseling_durations/data"
	handlerCounselingDuration "FinalProject/features/counseling_durations/handler"
	serviceCounselingDuration "FinalProject/features/counseling_durations/service"

	dataCounselingTopic "FinalProject/features/counseling_topics/data"
	handlerCounselingTopic "FinalProject/features/counseling_topics/handler"
	serviceCounselingTopic "FinalProject/features/counseling_topics/service"

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

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = configs.InitConfig()

	var cld = cloudinary.InitCloud(*config)
	var midtrans = midtrans.InitMidtrans(*config)

	var enkrip = enkrip.New()
	var slug = slug.New()
	var email = email.New(*config)
	var openai = openai.InitOpenAI(*config)
	var meet = helper.NewMeet()
	db, err := database.InitDB(*config)
	if err != nil {
		e.Logger.Fatal("cannot run database, ", err.Error())
	}

	mongo, err := database.InitMongoDb(*config)
	if err != nil {
		e.Logger.Fatal("cannot run mongo database, ", err.Error())
	}

	database.Migrate(db)

	// HANYA 1x SAJA
	// for _, seed := range seeds.All() {
	// 	if err := seed.Run(db); err != nil {
	// 		fmt.Printf("Running seed '%s', failed with error: %s", seed.Name, err)
	// 	}
	// }

	oauth := oauth.NewOauthGoogleConfig(*config)
	jwtInterface := helper.New(config.Secret, config.RefSecret)

	userModel := dataUser.New(db)
	userServices := serviceUser.New(userModel, jwtInterface, email, enkrip)
	userController := handlerUser.NewHandler(userServices, oauth, jwtInterface)

	transaksiModel := dataTransaksi.New(db)
	transaksiServices := serviceTransaksi.New(transaksiModel, cld, midtrans)
	transaksiController := handlerTransaksi.NewTransactionHandler(transaksiServices, jwtInterface)

	articleModel := dataArticle.New(db)
	articleServices := serviceArticle.New(articleModel, slug, cld)
	articleController := handlerArticle.NewHandler(articleServices, jwtInterface)

	articleCategoryModel := dataArticleCategory.New(db)
	articleCategoryServices := serviceArticleCategory.New(articleCategoryModel, slug)
	articleCategoryController := handlerArticleCategory.NewHandler(articleCategoryServices, jwtInterface)

	patientModel := dataPatient.New(db)
	patientServices := servicePatient.NewPatient(patientModel, cld, jwtInterface, enkrip)
	patientController := handlerPatient.NewHandlerPatient(patientServices, jwtInterface)

	doctorModel := dataDoctor.NewDoctor(db)
	doctorServices := serviceDoctor.NewDoctor(doctorModel, cld, email, meet)
	doctorController := handlerDoctor.NewHandlerDoctor(doctorServices, jwtInterface)

	withdrawModel := dataWithdraw.New(db)
	withdrawServices := serviceWithdraw.New(withdrawModel)
	withdrawController := handlerWithdraw.New(withdrawServices, jwtInterface)

	bundleModel := dataBundle.New(db)
	bundleServices := serviceBundle.New(bundleModel, cld)
	bundleController := handlerBundle.New(bundleServices, jwtInterface)

	socket := websocket.NewServer()

	counselingModel := dataCounseling.New(db)
	counselingServices := serviceCounseling.New(counselingModel)
	counselingController := handlerCounseling.New(counselingServices, jwtInterface)

	counselingMethodModel := dataCounselingMethod.New(db)
	counselingMethodServices := serviceCounselingMethod.New(counselingMethodModel)
	counselingMethodController := handlerCounselingMethod.NewHandler(counselingMethodServices)

	counselingDurationModel := dataCounselingDuration.New(db)
	counselingDurationServices := serviceCounselingDuration.New(counselingDurationModel)
	counselingDurationController := handlerCounselingDuration.NewHandler(counselingDurationServices)

	counselingTopicModel := dataCounselingTopic.New(db)
	counselingTopicServices := serviceCounselingTopic.New(counselingTopicModel)
	counselingTopicController := handlerCounselingTopic.NewHandler(counselingTopicServices)

	chatData := dataChat.New(db)
	chatServices := serviceChat.New(chatData, patientModel, userModel, socket, jwtInterface)
	chatController := handlerChat.New(chatServices)

	messageModel := dataMessage.New(db)
	messageServices := serviceMessage.New(messageModel, jwtInterface)
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
	routes.RouteCounselingMethod(e, counselingMethodController, *config)
	routes.RouteCounselingDuration(e, counselingDurationController, *config)
	routes.RouteCounselingTopic(e, counselingTopicController, *config)
	routes.RouteChat(e, chatController, *config)
	routes.RouteMessage(e, messageController, *config)
	routes.RouteChatBot(e, chatbotController, *config)
	routes.RouteChatBotCS(e, chatbotCsHandler, *config)

	e.Logger.Debug(db)

	//DEVMODE
	//test

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
