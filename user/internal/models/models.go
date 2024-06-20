package models

type User struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Refresh  string `json:"refresh"`
}
