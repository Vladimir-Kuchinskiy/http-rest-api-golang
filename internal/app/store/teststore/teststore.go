package teststore

import (
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/entity"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*entity.User),
	}

	return s.userRepository
}
