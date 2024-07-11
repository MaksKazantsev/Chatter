package models

import "time"

type Post struct {
	PostID          string    `db:"postid"`
	PostAuthorID    string    `db:"userid"`
	PostTitle       string    `db:"posttitle"`
	PostDescription string    `db:"postdesc"`
	PostFile        string    `db:"postfile"`
	PostLikesAmount int32     `db:"likesamount"`
	CreatedAt       time.Time `db:"createdat"`
	Comments        []Comment `db:"-"`
	Likes           []Like    `db:"-"`
}
type Comment struct {
	PostID    string       `db:"postid"`
	UserID    string       `db:"userid"`
	CommentID string       `db:"commentid"`
	CreatedAt time.Time    `db:"createdat"`
	ValueDb   string       `db:"val"`
	Value     CommentValue `db:"-"`
}

type CommentValue struct {
	TextValue string
	File      string
}

type Like struct {
	PostID string `db:"postid"`
	UserID string `db:"userid"`
}
