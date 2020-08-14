package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/s1moe2/gosrv/repositories"
	"net/http"
	"regexp"

	"github.com/s1moe2/gosrv/models"
)

// UsersHandler holds handler dependencies
type UsersHandler struct {
	userRepo models.UserRepository
}

type UserPayload struct {
	Name  string
	Email string
}

func (p *UserPayload) validate() []error {
	var errs []error

	if len(p.Name) < 3 {
		errs = append(errs, errors.New("name: invalid length"))
	}

	const emailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	emailRegexp := regexp.MustCompile(emailRegex)
	if !emailRegexp.MatchString(p.Email) {
		errs = append(errs, errors.New("email: invalid format"))
	}

	return errs
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
		respondInternalError(w)
		return
	}

	respond(w, users, http.StatusOK)
}

// GetByID tries to get a user by ID
func (h *UsersHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, ok := vars["id"]
	if !ok {
		respondError(w, newSimpleUserError(errors.New("invalid id param")))
		return
	}

	user, err := h.userRepo.FindByID(r.Context(), uid)
	if err != nil {
		respondInternalError(w)
		return
	}

	if user == nil {
		respondError(w, &userError{
			Status: http.StatusNotFound,
			Errors: []error{errors.New("user not found")},
		})
		return
	}

	respond(w, user, http.StatusOK)
}

// Create creates a new user
func (h *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var userPayload UserPayload
	err := decoder.Decode(&userPayload)
	if err != nil {
		respondInternalError(w)
		return
	}

	errs := userPayload.validate()
	if errs != nil {
		respondError(w, newUserError(errs))
		return
	}

	userCheck, err := h.userRepo.FindByEmail(r.Context(), userPayload.Email)
	if err != nil {
		respondInternalError(w)
		return
	}

	if userCheck != nil {
		respondError(w, newSimpleUserError(errors.New("email already in use")))
		return
	}

	user, err := h.userRepo.Create(&models.User{
		Name:  userPayload.Name,
		Email: userPayload.Email,
	})
	if err != nil {
		respondInternalError(w)
		return
	}

	respond(w, user, http.StatusCreated)
}

// Update updates a user
func (h *UsersHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, ok := vars["id"]
	if !ok {
		respondError(w, newSimpleUserError(errors.New("invalid id param")))
		return
	}

	decoder := json.NewDecoder(r.Body)

	var userPayload UserPayload
	err := decoder.Decode(&userPayload)
	if err != nil {
		respondInternalError(w)
		return
	}

	errs := userPayload.validate()
	if errs != nil {
		respondError(w, newUserError(errs))
		return
	}

	user, err := h.userRepo.Update(&models.User{
		ID:    uid,
		Name:  userPayload.Name,
		Email: userPayload.Email,
	})
	if err != nil {
		if e, ok := err.(*repositories.ConflictError); ok {
			respondError(w, newSimpleUserError(e))
			return
		}

		respondInternalError(w)
		return
	}

	if user == nil {
		respondError(w, &userError{
			Status: http.StatusNotFound,
			Errors: []error{errors.New("user not found")},
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
		respondError(w, newSimpleUserError(errors.New("invalid id param")))
		return
	}

	deleted, err := h.userRepo.Delete(uid)
	if err != nil {
		respondInternalError(w)
		return
	}

	if !deleted {
		respondError(w, &userError{
			Status: http.StatusNotFound,
			Errors: []error{errors.New("user not found")},
		})
		return
	}

	respond(w, nil, http.StatusNoContent)
}
