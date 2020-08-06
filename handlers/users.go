package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
	users, err := h.userRepo.GetAll(r.Context())
	if err != nil {
		respondError(w, internalError())
		return
	}

	respond(w, users, http.StatusOK)
}

// GetByID tries to get a user by ID
func (h *UsersHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, ok := vars["id"]
	if !ok {
		respondError(w, appError{
			Status:  http.StatusBadRequest,
			Message: "invalid id param",
		})
		return
	}

	user, err := h.userRepo.FindByID(r.Context(), uid)
	if err != nil {
		respondError(w, internalError())
		return
	}

	if user == nil {
		respondError(w, appError{
			Status:  http.StatusNotFound,
			Message: "user not found",
		})
		return
	}

	respond(w, user, http.StatusOK)
}

// Create creates a new user
func (h *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var userPayload models.User
	err := decoder.Decode(&userPayload)
	if err != nil {
		respondError(w, internalError())
		return
	}

	userCheck, err := h.userRepo.FindByEmail(r.Context(), userPayload.Email)
	if err != nil {
		respondError(w, internalError())
		return
	}

	if userCheck != nil {
		respondError(w, appError{
			Status:  http.StatusBadRequest,
			Message: "email already in use",
		})
		return
	}

	user, err := h.userRepo.Create(&userPayload)
	if err != nil {
		respondError(w, internalError())
		return
	}

	respond(w, user, http.StatusCreated)
}

// Update updates a user
func (h *UsersHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, ok := vars["id"]
	if !ok {
		respondError(w, appError{
			Status:  http.StatusBadRequest,
			Message: "invalid id param",
		})
		return
	}

	decoder := json.NewDecoder(r.Body)

	var userPayload models.User
	err := decoder.Decode(&userPayload)
	if err != nil {
		respondError(w, internalError())
		return
	}
	userPayload.ID = uid

	user, err := h.userRepo.Update(&userPayload)
	if err != nil {
		respondError(w, internalError())
		return
	}

	if user == nil {
		respondError(w, appError{
			Status:  http.StatusNotFound,
			Message: "user not found",
		})
		return
	}

	respond(w, user, http.StatusOK)
}

// Delete deletes a user
func (h *UsersHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, ok := vars["id"]
	if !ok {
		respondError(w, appError{
			Status:  http.StatusBadRequest,
			Message: "invalid id param",
		})
		return
	}

	deletedID, err := h.userRepo.Delete(uid)
	if err != nil {
		respondError(w, internalError())
		return
	}

	if deletedID == "" {
		respondError(w, appError{
			Status:  http.StatusNotFound,
			Message: "user not found",
		})
		return
	}

	respond(w, nil, http.StatusNoContent)
}
