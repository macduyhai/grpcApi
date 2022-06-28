package services

import (
	"github.com/macduyhai/grpcApi/server/config"
	"github.com/macduyhai/grpcApi/server/daos"
	"github.com/macduyhai/grpcApi/server/models"
)

type UserService interface {
	Add(user models.User) (*models.User, error)
}

type UserServiceImpl struct {
	conf    *config.Config
	userDao daos.UserDao
}

func NewUserService(conf *config.Config, userDao daos.UserDao) UserService {
	return &UserServiceImpl{
		conf:    conf,
		userDao: userDao,
	}
}
func (service *UserServiceImpl) Add(user models.User) (*models.User, error) {
	return service.userDao.Create(user)
}
