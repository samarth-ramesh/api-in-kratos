package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"time"
)

type TransactionRepo interface {
	Save(context.Context, *Transaction) (*Transaction, error)
	Update(context.Context, *Transaction) (*Transaction, error)
	Delete(ctx context.Context, transaction *Transaction) error
	FindByID(context.Context, int64) (*Transaction, error)
}

type Transaction struct {
	Id              string
	Account1        string
	Account2        string
	Amount          int64
	TransactionTime time.Time
}

// TransactionUseCase is an Transactions usecase.
type TransactionUseCase struct {
	transactionRepo TransactionRepo
	accountRepo     AccountRepo
	log             *log.Helper
}

func NewTransactionUseCase(transactionRepo TransactionRepo, accountRepo AccountRepo, logger log.Logger) *TransactionUseCase {
	return &TransactionUseCase{transactionRepo: transactionRepo, accountRepo: accountRepo, log: log.NewHelper(logger)}
}

func (uc *TransactionUseCase) CreateTransaction(ctx context.Context, transaction *Transaction) (*Transaction, error) {
	newId1, _ := strconv.Atoi(transaction.Account1)
	acc1, err := uc.accountRepo.FindByID(ctx, int64(newId1))
	if err != nil {
		return nil, err
	}
	if acc1 == nil || acc1.UserId != UserIdFromContext(ctx) {
		return nil, ErrNotFound
	}

	newId2, _ := strconv.Atoi(transaction.Account2)
	acc2, err := uc.accountRepo.FindByID(ctx, int64(newId2))
	if err != nil {
		return nil, err
	}
	if acc2 == nil || acc2.UserId != UserIdFromContext(ctx) {
		return nil, ErrNotFound
	}

	save, err := uc.transactionRepo.Save(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return save, err
}

func (uc *TransactionUseCase) UpdateTransaction(ctx context.Context, transaction *Transaction) (*Transaction, error) {
	newId1, _ := strconv.Atoi(transaction.Account1)
	acc1, err := uc.accountRepo.FindByID(ctx, int64(newId1))
	if err != nil {
		return nil, err
	}
	if acc1 == nil || acc1.UserId != UserIdFromContext(ctx) {
		return nil, ErrNotFound
	}

	newId2, _ := strconv.Atoi(transaction.Account2)
	acc2, err := uc.accountRepo.FindByID(ctx, int64(newId2))
	if err != nil {
		return nil, err
	}
	if acc2 == nil || acc2.UserId != UserIdFromContext(ctx) {
		return nil, ErrNotFound
	}
	return uc.transactionRepo.Update(ctx, transaction)
}

func (uc *TransactionUseCase) DeleteTransaction(ctx context.Context, _transactionId string) error {
	transactionId, _ := strconv.Atoi(_transactionId)
	transaction, err := uc.transactionRepo.FindByID(ctx, int64(transactionId))
	if err != nil {
		return err
	}
	if transaction == nil {
		return ErrNotFound
	}
	newId1, _ := strconv.Atoi(transaction.Account1)
	acc1, err := uc.accountRepo.FindByID(ctx, int64(newId1))
	if err != nil {
		return err
	}
	if acc1 == nil || acc1.UserId != UserIdFromContext(ctx) {
		return ErrNotFound
	}

	newId2, _ := strconv.Atoi(transaction.Account2)
	acc2, err := uc.accountRepo.FindByID(ctx, int64(newId2))
	if err != nil {
		return err
	}
	if acc2 == nil || acc2.UserId != UserIdFromContext(ctx) {
		return ErrNotFound
	}
	err = uc.transactionRepo.Delete(ctx, transaction)
	if err != nil {
		return err
	}
	return nil
}
