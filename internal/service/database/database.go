package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"test-api/internal/config"
)

var db *sqlx.DB

func New(cfg *config.Db) (*sqlx.DB, error) {
	if db == nil {
		var err error
		db, err = sqlx.Open(cfg.Dsn, ":memory:")

		if err != nil {
			return nil, errors.Wrap(err, "failed to connect database")
		}
	}

	return db, nil
}
