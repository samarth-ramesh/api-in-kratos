package data

import (
	"accountsapi/internal/conf"
	"database/sql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Db        *sql.DB
	JwtSecret string
}

// NewData .
func NewData(confData *conf.Data, confServer *conf.Server, logger log.Logger) (*Data, func(), error) {
	db, err := sql.Open(confData.Database.Driver, confData.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		Db:        db,
		JwtSecret: confServer.RandomKey,
	}, cleanup, nil
}
