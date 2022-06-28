package main

import (
	"context"

	"github.com/macduyhai/grpcApi/apiproto"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Add(ctx context.Context, req *apiproto.Addrequest, opts ...grpc.CallOption) (*apiproto.AddResponse, error) {
	// username, password, fullname, old := req.Username, req.Password, req.Fullname, req.Old

	response := &apiproto.AddResponse{}
	return response, nil
}

func main() {

}
