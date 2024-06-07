package models

type SignupReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetReq struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
