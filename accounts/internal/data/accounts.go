package data

import (
	"context"
	"database/sql"
	"fmt"

	"accountsapi/accounts/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AccountsRepo struct {
	data *Data
	log  *log.Helper
}

// NewAccountsRepo .
func NewAccountsRepo(data *Data, logger log.Logger) *AccountsRepo {
	return &AccountsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *AccountsRepo) Save(ctx context.Context, g *biz.Account) (*biz.Account, error) {
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

func (r *AccountsRepo) Update(ctx context.Context, g *biz.Account) (*biz.Account, error) {
	return g, nil
}

func (r *AccountsRepo) FindByID(ctx context.Context, id int64) (*biz.Account, error) {
	row := r.data.Db.QueryRowContext(ctx, "SELECT rowid, name, userId FROM account WHERE ROWID = ?", id)
	rv := new(biz.Account)
	err := row.Scan(&rv.Id, &rv.Name, &rv.UserId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (r *AccountsRepo) ListAll(ctx context.Context, userId string) (res []*biz.Account, err error) {
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

func (r *AccountsRepo) FindByName(ctx context.Context, name, userId string) (res []*biz.Account, err error) {
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
