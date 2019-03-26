package config

import "context"

type Options struct {
	*Db
}

type Db struct {
	Dsn string
}

func Get(ctx context.Context) (*Options, error) {
	var err error

	o := &Options{}

	o.Db, err = dbConfig(ctx)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func dbConfig(ctx context.Context) (*Db, error) {
	return &Db{Dsn: DbDsn}, nil
}
