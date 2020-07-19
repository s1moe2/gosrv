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
			t.Errorf("expected %d response", http.StatusOK)
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
			t.Errorf("expected %d response", http.StatusOK)
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
			t.Errorf("expected %d response", http.StatusInternalServerError)
		}
	})
}