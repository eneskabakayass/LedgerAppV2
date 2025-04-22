package models

import (
	"errors"
	"regexp"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

func (u *User) Validate() error {
	if u.ID == "" {
		return errors.New("ID is required")
	}
	if u.Username == "" {
		return errors.New("Username is required")
	}
	if !isValidEmail(u.Email) {
		return errors.New("Invalid email format")
	}
	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
