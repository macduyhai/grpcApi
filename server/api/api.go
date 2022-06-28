package api

import (
	"context"
	"log"

	"github.com/macduyhai/grpcApi/apiproto"
	"github.com/macduyhai/grpcApi/server/models"
	"github.com/macduyhai/grpcApi/server/services"
	"github.com/macduyhai/grpcApi/server/utility"
)

type ServerApi struct {
	apiproto.UnimplementedUserServiceServer
	userService services.UserService
}

func NewServerApi(userService services.UserService) *ServerApi {
	return &ServerApi{
		userService: userService,
	}
}

func (s *ServerApi) Add(ctx context.Context, req *apiproto.Addrequest) (*apiproto.AddResponse, error) {
	response := &apiproto.AddResponse{}
	timeNow := utility.TimeIn("Asia/Ho_Chi_Minh")
	user := models.User{
		Username:   req.Username,
		Password:   req.Password,
		Fullname:   req.Fullname,
		Salary:     req.Salary,
		CreateTime: &timeNow,
	}
	_, err := s.userService.Add(user)
	if err != nil {
		response.Code = 400
		response.Message = "Create user error"
		log.Println(response)
		return response, err
	}
	response.Code = 200
	response.Message = "Create user success"
	return response, nil
}
