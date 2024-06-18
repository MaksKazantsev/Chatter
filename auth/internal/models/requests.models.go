package models

type RegReq struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Refresh  string `json:"-"`
}

type LogReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Refresh  string `json:"-"`
}

type ResReq struct {
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
	Token       string `json:"-"`
}
