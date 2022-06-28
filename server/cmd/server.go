package cmd

import (
	"net"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/macduyhai/grpcApi/apiproto"
	"github.com/macduyhai/grpcApi/server/api"
	"github.com/macduyhai/grpcApi/server/config"
	"github.com/macduyhai/grpcApi/server/daos"
	"github.com/macduyhai/grpcApi/server/services"
	"google.golang.org/grpc"
)

func RunServer(conf *config.Config, db *gorm.DB, port string) error {
	// Run grpc server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Error("Error while create listen %v", err)
		return err
	}
	defer func() {
		if lis.Close() != nil {
			log.Fatal("Close tcp server error:" + err.Error())
		}
	}()
	userDao := daos.NewUserDao(db)
	UserService := services.NewUserService(conf, userDao)
	serverApi := api.NewServerApi(UserService)

	s := grpc.NewServer()
	apiproto.RegisterUserServiceServer(s, serverApi)
	err = s.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
