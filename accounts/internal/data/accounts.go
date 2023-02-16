package data

import (
	"context"
	"fmt"

	"accountsapi/accounts/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewAccountsRepo .
func NewAccountsRepo(data *Data, logger log.Logger) *greeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Account) (*biz.Account, error) {
	res, err := r.data.Db.ExecContext(ctx, "INSERT INTO account VALUES (?, ?)", g.Name, g.UserId)
	if err != nil {
		return nil, err
	}
	lastID, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	g.Id = fmt.Sprint(lastID)
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Account) (*biz.Account, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Account, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(ctx context.Context, userId string) (res []*biz.Account, err error) {
	rows, err := r.data.Db.QueryContext(ctx, "SELECT rowid, name FROM account WHERE userId = ?", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		acc := new(biz.Account)
		err = rows.Scan(&acc.Id, &acc.Name)
		if err != nil {
			return res, err
		}
		res = append(res, acc)
	}
	return res, nil
}

func (r *greeterRepo) FindByName(ctx context.Context, name, userId string) (res []*biz.Account, err error) {
	rows, err := r.data.Db.QueryContext(ctx, "SELECT rowid, name FROM account WHERE name = ? AND userId = ?", name, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		acc := new(biz.Account)
		err = rows.Scan(&acc.Id, &acc.Name)
		if err != nil {
			return res, err
		}
		res = append(res, acc)
	}
	return res, nil
}
