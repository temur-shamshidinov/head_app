package models

import (
	"time"

	"github.com/google/uuid"
)

type SubCategory struct {
	SubCategoryID uuid.UUID `json:"sub_category_id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"created_at"`
	CategoryID    uuid.UUID `json:"category_id"`
}

type CreateSubCategoryReq struct {
	Name       string    `json:"name"`
	CategoryID uuid.UUID `json:"category_id"`
}

type UpdateSubCategoryReq struct {
	SubCategoryID uuid.UUID `json:"sub_category_id"`
	Name          string    `json:"name"`
	CategoryID    uuid.UUID `json:"category_id"`
}

type GetSubCategoriesLisResp struct {
	SubCategories []*SubCategory `json:"sub_categories"`
	Count         int            `json:"count"`
}
