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

type FriendShipReq struct {
	Token    string `json:"-"`
	Receiver string `json:"receiver"`
}
type RefuseFriendShipReq struct {
	Token  string `json:"-"`
	Sender string `json:"sender"`
}

type GetHistoryReq struct {
	ChatID string `json:"chatID"`
	Token  string `json:"token"`
}

type UploadReq struct {
	PhotoID string `json:"photoID"`
	Photo   []byte `json:"photo"`
	Token   string `json:"token"`
}
