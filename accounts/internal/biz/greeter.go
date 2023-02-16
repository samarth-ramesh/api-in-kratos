package biz

import (
	"context"

	"accountsapi/api/accounts"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrAccountExists is user not found.
	ErrAccountExists = errors.BadRequest(accounts.ErrorReason_ACCOUNT_EXISTS.String(), "account exists")
)

// Account is a Account model.
type Account struct {
	Name string `field:"name"`
	Id   string `field:"rowid"`
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Account) (*Account, error)
	Update(context.Context, *Account) (*Account, error)
	FindByID(context.Context, int64) (*Account, error)
	ListAll(context.Context) ([]*Account, error)
	FindByName(context.Context, string) ([]*Account, error)
}

// AccountsUseCase is a Greeter usecase.
type AccountsUseCase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewAccountsUseCase new a Greeter usecase.
func NewAccountsUseCase(repo GreeterRepo, logger log.Logger) *AccountsUseCase {
	return &AccountsUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *AccountsUseCase) CreateGreeter(ctx context.Context, g *Account) (*Account, error) {
	rows, err := uc.repo.FindByName(ctx, g.Name)
	if err != nil {
		return nil, err
	}
	if len(rows) > 0 {
		return nil, ErrAccountExists
	}
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
