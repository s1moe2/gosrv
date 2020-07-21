package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/s1moe2/gosrv/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersHandler_Get(t *testing.T) {
	t.Run("expect GET /users to return 200 and a list of users", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.getAllImpl = func() ([]*models.User, error) {
			return mock.mockUsers, nil
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()
		uh.Get(w, r)

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
		uh.Get(w, r)

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
		uh.Get(w, r)

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
			return mock.mockUsers[0], nil
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("GET", "/users/1", nil)
		w := httptest.NewRecorder()
		uh.GetByID(w, r)

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

		if user.ID != mock.mockUsers[0].ID {
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
		uh.GetByID(w, r)

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
		uh.GetByID(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected %d response", http.StatusInternalServerError)
		}
	})
}

func TestUsersHandler_Create(t *testing.T) {
	mockPayload := bytes.NewBufferString(`
		"email": "johndoe@gosrv.com",
		"name": "John Doe"
	`)

	t.Run("expect POST /users to return 201", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.createImpl = func(user *models.User) (*models.User, error) {
			return user, nil
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("POST", "/users/1", mockPayload)
		w := httptest.NewRecorder()
		uh.Create(w, r)
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

		if user.ID != mock.mockUsers[0].ID {
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

		r := httptest.NewRequest("POST", "/users", mockPayload)
		w := httptest.NewRecorder()
		uh.Create(w, r)
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

		r := httptest.NewRequest("POST", "/users", mockPayload)
		w := httptest.NewRecorder()
		uh.Create(w, r)
		resp := w.Result()

		if resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected %d response", http.StatusInternalServerError)
		}
	})

	t.Run("expect POST /users to return 500 on find internal error", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.createImpl = func(user *models.User) (*models.User, error) {
			return nil, errors.New("repo error")
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("POST", "/users", mockPayload)
		w := httptest.NewRecorder()
		uh.Create(w, r)
		resp := w.Result()

		if resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected %d response", http.StatusInternalServerError)
		}
	})
}