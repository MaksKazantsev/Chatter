package models

type PostReq struct {
	PostID          string
	PostAuthorID    string
	PostTitle       string
	PostDescription string
	PostFile        string
}

type EditCommentReq struct {
	CommentID string
	UserID    string
	Value     CommentValue
}
