package service

import (
	"context"

	"accountsapi/accounts/internal/biz"
	pb "accountsapi/api/accounts"
)

type AccountsService struct {
	pb.UnimplementedAccountsServer
	uc *biz.AccountsUseCase
}

func NewAccountsService(uc *biz.AccountsUseCase) *AccountsService {
	return &AccountsService{
		uc: uc,
	}
}

func (s *AccountsService) CreateAccounts(ctx context.Context, req *pb.CreateAccountsRequest) (*pb.CreateAccountsReply, error) {
	rv, err := s.uc.CreateGreeter(ctx, &biz.Account{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountsReply{
		Id: rv.Id,
	}, nil
}
func (s *AccountsService) UpdateAccounts(ctx context.Context, req *pb.UpdateAccountsRequest) (*pb.UpdateAccountsReply, error) {
	return &pb.UpdateAccountsReply{}, nil
}
func (s *AccountsService) DeleteAccounts(ctx context.Context, req *pb.DeleteAccountsRequest) (*pb.DeleteAccountsReply, error) {
	return &pb.DeleteAccountsReply{}, nil
}
func (s *AccountsService) GetAccounts(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	s.GetAccount(ctx, req)
	return &pb.GetAccountReply{}, nil
}
func (s *AccountsService) ListAccounts(ctx context.Context, req *pb.ListAccountsRequest) (*pb.ListAccountsReply, error) {
	return &pb.ListAccountsReply{}, nil
}
