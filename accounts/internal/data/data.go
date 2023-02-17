package data

import (
	"accountsapi/accounts/internal/biz"
	"accountsapi/accounts/internal/conf"
	"database/sql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet2 is data providers.
var ProviderSet2 = wire.NewSet(NewData, wire.Bind(new(biz.AccountRepo), new(*AccountsRepo)), NewAccountsRepo)

type Data struct {
	// TODO wrapped database client
	Db *sql.DB
}

// NewData .
func NewData(confData *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := sql.Open(confData.Database.Driver, confData.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		Db: db,
	}, cleanup, nil
}
