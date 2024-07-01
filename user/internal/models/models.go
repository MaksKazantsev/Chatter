package models

import "time"

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

type Friend struct {
	FriendID string `db:"friendid"`
	Avatar   string `db:"avatar"`
	Username string `db:"username"`
}

type UserProfile struct {
	UUID       string    `db:"uuid"`
	Avatar     string    `db:"avatar"`
	Username   string    `db:"username"`
	Birthday   time.Time `db:"birthday"`
	Bio        string    `db:"bio"`
	LastOnline time.Time `db:"lastonline"`
}

type GetUserProfile struct {
	Avatar     string    `db:"avatar"`
	Username   string    `db:"username"`
	Birthday   time.Time `db:"birthday"`
	Bio        string    `db:"bio"`
	LastOnline time.Time `db:"lastonline"`
}

type FsReq struct {
	ReqID    string `db:"requestid"`
	Avatar   string `db:"avatar"`
	Username string `db:"username"`
}
