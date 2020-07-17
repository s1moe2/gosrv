package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/s1moe2/gosrv/config"
)

// ConnectDB opens a connection to the database.
// Since this is expected to successfully happen at server start, it will panic in case of error.
func ConnectDB(dbConfig config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open(dbConfig.Driver, dbConfig.URI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open db connection")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to ping db")
	}
	return db, nil
}
