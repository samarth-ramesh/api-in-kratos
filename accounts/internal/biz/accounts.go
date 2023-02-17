package biz

import (
	"context"
	"encoding/json"
	"strconv"

	"accountsapi/api/accounts"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
)

var (
	// ErrAccountExists is user not found.
	ErrAccountExists = errors.BadRequest(accounts.ErrorReason_ACCOUNT_EXISTS.String(), "account exists")
	ErrNotFound      = errors.NotFound("Not Found", "")
)

// Account is an Account model.
type Account struct {
	Name   string `field:"name"`
	Id     string `field:"rowid"`
	UserId string `field:"userId"`
}

// AccountRepo is an Accounts  accountRepo.
type AccountRepo interface {
	Save(context.Context, *Account) (*Account, error)
	Update(context.Context, *Account) (*Account, error)
	FindByID(context.Context, int64) (*Account, error)
	ListAll(context.Context, string) ([]*Account, error)
	FindByName(context.Context, string, string) ([]*Account, error)
}

// AccountsUseCase is an Accounts usecase.
type AccountsUseCase struct {
	accountRepo AccountRepo
	log         *log.Helper
}

// NewAccountsUseCase return a new Account usecase.
func NewAccountsUseCase(accountRepo AccountRepo, logger log.Logger) *AccountsUseCase {
	return &AccountsUseCase{accountRepo: accountRepo, log: log.NewHelper(logger)}
}

func UserIdFromContext(ctx context.Context) string {
	claims, _ := jwt.FromContext(ctx)
	b, _ := json.Marshal(claims)
	m := new(map[string]interface{})
	json.Unmarshal(b, m)
	m2 := *m
	return string(m2["sub"].(string))
}

// CreateAccount creates a Greeter, and returns the new Greeter.
func (uc *AccountsUseCase) CreateAccount(ctx context.Context, g *Account) (*Account, error) {
	rows, err := uc.accountRepo.FindByName(ctx, g.Name, UserIdFromContext(ctx))
	if err != nil {
		return nil, err
	}
	if len(rows) > 0 {
		return nil, ErrAccountExists
	}
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Name)
	g.UserId = UserIdFromContext(ctx)
	return uc.accountRepo.Save(ctx, g)
}

func (uc *AccountsUseCase) ListAccounts(ctx context.Context) ([]*Account, error) {
	rv, err := uc.accountRepo.ListAll(ctx, UserIdFromContext(ctx))
	return rv, err
}

func (uc *AccountsUseCase) ListAccountById(ctx context.Context, accountId int64) (*Account, error) {
	account, err := uc.accountRepo.FindByID(ctx, accountId)
	if err != nil {
		return nil, err
	}
	if account == nil {
		uc.log.Debug("No Account Found for ID")
		return nil, ErrNotFound
	}
	if account.UserId != UserIdFromContext(ctx) {
		uc.log.Debug("Perm denied")
		return nil, ErrNotFound
	}
	return account, nil
}

func (uc *AccountsUseCase) UpdateAccountById(ctx context.Context, account *Account) (*Account, error) {
	accountId, _ := strconv.Atoi(account.Id)
	acc, err := uc.accountRepo.FindByID(ctx, int64(accountId))
	if err != nil {
		return nil, err
	}
	if acc == nil {
		return nil, ErrNotFound
	}
	if acc.UserId != UserIdFromContext(ctx) {
		return nil, ErrNotFound
	}
	updated, err := uc.accountRepo.Update(ctx, account)
	if err != nil {
		return nil, err
	}
	return updated, nil
}
