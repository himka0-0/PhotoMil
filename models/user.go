package models

type User struct {
	Login    string `json:"email"`
	Password string `json:"password"`
}
