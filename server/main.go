package server

import (
	"github.com/gorilla/mux"
	"github.com/s1moe2/gosrv/models"
	"net/http"
	"time"

	"github.com/s1moe2/gosrv/config"
	"github.com/s1moe2/gosrv/db"
	"github.com/s1moe2/gosrv/handlers"
	"github.com/s1moe2/gosrv/repositories"
)

func Start() {
	conf := config.New()

	dbConn := db.ConnectDB(conf.Database)
	userRepo := repositories.NewUserRepo(dbConn)

	router := mux.NewRouter()
	setupUsersRouter(router, userRepo)

	srv := newServer(conf.Server, router)

	srv.ListenAndServe()

}

func newServer(serverConfig config.ServerConfig, router *mux.Router) *http.Server {
	return &http.Server{
		Addr: serverConfig.Address,
		//ErrorLog:     log.New(logrus.New().Writer(), "", 0),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}

func setupUsersRouter(router *mux.Router, repo models.UserRepository) {
	h := handlers.NewUsersHandler(repo)

	ur := router.PathPrefix("/users").Subrouter()

	ur.HandleFunc("/", h.Get).
		Methods(http.MethodGet)
	ur.HandleFunc("/{id}", h.GetByID).
		Methods(http.MethodGet)
}
