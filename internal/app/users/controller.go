package users

import (
	"encoding/json"
	"net/http"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/httphelper"
)

// Controller ...
type Controller struct {
	usersService ServiceI
}

// NewController ...
func NewController(
	usersService ServiceI,
) *Controller {
	return &Controller{
		usersService: usersService,
	}
}

// HandleUsersCreate ...
func (c *Controller) HandleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := &CreateDTO{}

		if err := json.NewDecoder(r.Body).Decode(params); err != nil {
			httphelper.RespondWithError(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := c.usersService.Create(params)

		if err != nil {
			httphelper.RespondWithError(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		httphelper.Respond(w, r, http.StatusCreated, u)
	}
}
