package server

import (
	"github.com/gorilla/mux"
	"github.com/s1moe2/gosrv/config"
	"github.com/s1moe2/gosrv/db"
	"github.com/s1moe2/gosrv/repositories"
)

// Run handles the API server configuration and setup before starting the HTTP server
func Run() error {
	conf := config.New()

	dbConn, err := db.ConnectDB(conf.Database)
	if err != nil {
		return err
	}
	userRepo := repositories.NewUserRepo(dbConn)

	router := mux.NewRouter()
	setupUsersRouter(router, userRepo)

	srv := newServer(conf.Server, router)
	return srv.start()
}
