package biz

import "github.com/google/wire"

// ProviderSet3 is biz providers.
var ProviderSet3 = wire.NewSet(NewAccountsUseCase)
