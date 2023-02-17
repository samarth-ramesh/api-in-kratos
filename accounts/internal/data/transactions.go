package data

import "github.com/go-kratos/kratos/v2/log"

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
