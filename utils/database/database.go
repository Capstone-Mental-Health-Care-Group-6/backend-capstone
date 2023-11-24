package database

import (
	"FinalProject/configs"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c configs.ProgrammingConfig) *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	// connStr := "root:@tcp(127.0.0.1:3306)/mentalhealth?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect database : ", err.Error())
	}

	return db
}
