package db

import (
	"fmt"

	"github.com/simon/orderManager/internal/config"
	logger "github.com/simon/orderManager/internal/log"
	"github.com/simon/orderManager/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dsn := createDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("create db error")
		panic("create db error")
	}
	DB = db
	logger.Log.Sugar().Info("create db success")
	Migration()
}

func Migration() {
	err := DB.AutoMigrate(&model.User{}, &model.OrderType{}, &model.OrderKind{})
	if err != nil {
		logger.Log.Sugar().Panic("crate table error")
		panic("crate table error")
	}
	logger.Log.Info("create table successful.")
}

func createDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Database.Username,
		config.Conf.Database.Password,
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.DBName,
	)
}
