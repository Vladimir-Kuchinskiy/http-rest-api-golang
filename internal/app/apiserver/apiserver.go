package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/adapters/database"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/authentication"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/users"
)

// Start ...
func Start(config *Config) error {
	db, err := database.NewDB(config.DatabaseURL)

	if err != nil {
		return err
	}

	defer db.Close()

	srv := newServer(buildDeps(db))

	configureRouter(srv)

	return http.ListenAndServe(config.BindAddr, srv)
}

func buildDeps(db *sql.DB) *deps {
	usersRepository := users.NewRepository(db)

	usersService := users.NewService(usersRepository)
	authenticationService := authentication.NewService()

	return &deps{
		usersController: users.NewController(usersService),
		authenticationController: authentication.NewController(
			authenticationService,
			usersRepository,
		),
	}
}
