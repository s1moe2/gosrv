package handlers

import (
	"fmt"
	"net/http"

	"github.com/s1moe2/gosrv/models"
)

// UsersHandler holds handler dependencies
type UsersHandler struct {
	userRepo models.UserRepository
}

// NewBaseHandler returns a new BaseHandler
func NewUsersHandler(userRepo models.UserRepository) *UsersHandler {
	return &UsersHandler{
		userRepo: userRepo,
	}
}

// Get gets all users
func (h *UsersHandler) Get(w http.ResponseWriter, r *http.Request) {
	if user, err := h.userRepo.FindByID("1"); err != nil {
		fmt.Println("Error", user)
	}

	w.Write([]byte("Hello, World"))
}

// GetByID tries to get a user by ID
func (h *UsersHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if user, err := h.userRepo.FindByID("1"); err != nil {
		fmt.Println("Error", user)
	}

	w.Write([]byte("Hello, World"))
}
