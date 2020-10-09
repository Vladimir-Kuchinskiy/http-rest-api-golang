package store

import (
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/entity"
)

// UserRepository ...
type UserRepository interface {
	Create(*entity.User) error
	FindByEmail(string) (*entity.User, error)
}
