package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	CategoryID uuid.UUID `json:"category_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreatedCategoryReq struct {
	Name string `json:"name"`
}

type GetCategoriesListResp struct {
	Categories []*Category
	Count      int32
}


