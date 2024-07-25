package storage

import (
	log "head_app/pkg/logger"
	"head_app/storage/postgres"
	repoi "head_app/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type StorageI interface {
	GetCategoryRepo() repoi.CategoryRepoI
	GetOwnerRepo() repoi.OwnerRepoI
	GetCommonRepo() repoi.CommonRepoI
	GetViewerRepo() repoi.ViewerRepoI
	GetArticleRepo() repoi.ArticleRepoI
	GetSubCategoryRepo() repoi.SubCategoryRepoI
}

type storage struct {
	articleRepo  repoi.ArticleRepoI
	categoryRepo repoi.CategoryRepoI
	ownerRepo    repoi.OwnerRepoI
	commonRepo   repoi.CommonRepoI
	viewerRepo   repoi.ViewerRepoI
	subcategoryRepo repoi.SubCategoryRepoI
}

func NewStorage(db *pgx.Conn, log log.Log) StorageI {
	return &storage{
		ownerRepo:       postgres.NewOwnerRepo(db, log),
		commonRepo:      postgres.NewCommonRepo(db, log),
		viewerRepo:      postgres.NewViewerRepo(db, log),
		articleRepo:     postgres.NewArticleRepo(db, log),
		categoryRepo:    postgres.NewCategoryRepo(db, log),
		subcategoryRepo: postgres.NewSubCategoryRepo(db, log),
	}
}

func (s *storage) GetCategoryRepo() repoi.CategoryRepoI {
	return s.categoryRepo
}

func (s *storage) GetOwnerRepo() repoi.OwnerRepoI {
	return s.ownerRepo
}

func (s *storage) GetCommonRepo() repoi.CommonRepoI {
	return s.commonRepo
}

func (s *storage) GetViewerRepo() repoi.ViewerRepoI {
	return s.viewerRepo
}

func (s *storage) GetArticleRepo() repoi.ArticleRepoI {
	return s.articleRepo
}

func (s *storage) GetSubCategoryRepo() repoi.SubCategoryRepoI {
	return s.subcategoryRepo
}
