package service

import (
	"context"

	"github.com/carousell/chope-assignment/model"
	pb "github.com/carousell/chope-assignment/proto"
	"github.com/carousell/chope-assignment/store"
)

type Svc struct {
	Storage store.StorageService
	pb.UnimplementedUserLoginServer
}
type UserLoginServer interface {
	pb.UserLoginServer
}

func (s *Svc) Login(ctx context.Context, req *pb.LogInRequest) (*pb.LogInResponse, error) {
	user := model.AccountsUser{}
	user.Email.Scan("ganesh")
	user.Passowrd.Scan("sqweqwe")
	user.Username.Scan("GA")
	s.Storage.CreateUser(ctx, &user)
	return &pb.LogInResponse{Token: "NAE"}, nil
}

func (s *Svc) mustEmbedUnimplementedUserLoginServer() {
}
