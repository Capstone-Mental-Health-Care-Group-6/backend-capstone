package database

import (
	"FinalProject/configs"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c configs.ProgrammingConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPass,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)

	// fmt.Println(dsn)
	//DEV MODE
	// dsn := "root:@tcp(127.0.0.1:3306)/mentalhealth?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("terjadi kesalahan pada database, error:", err.Error())
		return nil, err
	}

	return db, nil
}
