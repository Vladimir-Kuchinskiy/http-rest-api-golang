package users_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/adapters/database"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/users"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := database.TestDB(t)
	defer teardown("users")

	userRepository := users.NewRepository(db)
	u := users.TestUser(t)
	err := userRepository.Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u.ID)
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := database.TestDB(t)
	defer teardown("users")

	userRepository := users.NewRepository(db)
	email := "user@example.com"
	_, err := userRepository.FindByEmail(email)
	assert.EqualError(t, err, database.ErrRecordNotFound.Error())

	u := users.TestUser(t)
	u.Email = email
	userRepository.Create(u)

	u, err = userRepository.FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.Email, email)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := database.TestDB(t)
	defer teardown("users")

	userRepository := users.NewRepository(db)

	u := users.TestUser(t)
	userRepository.Create(u)

	u2, err := userRepository.Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
	assert.Equal(t, u.ID, u2.ID)
}
