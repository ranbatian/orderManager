package bootstrap

import (
	"github.com/simon/orderManager/internal/config"
	"github.com/simon/orderManager/internal/db"
	logger "github.com/simon/orderManager/internal/log"
)

func Init() {
	config.InitConf()
	logger.InitLogger()
	db.InitDB()
}
