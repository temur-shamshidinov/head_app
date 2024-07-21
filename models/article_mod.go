package models

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ArticleID     uuid.UUID  `json:"article_id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	CategoryID    uuid.UUID  `json:"category_id"`
	SubCategoryID uuid.UUID  `json:"sub_category_id"`
}

