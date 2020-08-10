package server

import (
	"github.com/gorilla/mux"
	"github.com/s1moe2/gosrv/config"
	"github.com/s1moe2/gosrv/db"
	"github.com/s1moe2/gosrv/repositories"
	"net/http"
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
	router.Use(loggingMiddleware)
	setupUsersRouter(router, userRepo)

	fs := http.FileServer(http.Dir("./swaggerui/"))
	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", fs))

	srv := newServer(conf.Server, router)
	return srv.start()
}
