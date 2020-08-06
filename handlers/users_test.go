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
			mockUsers := []*models.User{
				&models.User{
					ID:    "1",
					Name:  "user1",
					Email: "user1@eml.com",
				},
			}
			return mockUsers, nil
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodGet, "/users", uh.Get)
		router.ServeHTTP(w, r)

		resp := w.Result()

		assertStatusCode(t, resp, http.StatusOK)
		assertContentType(t, resp)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var user []*models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			t.Fatal("failed to parse response body")
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

		assertStatusCode(t, resp, http.StatusOK)
		assertContentType(t, resp)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var users []*models.User
		err = json.Unmarshal(body, &users)
		if err != nil {
			t.Fatal("failed to parse response body")
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

		assertStatusCode(t, resp, http.StatusInternalServerError)
	})
}

func TestUsersHandler_GetByID(t *testing.T) {
	t.Run("expect GET /users/{id} to return 200", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.findByIDImpl = func(ID string) (*models.User, error) {
			return &models.User{
				ID:    "1",
				Name:  "user1",
				Email: "user1@eml.com",
			}, nil
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("GET", "/users/1", nil)
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodGet, "/users/{id}", uh.GetByID)
		router.ServeHTTP(w, r)

		resp := w.Result()

		assertStatusCode(t, resp, http.StatusOK)
		assertContentType(t, resp)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var user *models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			t.Fatal("failed to parse response body")
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

		assertStatusCode(t, resp, http.StatusNotFound)
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

		assertStatusCode(t, resp, http.StatusInternalServerError)
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
			return user, nil
		}
		uh := NewUsersHandler(mock)

		body, _ := json.Marshal(mockPayload)
		r := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
		r.Header.Set("Content-Type", "application/json; charset=utf-8")
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodPost, "/users", uh.Create)
		router.ServeHTTP(w, r)
		resp := w.Result()

		assertStatusCode(t, resp, http.StatusCreated)
		assertContentType(t, resp)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var user *models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			t.Fatal("failed to parse response body")
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

		assertStatusCode(t, resp, http.StatusBadRequest)
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

		assertStatusCode(t, resp, http.StatusInternalServerError)
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

		assertStatusCode(t, resp, http.StatusInternalServerError)
	})
}

func TestUsersHandler_Update(t *testing.T) {
	mockPayload := map[string]interface{}{
		"email": "johndoe@gosrv.com",
		"name":  "John Doe",
	}

	t.Run("expect PUT /users/{id} to return 200", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.updateImpl = func(user *models.User) (*models.User, error) {
			return user, nil
		}
		uh := NewUsersHandler(mock)

		body, _ := json.Marshal(mockPayload)
		r := httptest.NewRequest("PUT", "/users/1", bytes.NewReader(body))
		r.Header.Set("Content-Type", "application/json; charset=utf-8")
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodPut, "/users/{id}", uh.Update)
		router.ServeHTTP(w, r)
		resp := w.Result()

		assertStatusCode(t, resp, http.StatusOK)
		assertContentType(t, resp)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("failed to read response body")
		}

		var user *models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			t.Fatal("failed to parse response body")
		}
	})

	t.Run("expect PUT /users/{id} to return 404 when the user does not exist", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.updateImpl = func(user *models.User) (*models.User, error) {
			return nil, nil
		}
		uh := NewUsersHandler(mock)

		body, _ := json.Marshal(mockPayload)
		r := httptest.NewRequest("PUT", "/users/1", bytes.NewReader(body))
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodPut, "/users/{id}", uh.Update)
		router.ServeHTTP(w, r)
		resp := w.Result()

		assertStatusCode(t, resp, http.StatusNotFound)
	})

	t.Run("expect PUT /users/{id} to return 500 on internal error", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.updateImpl = func(user *models.User) (*models.User, error) {
			return nil, errors.New("repo error")
		}
		uh := NewUsersHandler(mock)

		body, _ := json.Marshal(mockPayload)
		r := httptest.NewRequest("PUT", "/users/1", bytes.NewReader(body))
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodPut, "/users/{id}", uh.Update)
		router.ServeHTTP(w, r)
		resp := w.Result()

		assertStatusCode(t, resp, http.StatusInternalServerError)
	})
}

func TestUsersHandler_Delete(t *testing.T) {
	t.Run("expect DELETE /users/{id} to return 204", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.deleteImpl = func(ID string) (string, error) {
			return ID, nil
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("DELETE", "/users/1", nil)
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodDelete, "/users/{id}", uh.Delete)
		router.ServeHTTP(w, r)

		resp := w.Result()

		assertStatusCode(t, resp, http.StatusNoContent)
		assertContentType(t, resp)
	})

	t.Run("expect DELETE /users/{id} to return 404 when user does not exist", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.deleteImpl = func(ID string) (string, error) {
			return "", nil
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("DELETE", "/users/1", nil)
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodDelete, "/users/{id}", uh.Delete)
		router.ServeHTTP(w, r)

		resp := w.Result()

		assertStatusCode(t, resp, http.StatusNotFound)
	})

	t.Run("expect DELETE /users/{id} to return 500 on internal error", func(t *testing.T) {
		mock := newUserRepoMockDefault()
		mock.deleteImpl = func(ID string) (string, error) {
			return "", errors.New("repo error")
		}
		uh := NewUsersHandler(mock)

		r := httptest.NewRequest("DELETE", "/users/1", nil)
		w := httptest.NewRecorder()
		router := prepareRouter(http.MethodDelete, "/users/{id}", uh.Delete)
		router.ServeHTTP(w, r)

		resp := w.Result()

		assertStatusCode(t, resp, http.StatusInternalServerError)
	})
}
