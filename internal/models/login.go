package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// In Memory DB

var users []User

// Model Actions

func GenerateUsers() {
	user := User{
		ID:       "1",
		Username: "user-test",
		Name:     "Test User",
		Email:    "test@test.com",
		Password: "Test1234",
	}
	securePassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic("Error del servidor al generar el token de prueba")
	}

	user.Password = string(securePassword)
	users = append(users, user)
}

func SearchUser(user *User) (*User, bool) {
	for _, u := range users {
		if u.Username == user.Username {
			if samePassword(u.Password, user.Password) {
				return &u, true
			}
		}
	}
	return nil, false
}

// Helpers

func samePassword(registeredPassword string, loginPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(registeredPassword), []byte(loginPassword)) == nil
}
