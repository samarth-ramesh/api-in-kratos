package data

import (
	"accountsapi/accounts/internal/biz"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type TransactionRepo struct {
	data *Data
	log  *log.Helper
}

// NewTransactionRepo .
func NewTransactionRepo(data *Data, logger log.Logger) *TransactionRepo {
	return &TransactionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *TransactionRepo) Save(ctx context.Context, transaction *biz.Transaction) (*biz.Transaction, error) {
	result, err := repo.data.Db.ExecContext(ctx, "INSERT INTO `transaction` (account1, account2, amount, date) VALUES (?,?,?,?)", transaction.Account1, transaction.Account2, transaction.Amount, transaction.TransactionTime.Unix())
	if err != nil {
		return nil, err
	}
	newId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	transaction.Id = fmt.Sprint(newId)
	return transaction, nil
}
