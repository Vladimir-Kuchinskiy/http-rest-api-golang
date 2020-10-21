package apiserver

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/authentication"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/httphelper"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/users"
)

type deps struct {
	authenticationController *authentication.Controller
	usersController          *users.Controller
}

type server struct {
	router *mux.Router
	logger *logrus.Logger
	*deps
}

func newServer(d *deps) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),

		deps: d,
	}

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), httphelper.CtxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(httphelper.CtxKeyRequestID),
		})

		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}
