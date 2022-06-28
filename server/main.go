package main

import (
	"github.com/jinzhu/gorm"

	"github.com/macduyhai/grpcApi/server/cmd"
	"github.com/macduyhai/grpcApi/server/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Server grpc is starting ...")
	conf := config.NewConfig()
	if conf == nil {
		log.Fatal("Read config error")

	}
	// Connect DB Mysql
	db, err := gorm.Open("mysql", conf.MYSQLURL)
	if err != nil {
		log.Error("Connect DB error")

	} else {
		log.Info("DB Connected")
	}
	defer func() {
		if db.Close() != nil {
			log.Fatal("Close DB error:" + err.Error())

		}
	}()
	err = cmd.RunServer(conf, db, ":8899")
	if err != nil {
		log.Fatal("Run server err:" + err.Error())
	}

}
