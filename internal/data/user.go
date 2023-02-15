package data

import (
	"context"
	"fmt"

	"accountsapi/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	ret, err := r.data.Db.Exec("INSERT INTO user VALUES (?, ?)", g.Username, g.Password)
	if err != nil {
		return nil, err
	}
	lastId, err := ret.LastInsertId()
	if err != nil {
		return nil, err
	}
	g.UserId = fmt.Sprint(lastId)
	return g, nil
}

func (r *userRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *userRepo) FindByID(context.Context, int64) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) ListByUserName(ctx context.Context, term string) (rv []*biz.User, err error) {
	rows, err := r.data.Db.QueryContext(ctx, "SELECT username, password, rowid FROM user WHERE username = ?", term)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := new(biz.User)
		err = rows.Scan(&user.Username, &user.Password, &user.UserId)
		if err != nil {
			return rv, err
		}
		rv = append(rv, user)
	}
	return rv, nil
}

func (r *userRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}

func (r *userRepo) GetSigningKey() string {
	return r.data.JwtSecret
}
