package postgres

import (
	"context"
	"fmt"
	"head_app/models"
	log "head_app/pkg/logger"
	repoi "head_app/storage/repoI"

	"github.com/jackc/pgx/v5"
	"github.com/saidamir98/udevs_pkg/logger"
)

type commonRepo struct {
	db  *pgx.Conn
	log log.Log
}

func NewCommonRepo(db *pgx.Conn, log log.Log) repoi.CommonRepoI {
	return &commonRepo{db, log}
}

func (c *commonRepo) CheckIsExists(ctx context.Context, req *models.Common) (bool, error) {

	var isExists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s = '%s')", req.TableName, req.ColumnName, req.ExpValue)

	err := c.db.QueryRow(ctx, query).Scan(&isExists)
	if err != nil {
		c.log.Error("error on checking is exists", logger.Error(err))
		return false, err
	}

	return isExists, nil
}
