package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaksKazantsev/Chatter/posts/internal/log"
	"github.com/MaksKazantsev/Chatter/posts/internal/models"
	"github.com/MaksKazantsev/Chatter/posts/internal/utils"
	"strings"
)

func (p *Postgres) CreatePost(ctx context.Context, req models.PostReq) error {
	q := `INSERT INTO posts(userid,postid,posttitle,postdesc,postfile) VALUES($1,$2,$3,$4,$5)`
	_, err := p.Exec(q, req.PostAuthorID, req.PostID, req.PostTitle, req.PostDescription, req.PostFile)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}

func (p *Postgres) DeletePost(ctx context.Context, userID, postID string) error {
	q := `DELETE FROM posts WHERE userid = $1 AND postid = $2`
	res, err := p.Exec(q, userID, postID)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return utils.NewError("post with this id and userid does not exist", utils.ErrNotFound)
	}

	q = `DELETE FROM comments WHERE postid = $1`
	_, err = p.Exec(q, postID)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	q = `DELETE FROM likes WHERE postid = $1`
	_, err = p.Exec(q, postID)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}

func (p *Postgres) EditPost(ctx context.Context, req models.PostReq) error {
	q := `SELECT * FROM posts WHERE userid = $1 AND postid = $2`
	_, err := p.Exec(q, req.PostAuthorID, req.PostID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("post does not exist", utils.ErrNotFound)
		}
	}
	if req.PostTitle != "" {
		q = `UPDATE posts SET posttitle = $1 WHERE userid = $2 AND postid = $3`
		_, err = p.Exec(q, req.PostTitle, req.PostAuthorID, req.PostID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}
	if req.PostDescription != "" {
		q = `UPDATE posts SET postdesc = $1 WHERE userid = $2 AND postid = $3`
		_, err = p.Exec(q, req.PostDescription, req.PostAuthorID, req.PostID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}
	if req.PostFile != "" {
		q = `UPDATE posts SET postfile = $1 WHERE userid = $2 AND postid = $3`
		_, err = p.Exec(q, req.PostFile, req.PostAuthorID, req.PostID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}
	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}
func (p *Postgres) LikePost(ctx context.Context, userID, postID, likeID string) error {
	q := `SELECT * FROM posts WHERE postid = $1`
	res, err := p.Exec(q, postID)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	amount, _ := res.RowsAffected()
	if amount == 0 {
		return utils.NewError("post does not exist", utils.ErrNotFound)
	}

	q = `INSERT INTO likes(userid,postid,likeid) VALUES($1,$2,$3)`
	_, err = p.Exec(q, userID, postID, likeID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return utils.NewError("user can like post only once", utils.ErrBadRequest)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	q = `UPDATE posts SET likesamount = likesamount + 1 WHERE postid = $1`
	res, err = p.Exec(q, postID)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return utils.NewError("post with this id and userid does not exist", utils.ErrNotFound)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}

func (p *Postgres) LeaveComment(ctx context.Context, comment models.Comment) error {
	q := `SELECT * FROM posts WHERE userid = $1 AND postid = $2`
	_, err := p.Exec(q, comment.UserID, comment.PostID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("post does not exist", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	commentVal := comment.Value.TextValue + ":" + comment.Value.File
	q = `INSERT INTO comments(userid,postid,commentid,val) VALUES($1,$2,$3,$4)`
	_, err = p.Exec(q, comment.UserID, comment.PostID, comment.CommentID, commentVal)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}

func (p *Postgres) GetUserPosts(ctx context.Context, userID string) ([]models.Post, error) {
	var res []models.Post

	q := `SELECT * FROM posts WHERE userid = $1 ORDER BY createdat DESC`
	rows, err := p.Queryx(q, userID)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.ErrInternal)
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var post models.Post
		err := rows.StructScan(&post)
		if err != nil {
			return nil, utils.NewError(err.Error(), utils.ErrInternal)
		}

		likesQuery := `SELECT postid, userid FROM likes WHERE postid = $1`
		likesRows, err := p.Queryx(likesQuery, post.PostID)
		if err != nil {
			return nil, utils.NewError(err.Error(), utils.ErrInternal)
		}

		for likesRows.Next() {
			var like models.Like
			err := likesRows.StructScan(&like)
			if err != nil {
				return nil, utils.NewError(err.Error(), utils.ErrInternal)
			}
			post.Likes = append(post.Likes, like)
		}

		commentsQuery := `SELECT * FROM comments WHERE postid = $1`
		commentsRows, err := p.Queryx(commentsQuery, post.PostID)
		if err != nil {
			return nil, utils.NewError(err.Error(), utils.ErrInternal)
		}

		for commentsRows.Next() {
			var comment models.Comment
			err := commentsRows.StructScan(&comment)
			if err != nil {
				return nil, utils.NewError(err.Error(), utils.ErrInternal)
			}
			val := strings.Split(comment.ValueDb, ":")
			if len(val) < 2 {
				val = append(val, "")
			}
			comment.Value = models.CommentValue{
				TextValue: val[0],
				File:      val[1],
			}
			post.Comments = append(post.Comments, comment)

		}

		res = append(res, post)
	}

	if err := rows.Err(); err != nil {
		return nil, utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return res, nil
}

func (p *Postgres) DeleteComment(ctx context.Context, commentID, userID string) error {
	q := `DELETE FROM comments WHERE commentid = $1 AND userid = $2`
	res, err := p.Exec(q, commentID, userID)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	amount, _ := res.RowsAffected()
	if amount == 0 {
		return utils.NewError("comment does not exist", utils.ErrNotFound)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}

func (p *Postgres) UnlikePost(ctx context.Context, userID, postID string) error {
	q := `DELETE FROM likes WHERE userid = $1 AND postid = $2`
	res, err := p.Exec(q, userID, postID)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	amount, _ := res.RowsAffected()
	if amount == 0 {
		return utils.NewError("like does not exist", utils.ErrNotFound)
	}

	q = `UPDATE posts SET likesamount = likesamount - 1 WHERE postid = $1`
	res, err = p.Exec(q, postID)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	amount, _ = res.RowsAffected()
	if amount == 0 {
		return utils.NewError("like does not exist", utils.ErrNotFound)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}

func (p *Postgres) EditComment(ctx context.Context, dbValue, commentID, userID string) error {
	q := `UPDATE comments SET val = $1 WHERE commentid = $2 AND userid = $3`
	res, err := p.Exec(q, dbValue, commentID, userID)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	amount, _ := res.RowsAffected()
	if amount == 0 {
		return utils.NewError("comment does not exist", utils.ErrNotFound)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}
