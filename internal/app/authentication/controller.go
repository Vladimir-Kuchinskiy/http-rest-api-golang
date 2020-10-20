package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/httphelper"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/users"
)

const (
	authTokenKey = "X-Auth-Token"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

// Controller ...
type Controller struct {
	authenticationService ServiceI
	usersRepository       users.RepositoryI
}

// NewController ...
func NewController(
	authenticationService ServiceI,
	usersRepository users.RepositoryI,
) *Controller {
	return &Controller{
		authenticationService: authenticationService,
		usersRepository:       usersRepository,
	}
}

// AuthenticateUser ...
func (c *Controller) AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := extractAuthToken(r)
		if err != nil {
			httphelper.RespondWithError(w, r, http.StatusUnauthorized, err)
			return
		}

		accessDetails, err := c.authenticationService.ExtractTokenMetadata(tokenString)

		if err != nil {
			httphelper.RespondWithError(w, r, http.StatusUnauthorized, err)
			return
		}

		u, err := c.usersRepository.Find(accessDetails.UserID)
		if err != nil {
			httphelper.RespondWithError(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}
		fmt.Println("u:", u)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), httphelper.CtxKeyUser, u)))
	})
}

// HandleAuth ...
func (c *Controller) HandleAuth() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			httphelper.RespondWithError(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := c.usersRepository.FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			httphelper.RespondWithError(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		token, err := c.authenticationService.CreateAuthToken(u.ID)
		if err != nil {
			httphelper.RespondWithError(w, r, http.StatusInternalServerError, err)
			return
		}

		cookie := &http.Cookie{
			Name:     authTokenKey,
			Value:    token,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)

		httphelper.Respond(w, r, http.StatusOK, nil)
	}
}

// HandleWhoami ...
func (c *Controller) HandleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httphelper.Respond(w, r, http.StatusOK, r.Context().Value(httphelper.CtxKeyUser).(*users.User))
	}
}

func extractAuthToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(authTokenKey)

	if err != nil {
		return "", err
	}

	return strings.Split(cookie.String(), "=")[1], nil
}
