package service

import (
	"context"
	"strconv"
	"time"

	"accountsapi/accounts/internal/biz"
	pb "accountsapi/api/accounts"
)

type AccountsService struct {
	pb.UnimplementedAccountsServer
	accountsUseCase    *biz.AccountsUseCase
	transactionUseCase *biz.TransactionUseCase
}

func NewAccountsService(accountsUseCase *biz.AccountsUseCase, transactionUseCase *biz.TransactionUseCase) *AccountsService {
	return &AccountsService{
		accountsUseCase:    accountsUseCase,
		transactionUseCase: transactionUseCase,
	}
}

func (s *AccountsService) CreateAccounts(ctx context.Context, req *pb.CreateAccountsRequest) (*pb.CreateAccountsReply, error) {
	rv, err := s.accountsUseCase.CreateAccount(ctx, &biz.Account{
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
	rv, err := s.accountsUseCase.UpdateAccountById(ctx, &biz.Account{
		Name: req.Name,
		Id:   req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAccountsReply{
		Id: rv.Id,
	}, nil
}
func (s *AccountsService) DeleteAccounts(ctx context.Context, req *pb.DeleteAccountsRequest) (*pb.DeleteAccountsReply, error) {
	return &pb.DeleteAccountsReply{}, nil
}

func (s *AccountsService) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	accountId, _ := strconv.Atoi(req.Id)
	rv, err := s.accountsUseCase.ListAccountById(ctx, int64(accountId))
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountReply{
		Id:   rv.Id,
		Name: rv.Name,
	}, nil
}

func (s *AccountsService) ListAccounts(ctx context.Context, _ *pb.ListAccountsRequest) (*pb.ListAccountsReply, error) {
	rows, err := s.accountsUseCase.ListAccounts(ctx)
	if err != nil {
		return nil, err
	}
	rv := new(pb.ListAccountsReply)
	rv.Accounts = make([]*pb.GetAccountReply, 0)
	for _, j := range rows {
		rv.Accounts = append(rv.Accounts, &pb.GetAccountReply{
			Id:   j.Id,
			Name: j.Name,
		})
	}
	return rv, nil
}

func (s *AccountsService) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionReply, error) {
	transaction, err := s.transactionUseCase.CreateTransaction(ctx, &biz.Transaction{
		Account1:        req.AccountSource,
		Account2:        req.AccountDest,
		Amount:          int64(req.Amount),
		TransactionTime: time.Unix(req.Time, 0),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateTransactionReply{
		Id: transaction.Id,
	}, nil
}
func (s *AccountsService) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionRequest, error) {
	return &pb.UpdateTransactionRequest{}, nil
}
func (s *AccountsService) DeleteTransaction(ctx context.Context, req *pb.DeleteTransactionRequest) (*pb.DeleteAccountsReply, error) {
	return &pb.DeleteAccountsReply{}, nil
}
func (s *AccountsService) GetTransaction(ctx context.Context, req *pb.DeleteAccountsRequest) (*pb.UpdateAccountsReply, error) {
	return &pb.UpdateAccountsReply{}, nil
}
func (s *AccountsService) ListTransactions(ctx context.Context, req *pb.ListTransactionsRequest) (*pb.ListTransactionsReply, error) {
	return &pb.ListTransactionsReply{}, nil
}
