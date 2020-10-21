package apiserver

import "github.com/gorilla/handlers"

func configureRouter(s *server) {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.Handle("/users", s.usersController.HandleUsersCreate()).Methods("POST")
	s.router.Handle("/auth", s.authenticationController.HandleAuth()).Methods("POST")
	s.router.Handle("/unauth", s.authenticationController.HandleUnauth()).Methods("POST")

	s.router.Handle(
		"/whoami",
		s.authenticationController.AuthenticateUser(
			s.authenticationController.HandleWhoami(),
		),
	).Methods("GET")
}
