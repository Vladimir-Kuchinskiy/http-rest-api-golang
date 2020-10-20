package apiserver

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/httphelper"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/users"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/authentication"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/store"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

const (
	sessionName = "httpgolangapp"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

type server struct {
	router                   *mux.Router
	logger                   *logrus.Logger
	store                    store.Store
	sessionStore             sessions.Store
	authenticationController *authentication.Controller
	usersController          *users.Controller
}

func newServer(
	store store.Store,
	sessionStore sessions.Store,
	db *sql.DB,
) *server {
	usersRepository := users.NewRepository(db)

	usersService := users.NewService(usersRepository)
	authenticationService := authentication.NewService()

	usersController := users.NewController(usersService)
	authenticationController := authentication.NewController(
		authenticationService,
		usersRepository,
	)

	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,

		usersController:          usersController,
		authenticationController: authenticationController,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.Handle("/users", s.usersController.HandleUsersCreate()).Methods("POST")
	s.router.Handle("/auth", s.authenticationController.HandleAuth()).Methods("POST")

	s.router.Handle(
		"/whoami",
		s.authenticationController.AuthenticateUser(
			s.authenticationController.HandleWhoami(),
		),
	).Methods("GET")
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
