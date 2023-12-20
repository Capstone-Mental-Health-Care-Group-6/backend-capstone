package database

import (
	DataArticleCategory "FinalProject/features/article_categories/data"
	DataArticle "FinalProject/features/articles/data"
	DataBundle "FinalProject/features/bundle_counseling/data"
	DataCounselingDuration "FinalProject/features/counseling_durations/data"
	DataCounselingMethod "FinalProject/features/counseling_methods/data"
	DataCounseling "FinalProject/features/counseling_session/data"
	DataCounselingTopic "FinalProject/features/counseling_topics/data"
	DataDoctor "FinalProject/features/doctor/data"
	DataPatient "FinalProject/features/patients/data"
	DataTransaction "FinalProject/features/transaction/data"
	DataUser "FinalProject/features/users/data"
	DataWithdraw "FinalProject/features/withdraw/data"

	message "FinalProject/features/chat_messages"
	chat "FinalProject/features/chats"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	// db.Migrator().DropTable(DataArticleCategory.ArticleCategory{})
	// db.Migrator().DropTable(DataCounselingMethod.CounselingMethod{})
	// db.Migrator().DropTable(DataCounselingDuration.CounselingDuration{})
	// db.Migrator().DropTable(DataCounselingTopic.CounselingTopic{})

	db.AutoMigrate(DataUser.User{})
	db.AutoMigrate(DataArticleCategory.ArticleCategory{})
	db.AutoMigrate(DataArticle.Article{})
	db.AutoMigrate(DataDoctor.Doctor{})
	db.AutoMigrate(DataTransaction.Transaction{})
	db.AutoMigrate(DataUser.UserResetPass{})
	db.AutoMigrate(DataPatient.PatientAccount{})
	db.AutoMigrate(DataBundle.BundleCounseling{})
	db.AutoMigrate(DataDoctor.DoctorExpertiseRelation{})
	db.AutoMigrate(DataDoctor.DoctorWorkadays{})
	db.AutoMigrate(DataDoctor.DoctorExperience{})
	db.AutoMigrate(DataDoctor.DoctorEducation{})
	db.AutoMigrate(DataDoctor.DoctorRating{})
	db.AutoMigrate(DataCounseling.CounselingSession{})
	db.AutoMigrate(DataCounselingMethod.CounselingMethod{})
	db.AutoMigrate(DataCounselingDuration.CounselingDuration{})
	db.AutoMigrate(DataCounselingTopic.CounselingTopic{})
	db.AutoMigrate(chat.Chat{})
	db.AutoMigrate(message.Message{})
	db.AutoMigrate(DataWithdraw.Withdraw{})
}
