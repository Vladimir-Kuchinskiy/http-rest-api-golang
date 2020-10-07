package entity

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.com",
		Password: "Passw0rd",
	}
}
