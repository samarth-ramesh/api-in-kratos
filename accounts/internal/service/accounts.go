package service

import (
	"context"

	pb "accountsapi/api/accounts"
)

type AccountsService struct {
	pb.UnimplementedAccountsServer
}

func NewAccountsService() *AccountsService {
	return &AccountsService{}
}

func (s *AccountsService) CreateAccounts(ctx context.Context, req *pb.CreateAccountsRequest) (*pb.CreateAccountsReply, error) {
	return &pb.CreateAccountsReply{}, nil
}
func (s *AccountsService) UpdateAccounts(ctx context.Context, req *pb.UpdateAccountsRequest) (*pb.UpdateAccountsReply, error) {
	return &pb.UpdateAccountsReply{}, nil
}
func (s *AccountsService) DeleteAccounts(ctx context.Context, req *pb.DeleteAccountsRequest) (*pb.DeleteAccountsReply, error) {
	return &pb.DeleteAccountsReply{}, nil
}
func (s *AccountsService) GetAccounts(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	return &pb.GetAccountReply{}, nil
}
func (s *AccountsService) ListAccounts(ctx context.Context, req *pb.ListAccountsRequest) (*pb.ListAccountsReply, error) {
	return &pb.ListAccountsReply{}, nil
}
