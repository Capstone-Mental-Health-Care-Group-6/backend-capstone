package routes

import (
	"FinalProject/configs"
	articlecategories "FinalProject/features/article_categories"
	"FinalProject/features/articles"
	bundlecounseling "FinalProject/features/bundle_counseling"
	"FinalProject/features/chatbot"
	"FinalProject/features/chatbotcs"
	"FinalProject/features/doctor"
	"FinalProject/features/patients"
	transaction "FinalProject/features/transaction"
	"FinalProject/features/users"
	"FinalProject/features/withdraw"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uh users.UserHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/register", uh.Register())
	e.POST("/login", uh.Login())
	e.GET("/login/google", uh.LoginGoogle())
	e.GET("/login/google/callback", uh.CallbackGoogle())
}

func RouteTransaction(e *echo.Echo, th transaction.TransactionHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/transaksi", th.CreateTransaction())
	e.POST("/transaksi/notif", th.NotifTransaction())
	e.GET("/transaksi/:id", th.GetTransaction())
	e.GET("/transaksi", th.GetTransactions())
	e.DELETE("/transaksi/:id", th.DeleteTransaction())
	e.GET("/transaksi/check/:id", th.GetTransactionByMidtransID())
	e.PUT("/transaksi/:id", th.UpdateTransaction())
	// e.POST("/transaksi/manual", th.CreateManualTransaction())
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

func RoutePatient(e *echo.Echo, ph patients.PatientHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/patient", ph.GetPatients(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/patient/account/:id", ph.GetPatient(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/patient/register", ph.CreatePatient())
	e.POST("/patient/login", ph.LoginPatient())
	e.PUT("/patient/account/update", ph.UpdatePatient(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/patient/account/update/password", ph.UpdatePassword(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteBundle(e *echo.Echo, ph bundlecounseling.BundleCounselingHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/bundle", ph.GetAllBundle(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/bundle/:id", ph.GetBundle(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/bundle", ph.CreateBundle(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/bundle/:id", ph.UpdateBundle(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/bundle/:id", ph.DeleteBundle(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteChatBot(e *echo.Echo, ch chatbot.ChatbotHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/chatbot", ch.GetAllChatBot(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/chatbot", ch.CreateChatBot(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteChatBotCS(e *echo.Echo, ch chatbotcs.ChatbotCsHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/chatbotcs", ch.ChatBotCs())
	e.POST("/chatbotcs", ch.CreateMessage())
}
