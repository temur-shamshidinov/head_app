package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	CommentID uuid.UUID `json:"comment_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ArticleID uuid.UUID `json:"article_id"`
	ViewerID  uuid.UUID `json:"viewer_id"`
}

type CreateCommentReq struct {
	Content   string    `json:"content"`
	ArticleID uuid.UUID `json:"article_id"`
	ViewerID  uuid.UUID `json:"viewer_id"`
}

type UpdateCommentReq struct {
	CommentID uuid.UUID `json:"comment_id"`
	Content   string    `json:"content"`
	ArticleID uuid.UUID `json:"article_id"`
}

type GetCommentListResp struct {
	Comments []*Comment `json:"comments"`
	Count    int        `json:"count"`
}
