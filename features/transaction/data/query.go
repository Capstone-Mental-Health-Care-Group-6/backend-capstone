package data

import (
	"FinalProject/features/transaction"

	"gorm.io/gorm"
)

type TransactionData struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaction.TransactionDataInterface {
	return &TransactionData{
		db: db,
	}
}

func (ad *TransactionData) GetAll() ([]transaction.TransactionInfo, error) {
	// var listTransaction = []transaction.ArticleInfo{}
	// var qry = ad.db.Table("articles").Select("transaction.*,users.name as user_name,article_categories.name as category_name").
	// 	Joins("JOIN users on users.id = transaction.user_id").
	// 	Joins("JOIN article_categories ON article_categories.id = transaction.category_id").
	// 	Where("transaction.deleted_at is null").
	// 	Scan(&listArticle)

	// if err := qry.Error; err != nil {
	// 	logrus.Info("DB error : ", err.Error())
	// 	return nil, err
	// }

	// return listArticle, nil
	return nil, nil
}

func (ad *TransactionData) GetByID(id int) ([]transaction.TransactionInfo, error) {
	// var listArticle = []transaction.ArticleInfo{}
	// var qry = ad.db.Table("articles").Select("transaction.*,users.name as user_name,article_categories.name as category_name").
	// 	Joins("LEFT JOIN users on users.id = transaction.user_id").Joins("LEFT JOIN article_categories ON article_categories.id = transaction.category_id").
	// 	Where("transaction.id = ?", id).
	// 	Where("transaction.deleted_at is null").
	// 	Scan(&listArticle)

	// if err := qry.Error; err != nil {
	// 	logrus.Info("DB error : ", err.Error())
	// 	return nil, err
	// }
	// return listArticle, nil
	return nil, nil
}

func (ad *TransactionData) Insert(newData transaction.Transaction) (*transaction.Transaction, error) {
	var dbData = new(Transaction)

	dbData.UserID = newData.UserID
	dbData.MidtransID = *newData.MidtransID

	dbData.PriceResult = *newData.PriceResult

	dbData.PaymentStatus = *&newData.PaymentStatus
	dbData.PaymentType = *&newData.PaymentType

	if err := ad.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}
