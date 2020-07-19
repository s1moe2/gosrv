package handlers

import (
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
			t.Errorf("expected %d response, got %d", http.StatusOK, resp.StatusCode)
		}

		if resp.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected 'application/json', got '%s'", resp.Header.Get("Content-Type"))
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error("failed to read response body")
		}

		var users []*models.User
		json.Unmarshal(body, &users)

		if len(users) != len(mock.mockUsers) {
			t.Error("wrong user set")
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
			t.Errorf("expected %d response, got %d", http.StatusOK, resp.StatusCode)
		}

		if resp.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected 'application/json', got '%s'", resp.Header.Get("Content-Type"))
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error("failed to read response body")
		}

		var users []*models.User
		json.Unmarshal(body, &users)

		if len(users) != 0 {
			t.Error("wrong user set")
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
			t.Errorf("expected %d response, got %d", http.StatusInternalServerError, resp.StatusCode)
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
			t.Errorf("expected %d response, got %d", http.StatusOK, resp.StatusCode)
		}

		if resp.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected 'application/json', got '%s'", resp.Header.Get("Content-Type"))
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error("failed to read response body")
		}

		var user *models.User
		json.Unmarshal(body, &user)

		if user != mock.mockUsers[0] {
			t.Error("wrong user returned")
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
			t.Errorf("expected %d response, got %d", http.StatusNotFound, resp.StatusCode)
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
			t.Errorf("expected %d response", http.StatusInternalServerError)
		}
	})
}