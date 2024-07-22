package repoi

import (
	"context"
	"head_app/models"
)

type ContentRepoI interface {
	// Category
	CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	GetCategories(ctx context.Context, page, limit int32) (*models.GetCategoriesListResp, error)
	GetCategory(ctx context.Context, id string) (*models.Category, error)
	UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	DeleteCategory(ctx context.Context, id string) error

	// SubCategory
	CreateSubCategory(ctx context.Context, category *models.SubCategory) (*models.SubCategory, error)
	GetSubCategories(ctx context.Context, page, limit int32) ([]*models.SubCategory, error)
	GetSubCategory(ctx context.Context, id string) (*models.SubCategory, error)
	UpdateSubCategory(ctx context.Context, category *models.SubCategory) (*models.SubCategory, error)
	DeleteSubCategory(ctx context.Context, id string) error

	// Article

	CreateArticle(ctx context.Context, article *models.Article) (*models.Article, error)
	GetArticles(ctx context.Context, page, limit int32) ([]*models.Article, error)
	GetArticle(ctx context.Context, id string) (*models.Article, error)
	UpdateArticle(ctx context.Context, article *models.Article) (*models.Article, error)
	DeleteArticle(ctx context.Context, id string) error
}
  
// owner

type OwnerRepoI interface{
	Login(ctx context.Context, login *models.LoginOwner) (*models.LoginOwner, error)

}

type CommonRepoI interface {
	CheckIsExists(ctx context.Context, req *models.Common) (bool, error)
}

type ViewerRepoI interface {
	CreateViewer(ctx context.Context,viewer *models.Viewer) (*models.Claim, error) 
	LogIn(ctx context.Context,login *models.LogInViewer) (*models.Claim, error) 
}