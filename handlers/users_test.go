package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/s1moe2/gosrv/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func prepareRouter(method string, path string, h func(http.ResponseWriter, *http.Request)) *mux.Router{
	router := mux.NewRouter()
	router.Methods(method).
		Path(path).
		HandlerFunc(h)
	return router
}

func TestUsersHandler_Get(t *testing.T) {
	t.Run("expect GET /users to return 200 and a list of users", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.getAllImpl = func() ([]*models.User, error) {
			var res []*models.User
			for  _, usr := range mock.mockUsers {
				res = append(res, usr)
			}
			return res, nil
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodGet, "/users", uh.Get)
		router.ServeHTTP(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected %d response, got %d", http.StatusOK, resp.StatusCode)
		}

		if resp.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("expected 'application/json', got '%s'", resp.Header.Get("Content-Type"))
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var users []*models.User
		json.Unmarshal(body, &users)

		if len(users) != len(mock.mockUsers) {
			t.Fatal("wrong user set")
		}
	})

	t.Run("expect GET /users to return 200 and an empty list of users", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.getAllImpl = func() ([]*models.User, error) {
			return []*models.User{}, nil
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodGet, "/users", uh.Get)
		router.ServeHTTP(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected %d response, got %d", http.StatusOK, resp.StatusCode)
		}

		if resp.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("expected 'application/json', got '%s'", resp.Header.Get("Content-Type"))
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var users []*models.User
		json.Unmarshal(body, &users)

		if len(users) != 0 {
			t.Fatal("wrong user set")
		}
	})

	t.Run("expect GET /users to return 500 on internal error", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.getAllImpl = func() ([]*models.User, error) {
			return nil, errors.New("repo error")
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodGet, "/users", uh.Get)
		router.ServeHTTP(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected %d response, got %d", http.StatusInternalServerError, resp.StatusCode)
		}
	})
}

func TestUsersHandler_GetByID(t *testing.T) {
	t.Run("expect GET /users/{id} to return 200", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.findByIDImpl = func(ID string) (*models.User, error) {
			return mock.mockUsers[ID], nil
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("GET", "/users/1", nil)
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodGet, "/users/{id}", uh.GetByID)
		router.ServeHTTP(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected %d response, got %d", http.StatusOK, resp.StatusCode)
		}

		if resp.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("expected 'application/json', got '%s'", resp.Header.Get("Content-Type"))
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var user *models.User
		json.Unmarshal(body, &user)

		if user.ID != mock.mockUsers["1"].ID {
			t.Fatal("wrong user returned")
		}
	})

	t.Run("expect GET /users/{id} to return 404 when user does not exist", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.findByIDImpl = func(ID string) (*models.User, error) {
			return nil, nil
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("GET", "/users/1", nil)
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodGet, "/users/{id}", uh.GetByID)
		router.ServeHTTP(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusNotFound {
			t.Fatalf("expected %d response, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})

	t.Run("expect GET /users/{id} to return 500 on internal error", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.findByIDImpl = func(ID string) (*models.User, error) {
			return nil, errors.New("repo error")
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("GET", "/users/1", nil)
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodGet, "/users/{id}", uh.GetByID)
		router.ServeHTTP(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected %d response", http.StatusInternalServerError)
		}
	})
}

func TestUsersHandler_Create(t *testing.T) {
	mockPayload := map[string]interface{}{
		"email": "johndoe@gosrv.com",
		"name":  "John Doe",
	}

	t.Run("expect POST /users to return 201", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.findByEmailImpl = func(email string) (*models.User, error) {
			return nil, nil
		}
		mock.createImpl = func(user *models.User) (*models.User, error) {
			user.ID = "3"
			mock.mockUsers[user.ID] = user
			return user, nil
		}
		uh := NewUsersHandler(mock)

		body, _ := json.Marshal(mockPayload)
		r := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodPost, "/users", uh.Create)
		router.ServeHTTP(w, r)
		resp := w.Result()

		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("expected %d response, got %d", http.StatusCreated, resp.StatusCode)
		}

		if resp.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("expected 'application/json', got '%s'", resp.Header.Get("Content-Type"))
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var user *models.User
		json.Unmarshal(body, &user)

		if user.ID != mock.mockUsers[user.ID].ID {
			t.Fatal("wrong user returned")
		}
	})

	t.Run("expect POST /users to return 400 when the email is in use", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.findByEmailImpl = func(email string) (*models.User, error) {
			return &models.User{
				ID:    "5",
				Name:  "John Doe",
				Email: email,
			}, nil
		}
		uh := NewUsersHandler(mock)

		body, _ := json.Marshal(mockPayload)
		r := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodPost, "/users", uh.Create)
		router.ServeHTTP(w, r)
		resp := w.Result()

		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected %d response, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})

	t.Run("expect POST /users to return 500 on find internal error", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.findByEmailImpl = func(email string) (*models.User, error) {
			return nil, errors.New("repo error")
		}
		uh := NewUsersHandler(mock)

		body, _ := json.Marshal(mockPayload)
		r := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodPost, "/users", uh.Create)
		router.ServeHTTP(w, r)
		resp := w.Result()

		if resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected %d response", http.StatusInternalServerError)
		}
	})

	t.Run("expect POST /users to return 500 on create internal error", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.findByEmailImpl = func(email string) (*models.User, error) {
			return nil, nil
		}
		mock.createImpl = func(user *models.User) (*models.User, error) {
			return nil, errors.New("repo error")
		}
		uh := NewUsersHandler(mock)

		body, _ := json.Marshal(mockPayload)
		r := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodPost, "/users", uh.Create)
		router.ServeHTTP(w, r)
		resp := w.Result()

		if resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected %d response", http.StatusInternalServerError)
		}
	})
}