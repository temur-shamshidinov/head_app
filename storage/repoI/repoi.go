package repoi

import (
	"context"
	"head_app/models"
)

type CategoryRepoI interface {
	// Category
	CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	GetCategories(ctx context.Context, page, limit int32) (*models.GetCategoriesListResp, error)
	GetCategory(ctx context.Context, id string) (*models.Category, error)
	UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	DeleteCategory(ctx context.Context, id string) error
}
type SubCategoryRepoI interface {
	CreateSubCategory(ctx context.Context, subcategory *models.SubCategory) (*models.SubCategory, error)
	GetSubCategories(ctx context.Context, page, limit int32) (*models.GetSubCategoriesLisResp, error)
	GetSubCategory(ctx context.Context, id string) (*models.SubCategory, error)
	UpdateSubCategory(ctx context.Context, subcategory *models.SubCategory) (*models.SubCategory, error)
	DeleteSubCategory(ctx context.Context, id string) error
}
type ArticleRepoI interface {
	CreateArticle(ctx context.Context, article *models.Article) (*models.Article, error)
	GetArticles(ctx context.Context, page, limit int32) (*models.GetArticleListResp, error)
	GetArticle(ctx context.Context, id string) (*models.Article, error)
	UpdateArticle(ctx context.Context, article *models.Article) (*models.Article, error)
	DeleteArticle(ctx context.Context, id string) error
}

// owner

type OwnerRepoI interface {
	Login(ctx context.Context, login *models.LoginOwner) (*models.LoginOwner, error)
}

type CommonRepoI interface {
	CheckIsExists(ctx context.Context, req *models.Common) (bool, error)
}

type ViewerRepoI interface {
	CreateViewer(ctx context.Context, viewer *models.Viewer) (*models.Claim, error)
	LogIn(ctx context.Context, login *models.LogInViewer) (*models.Claim, error)
	//
	// comment
	AddComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	GetComments(ctx context.Context, getListReq *models.GetListReq) (*models.GetCommentListResp, error)
	GetComment(ctx context.Context, id string) (*models.Comment, error)
	UpdateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	DeleteComment(ctx context.Context, id string) error
}
