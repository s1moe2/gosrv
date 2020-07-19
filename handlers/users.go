package handlers

import (
	"encoding/json"
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
	users, err := h.userRepo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// GetByID tries to get a user by ID
func (h *UsersHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("id")

	user, err := h.userRepo.FindByID(uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// Create creates a new user
func (h *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

// Update updates a user
func (h *UsersHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

// Delete deletes a user
func (h *UsersHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}