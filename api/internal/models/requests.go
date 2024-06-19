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

type VerifyCodeReq struct {
	Email string `json:"email"`
	Code  string `json:"-"`
	Type  string `json:"-"`
}

type RecoveryReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SendReq struct {
	Email string `json:"email"`
}
