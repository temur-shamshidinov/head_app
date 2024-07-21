package postgres

import (
	"context"
	"head_app/models"
	log "head_app/pkg/logger"
	repoi "head_app/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type ownerRepo struct {
	db *pgx.Conn
	log log.Log
}

func NewOwnerRepo(db *pgx.Conn,log log.Log) repoi.OwnerRepoI{
	return &ownerRepo{db,log}
}

func (o *ownerRepo) Login(ctx context.Context, login *models.LoginOwner) (*models.LoginOwner, error) {
	return nil, nil
}