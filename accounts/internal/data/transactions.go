package data

import (
	"accountsapi/accounts/internal/biz"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
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

func (repo *TransactionRepo) Update(ctx context.Context, transaction *biz.Transaction) (*biz.Transaction, error) {
	_, err := repo.data.Db.ExecContext(ctx, "UPDATE `transaction` SET account1 = ?, account2 = ?, amount = ?, date = ? WHERE ROWID = ?", transaction.Account1, transaction.Account2, transaction.Amount, transaction.TransactionTime.Unix(), transaction.Id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (repo *TransactionRepo) Delete(ctx context.Context, transaction *biz.Transaction) error {
	_, err := repo.data.Db.ExecContext(ctx, "DELETE FROM `transaction` WHERE ROWID = ?", transaction.Id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TransactionRepo) FindByID(ctx context.Context, transactionId int64) (*biz.Transaction, error) {
	row := repo.data.Db.QueryRowContext(ctx, "SELECT ROWID, account1, account2, amount, date FROM `transaction` WHERE ROWID = ?", transactionId)
	transaction := new(biz.Transaction)
	transactionTime := 0
	err := row.Scan(&transaction.Id, &transaction.Account1, &transaction.Account2, &transaction.Amount, &transactionTime)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	transaction.TransactionTime = time.Unix(int64(transactionTime), 0)
	return transaction, nil
}
