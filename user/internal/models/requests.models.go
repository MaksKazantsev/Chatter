package models

type RegReq struct {
	UUID     string
	Username string
	Password string
	Email    string
	Refresh  string
}

type LogReq struct {
	Email    string
	Password string
	Refresh  string
}

type VerifyCodeReq struct {
	Code  string
	Email string
	Type  string
}
