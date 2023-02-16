package server

import (
	"accountsapi/accounts/internal/conf"
	v1 "accountsapi/api/helloworld/v1"
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
)

func GenAuthMiddleWare(c *conf.Server) middleware.Middleware {
	return selector.Server(
		jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
			return []byte(c.RandomKey), nil
		}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256))).
		Match(func(ctx context.Context, operation string) bool {
			if (operation == v1.OperationUserLogIn) || (operation == v1.OperationUserSignUp) {
				return false
			} else {
				return true
			}
		}).Build()
}

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)
