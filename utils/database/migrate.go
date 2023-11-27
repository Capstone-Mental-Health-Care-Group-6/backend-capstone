package database

import (
	DataArticleCategory "FinalProject/features/article_categories/data"
	DataArticle "FinalProject/features/articles/data"
	DataDoctor "FinalProject/features/doctor/data"
	DataTransaction "FinalProject/features/transaction/data"
	DataUser "FinalProject/features/users/data"
	"FinalProject/utils/database/migration"
  DataPatient "FinalProject/features/patients/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	migrator := migration.NewMySqlMigrator(db)
	// migrator.DropTable([]migration.Table{
	// 	DataUser.User{},
	// 	DataArticleCategory.ArticleCategory{},
	// 	DataArticle.Article{},
	// 	DataTransaction.Transaction{},
  //  DataPatient.PatientAccount{},
	// }...)
  
	migrator.CreateTable([]migration.Table{
		DataUser.User{},
		DataArticleCategory.ArticleCategory{},
		DataArticle.Article{},
		DataDoctor.Doctor{},
		DataTransaction.Transaction{},
    DataPatient.PatientAccount{},
	}...)
}