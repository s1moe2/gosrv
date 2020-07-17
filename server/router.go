package server

import (
	"github.com/gorilla/mux"
	"github.com/s1moe2/gosrv/handlers"
	"github.com/s1moe2/gosrv/models"
	"net/http"
)

func setupUsersRouter(router *mux.Router, repo models.UserRepository) {
	h := handlers.NewUsersHandler(repo)

	ur := router.PathPrefix("/users").Subrouter()

	ur.HandleFunc("/", h.Get).
		Methods(http.MethodGet)
	ur.HandleFunc("/{id}", h.GetByID).
		Methods(http.MethodGet)
}
