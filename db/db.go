package db

import (
	"database/sql"
	"github.com/s1moe2/gosrv/config"
)

// ConnectDB opens a connection to the database.
// Since this is expected to successfully happen at server start, it will panic in case of error.
func ConnectDB(dbConfig config.DatabaseConfig) *sql.DB {
	db, err := sql.Open(dbConfig.Driver, dbConfig.URI)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
