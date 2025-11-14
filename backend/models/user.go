package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // Don't send password in JSON
	Name     string `json:"name"`
}
