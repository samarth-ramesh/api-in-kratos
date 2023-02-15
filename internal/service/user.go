package service

import (
	"context"

	pb "accountsapi/api/helloworld/v1"
	"accountsapi/internal/biz"
)

type GreeterService struct {
	pb.UnimplementedUserServer
	biz.UserUsecase
}

func NewGreeterService(gre *biz.UserUsecase) *GreeterService {
	return &GreeterService{
		UserUsecase: *gre,
	}
}

func (s *GreeterService) SignUp(ctx context.Context, req *pb.SignUpReq) (*pb.SignUpReply, error) {
	newUser, err := s.CreateUser(ctx, &biz.User{
		Username: req.Name,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}
	return &pb.SignUpReply{
		UserId: newUser.UserId,
	}, nil
}

func (s *GreeterService) LogIn(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	authCode, err := s.UserUsecase.SignIn(ctx, &biz.User{
		Username: req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{AuthCode: authCode}, nil
}
