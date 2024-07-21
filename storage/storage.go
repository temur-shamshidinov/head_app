package storage

import (
	log "head_app/pkg/logger"
	"head_app/storage/postgres"
	repoi "head_app/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type StorageI interface {
	GetContentRepo() repoi.ContentRepoI
	GetOwnerRepo() repoi.OwnerRepoI
	GetCommonRepo() repoi.CommonRepoI
}

type storage struct {
	contentRepo repoi.ContentRepoI
	ownerRepo   repoi.OwnerRepoI
	commonRepo  repoi.CommonRepoI
}

func NewStorage(db *pgx.Conn, log log.Log) StorageI {
	return &storage{
		contentRepo: postgres.NewContentRepo(db, log),
		ownerRepo:   postgres.NewOwnerRepo(db, log),
		commonRepo:  postgres.NewCommonRepo(db,log),

	}
}

func (s *storage) GetContentRepo() repoi.ContentRepoI {
	return s.contentRepo
}

func (s *storage) GetOwnerRepo() repoi.OwnerRepoI {
	return s.ownerRepo
}

func (s *storage) GetCommonRepo() repoi.CommonRepoI{
	return s.commonRepo
}