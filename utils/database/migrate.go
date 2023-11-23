package database

import (
	DataArticleCategory "FinalProject/features/article_categories/data"
	DataArticle "FinalProject/features/articles/data"
	dataPatient "FinalProject/features/patients/data"
	DataTransaction "FinalProject/features/transaction/data"
	DataUser "FinalProject/features/users/data"
	"FinalProject/utils/database/migration"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	migrator := migration.NewMySqlMigrator(db)
	//migrator.DropTable([]migration.Table{
	//	DataUser.User{},
	//	DataArticleCategory.ArticleCategory{},
	//	DataArticle.Article{},
	//	DataTransaction.Transaction{},
	//	dataPatient.PatientAccount{},
	//}...)
	migrator.CreateTable([]migration.Table{
		DataUser.User{},
		DataArticleCategory.ArticleCategory{},
		DataArticle.Article{},
		DataTransaction.Transaction{},
		dataPatient.PatientAccount{},
	}...)
}
