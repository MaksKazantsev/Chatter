package models

import "time"

type Post struct {
	PostID          string
	PostAuthorID    string
	PostTitle       string
	PostDescription string
	PostFile        string
	PostLikesAmount int32
	CreatedAt       time.Time
}
type Comment struct {
	PostID    string
	UserID    string
	CommentID string
	Value     CommentValue
}

type CommentValue struct {
	TextValue string
	File      string
}

type Like struct {
	PostID string
	UserID string
}
