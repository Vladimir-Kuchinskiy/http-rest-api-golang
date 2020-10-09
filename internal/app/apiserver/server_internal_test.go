package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/entity"

	"github.com/stretchr/testify/assert"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/store/teststore"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		desc         string
		payload      interface{}
		expectedCode int
	}{
		{
			desc: "when params valid",
			payload: map[string]string{
				"email":    "user@example.com",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			desc:         "when invalid request payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			desc: "when invalid params",
			payload: map[string]string{
				"email":    "invalid",
				"password": "1",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tC.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tC.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := entity.TestUser(t)
	store := teststore.New()

	s := newServer(store, sessions.NewCookieStore([]byte("secret")))

	store.User().Create(u)

	testCases := []struct {
		desc         string
		payload      interface{}
		expectedCode int
	}{
		{
			desc: "when user exists and email and password match",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			desc:         "when payload is invalid",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			desc: "when user with given email does not exist",
			payload: map[string]string{
				"email":    "unknown@example.com",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			desc: "when user with given email exists and passwords doesn't match",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid-password",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tC.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tC.expectedCode, rec.Code)
		})
	}
}
