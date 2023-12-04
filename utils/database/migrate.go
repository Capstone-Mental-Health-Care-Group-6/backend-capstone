package database

import (
	DataArticleCategory "FinalProject/features/article_categories/data"
	DataArticle "FinalProject/features/articles/data"
	DataBundle "FinalProject/features/bundle_counseling/data"
	DataDoctor "FinalProject/features/doctor/data"
	DataPatient "FinalProject/features/patients/data"
	DataTransaction "FinalProject/features/transaction/data"
	DataUser "FinalProject/features/users/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(DataUser.User{})
	db.AutoMigrate(DataArticleCategory.ArticleCategory{})
	db.AutoMigrate(DataArticle.Article{})
	db.AutoMigrate(DataDoctor.Doctor{})
	db.AutoMigrate(DataTransaction.Transaction{})
	db.AutoMigrate(DataUser.UserResetPass{})
	db.AutoMigrate(DataPatient.PatientAccount{})
	db.AutoMigrate(DataBundle.BundleCounseling{})
	db.AutoMigrate(DataDoctor.DoctorExpertiseRelation{})
	db.AutoMigrate(DataDoctor.DoctorRating{})
	db.AutoMigrate(DataDoctor.DoctorWorkadays{})
}
