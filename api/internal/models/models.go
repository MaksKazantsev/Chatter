package models

import "time"

type Message struct {
	ChatID     string    `json:"-"`
	SenderID   string    `json:"-"`
	ReceiverID string    `json:"receiverID"`
	MessageID  string    `json:"-"`
	Value      string    `json:"value"`
	SentAt     time.Time `json:"-"`
	Token      string    `json:"-"`
}

type Friend struct {
	FriendID string `json:"friendID"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
}

type FsReq struct {
	ReqID    string `json:"requestID"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
}

type UserProfileReq struct {
	Avatar   string     `json:"avatar"`
	Birthday *time.Time `json:"birthday" example:"2024-03-11T14:30:00Z"`
	Bio      string     `json:"bio"`
	Token    string     `json:"-"`
}

type UserProfile struct {
	Avatar     string    `json:"avatar"`
	Bio        string    `json:"bio"`
	Birthday   time.Time `json:"birthday"`
	LastOnline time.Time `json:"lastOnline"`
	Username   string    `json:"username"`
}

type ProfileAvatar struct {
	Avatar string `json:"avatar"`
}

type Post struct {
	PostID          string    `json:"postID"`
	PostAuthorID    string    `json:"postAuthorID"`
	PostTitle       string    `json:"postTitle"`
	PostDescription string    `json:"postDescription"`
	PostFile        string    `json:"postFile"`
	PostLikesAmount int32     `json:"postLikesAmount"`
	CreatedAt       time.Time `json:"createdAt"`
	Comments        []Comment `json:"comments"`
	Likes           []Like    `json:"likes"`
}

type Comment struct {
	PostID    string       `json:"postID"`
	UserID    string       `json:"userID"`
	CommentID string       `json:"commentID"`
	CreatedAt time.Time    `json:"createdAt"`
	ValueDb   string       `json:"-"`
	Value     CommentValue `json:"value"`
}

type CommentValue struct {
	TextValue string `json:"textValue"`
	File      string `json:"file"`
}

type Like struct {
	PostID string `json:"postID"`
	UserID string `json:"userID"`
}
