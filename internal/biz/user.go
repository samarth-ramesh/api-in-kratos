package biz

import (
	"context"
	"fmt"

	v1 "accountsapi/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
	ErrUserExist    = errors.BadRequest(v1.ErrorReason_USER_EXISTS.String(), "user exists")
)

// User is a User model.
type User struct {
	Username string `field:"username"`
	Password string `field:"password"`
	UserId   string `field:"rowid"`
}

// UserRepo is a User repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListByUserName(context.Context, string) ([]*User, error)
	ListAll(context.Context) ([]*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWTGrant(userId string) string {
	return ""
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) (*User, error) {
	rows, err := uc.repo.ListByUserName(ctx, g.Username)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch rows: %w", err)
	}
	if len(rows) > 0 {
		return nil, errors.New(400, v1.ErrorReason_USER_EXISTS.String(), "")
	}
	hashed, err := hashPassword(g.Password)
	if err != nil {
		return nil, err
	}
	g.Password = hashed
	return uc.repo.Save(ctx, g)
}

func (uc *UserUsecase) SignIn(ctx context.Context, g *User) (string, error) {
	rows, err := uc.repo.ListByUserName(ctx, g.Username)
	if err != nil {
		return "", ErrUserNotFound
	}
	if len(rows) == 0 {
		return "", ErrUserNotFound
	}
	isCorrectPassword := checkPasswordHash(g.Password, rows[0].Password)
	if !isCorrectPassword {
		return "", ErrUserNotFound
	}
	return generateJWTGrant(rows[0].UserId), nil
}
