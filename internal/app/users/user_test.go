package users_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/users"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		desc    string
		u       func() *users.User
		isValid bool
	}{
		{
			desc: "When user is valid",
			u: func() *users.User {
				return users.TestUser(t)
			},
			isValid: true,
		},
		{
			desc: "When user has EnrcyptedPassword",
			u: func() *users.User {
				u := users.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encrypted-password"

				return u
			},
			isValid: true,
		},
		{
			desc: "When user has NO Email",
			u: func() *users.User {
				u := users.TestUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			desc: "When user has invalid Email",
			u: func() *users.User {
				u := users.TestUser(t)
				u.Email = "invalid"

				return u
			},
			isValid: false,
		},
		{
			desc: "When user has NO Password",
			u: func() *users.User {
				u := users.TestUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			desc: "When user has invalid Password",
			u: func() *users.User {
				u := users.TestUser(t)
				u.Password = "123"

				return u
			},
			isValid: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.isValid {
				assert.NoError(t, tC.u().Validate())
			} else {
				assert.Error(t, tC.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := users.TestUser(t)

	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
