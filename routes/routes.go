package routes

import (
	"FinalProject/configs"
	articlecategories "FinalProject/features/article_categories"
	"FinalProject/features/articles"
	bundlecounseling "FinalProject/features/bundle_counseling"
	. "FinalProject/features/chat_messages"
	"FinalProject/features/chatbot"
	"FinalProject/features/chatbotcs"
	. "FinalProject/features/chats"
	counselingdurations "FinalProject/features/counseling_durations"
	counselingmethod "FinalProject/features/counseling_methods"
	counselingsession "FinalProject/features/counseling_session"
	counselingtopics "FinalProject/features/counseling_topics"
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
	e.POST("/forget-password", uh.ForgetPasswordWeb())
	e.POST("/forget-password/verify", uh.ForgetPasswordVerify())
	e.POST("/reset-password", uh.ResetPassword())
	e.POST("/refresh-token", uh.RefreshToken(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/admin/update", uh.UpdateProfile(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteTransaction(e *echo.Echo, th transaction.TransactionHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/transaksi", th.CreateTransaction(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/transaksi/notif", th.NotifTransaction())
	e.GET("/transaksi/:id", th.GetTransaction(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/transaksi", th.GetTransactions(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/transaksi/:id", th.DeleteTransaction(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/transaksi/check/:id", th.GetTransactionByMidtransID(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/transaksi/patient/:id", th.GetTransactionByPatientID(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/transaksi/:id", th.UpdateTransaction(), echojwt.JWT([]byte(cfg.Secret)))
	// e.POST("/transaksi/manual", th.CreateManualTransaction())
}

func RouteArticle(e *echo.Echo, ah articles.ArticleHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/articles", ah.GetArticles())
	e.GET("/articles/:id", ah.GetArticle())
	e.GET("/doctor/articles", ah.GetArticlesByDoctorID(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/articles", ah.CreateArticle(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/articles/:id", ah.UpdateArticle(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/articles/:id/reject", ah.RejectArticle(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/articles/:id/publish", ah.PublishArticle(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/articles/dashboard", ah.ArticleDashboard(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteArticleCategory(e *echo.Echo, ach articlecategories.ArticleCategoryHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/article/categories", ach.GetArticleCategories())
	e.GET("/article/categories/:id", ach.GetArticleCategory())
	e.POST("/article/categories", ach.CreateArticleCategory(), echojwt.JWT([]byte(cfg.Secret)))
	e.PATCH("/article/categories/:id", ach.UpdateArticleCategory(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/article/categories/:id", ach.DeleteArticleCategory(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteDoctor(e *echo.Echo, ph doctor.DoctorHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/doctor", ph.GetDoctors(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/doctor/:id", ph.GetDoctor(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/doctor/dashboard/:id", ph.DoctorDashboard(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/doctor/patientlist/:id", ph.DoctorDashboardPatient(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/doctor/user/:id", ph.GetDoctorByUserId(), echojwt.JWT([]byte(cfg.Secret)))

	e.POST("/doctor/register", ph.CreateDoctor(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/doctor/insert/:type", ph.InsertDataDoctor(), echojwt.JWT([]byte(cfg.Secret)))

	e.PUT("/doctor/datapokok/:id", ph.UpdateDoctorDatapokok(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/doctor/workday/:id", ph.UpdateDoctorWorkdays(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/doctor/education/:id", ph.UpdateDoctorEducation(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/doctor/experience/:id", ph.UpdateDoctorExperience(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/doctor/rating/:id", ph.UpdateDoctorRating(), echojwt.JWT([]byte(cfg.Secret)))

	e.DELETE("/doctor/:id", ph.DeleteDoctor(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/doctor/:type/:id", ph.DeleteDoctorData(), echojwt.JWT([]byte(cfg.Secret)))

	// ROUTE UNTUK ADMIN
	e.GET("/doctor/dashboard", ph.DoctorDashboardAdmin(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/doctor/:user_id/approve", ph.ApproveDoctor(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/doctor/:user_id/deny", ph.DenyDoctor(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteWithdraw(e *echo.Echo, wh withdraw.WithdrawHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/withdraw", wh.GetAllWithdraw(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/doctor/withdraw", wh.GetAllWithdrawDokter(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/withdraw", wh.CreateWithdraw(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/withdraw/:id", wh.GetWithdraw(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/withdraw/:id/status", wh.UpdateStatus(), echojwt.JWT([]byte(cfg.Secret)))

}

func RoutePatient(e *echo.Echo, ph patients.PatientHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/patient/register", ph.CreatePatient())
	e.POST("/patient/login", ph.LoginPatient())

	e.GET("/patient", ph.GetPatients(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/patient/account/:id", ph.GetPatient(), echojwt.JWT([]byte(cfg.Secret)))

	// PATIENT
	e.PUT("/patient/account/update", ph.UpdatePatient(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/patient/account/update/password", ph.UpdatePassword(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/patient/inactivate", ph.InactivateAccount(), echojwt.JWT([]byte(cfg.Secret)))

	// ADMIN
	e.PUT("/patient/update/:id/status", ph.UpdateStatus(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/patient/dashboard", ph.PatientDashboard(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteBundle(e *echo.Echo, ph bundlecounseling.BundleCounselingHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/bundle", ph.GetAllBundle(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/mobile/bundle", ph.GetAllBundleFilter())
	e.GET("/bundle/:id", ph.GetBundle(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/bundle", ph.CreateBundle(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/bundle/:id", ph.UpdateBundle(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/bundle/:id", ph.DeleteBundle(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteCounseling(e *echo.Echo, ph counselingsession.CounselingSessionHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/counseling", ph.GetAllCounseling(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/counseling/user/:id", ph.GetCounselingByUserID(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/counseling/:id", ph.GetCounseling(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/counseling", ph.CreateCounseling(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/counseling/:id", ph.UpdateCounseling(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/counseling/:id", ph.DeleteCounseling(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/counseling/:id/approve", ph.ApprovePatient(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/counseling/:id/reject", ph.RejectPatient(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteCounselingMethod(e *echo.Echo, mh counselingmethod.CounselingMethodHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/counseling/methods", mh.GetCounselingMethods())
	e.GET("/counseling/methods/:id", mh.GetCounselingMethod())
}

func RouteCounselingDuration(e *echo.Echo, mh counselingdurations.CounselingDurationHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/counseling/durations", mh.GetCounselingDurations())
	e.GET("/counseling/durations/:id", mh.GetCounselingDuration())
}

func RouteCounselingTopic(e *echo.Echo, mh counselingtopics.CounselingTopicHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/counseling/topics", mh.GetCounselingTopics())
	e.GET("/counseling/topics/:id", mh.GetCounselingTopic())
}

func RouteChat(e *echo.Echo, h ChatHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/api/socket/:role/:id", h.Establish())
	group := e.Group("/api/chats")
	group.GET("", h.Index(), echojwt.JWT([]byte(cfg.Secret)))
	group.POST("", h.Store(), echojwt.JWT([]byte(cfg.Secret)))
	group.PUT("/:id", h.Edit(), echojwt.JWT([]byte(cfg.Secret)))
	group.DELETE("/:id", h.Destroy(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteMessage(e *echo.Echo, h MessageHandlerInterface, cfg configs.ProgrammingConfig) {
	group := e.Group("/api/chats/:id/messages")
	group.GET("", h.Index(), echojwt.JWT([]byte(cfg.Secret)))
	group.GET("/:message", h.Observe(), echojwt.JWT([]byte(cfg.Secret)))
	group.POST("", h.Store(), echojwt.JWT([]byte(cfg.Secret)))
	group.PUT("/:message", h.Edit(), echojwt.JWT([]byte(cfg.Secret)))
	group.DELETE("/:message", h.Destroy(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteChatBot(e *echo.Echo, ch chatbot.ChatbotHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/chatbot", ch.GetAllChatBot(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/chatbot", ch.CreateChatBot(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteChatBotCS(e *echo.Echo, ch chatbotcs.ChatbotCsHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET("/chatbotcs", ch.ChatBotCs())
	e.POST("/chatbotcs", ch.CreateMessage())
}
