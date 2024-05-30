package models

type User struct {
	Handle   string `json:"handle"`
	Username string `json:"name"`
	Password string `json:"pass"`
}
