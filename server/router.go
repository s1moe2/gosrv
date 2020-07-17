package server

import (
	"github.com/gorilla/mux"
	"github.com/s1moe2/gosrv/handlers"
	"github.com/s1moe2/gosrv/models"
	"net/http"
)

func setupUsersRouter(router *mux.Router, repo models.UserRepository) {
	h := handlers.NewUsersHandler(repo)

	ur := router.
		PathPrefix("/users").
		Subrouter()

	ur.Methods(http.MethodGet).
		Path("/").
		HandlerFunc(h.Get)

	ur.Methods(http.MethodGet).
		Path("/{id}").
		HandlerFunc(h.GetByID)
}
