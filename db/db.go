package db

import (
	"log"

	"github.com/Kritsana135/assessment/domain"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	isRelase := viper.GetString("GIN_MODE") == "release"
	var gormConfig *gorm.Config
	if !isRelase {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	} else {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}

	url := viper.GetString("DATABASE_URL")
	dial := postgres.Open(url)
	var err error
	db, err = gorm.Open(dial, gormConfig)
	if err != nil {
		log.Fatal("connect db error : ", err)
	}

	logrus.Info("database connected")

	autoMigrate := viper.GetBool("AUTO_MIGRATE")
	if autoMigrate {
		db.AutoMigrate(&domain.ExpenseTable{})
	}

	return db
}

func CloseDatabase() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}

func GetDb() *gorm.DB {
	return db
}
