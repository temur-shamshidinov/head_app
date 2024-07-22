package postgres

import (
	"context"
	"errors"
	"head_app/models"
	"head_app/pkg/helpers"
	log "head_app/pkg/logger"
	repoi "head_app/storage/repoI"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/saidamir98/udevs_pkg/logger"
)

type viewerRepo struct {
	db  *pgx.Conn
	log log.Log
}

func NewViewerRepo(db *pgx.Conn, log log.Log) repoi.ViewerRepoI {
	return &viewerRepo{
		db:  db,
		log: log,
	}
}

func (v *viewerRepo) CreateViewer(ctx context.Context, viewer *models.Viewer) (*models.Claim, error) {
	query := `
		INSERT INTO viewers (
			viewer_id,
			fullname,
			username,
			gmail,    
			password 
		) VALUES($1,$2,$3,$4,$5)
	`

	_, err := v.db.Exec(
		ctx, query,
		viewer.ViewerID.String(),
		viewer.Fullname,
		viewer.Username,
		viewer.Gmail,
		viewer.Password,
	)
	if err != nil {
		v.log.Error("error on creating new viewer", logger.Error(err))
		return nil, err
	}

	return v.GetClaims(ctx, viewer.ViewerID.String())
}
func (v *viewerRepo) LogIn(ctx context.Context, login *models.LogInViewer) (*models.Claim, error) {

	var (
		viewerId        uuid.UUID
		gmail, passwordHash string
	)
	query := `
		SELECT 
			viewer_id,
			gmail,
			password 
		FROM 
			viewers 
		WHERE 
			username = $1	
	`
	err := v.db.QueryRow(ctx, query, login.Username).Scan(&viewerId,&gmail,&passwordHash)
	if err != nil {
		return nil, err
	}

	if !helpers.CompareHashAndPassword(passwordHash,login.Password){
		return nil, errors.New("password in incorrect")
	}

	return &models.Claim{
		UserID: viewerId.String(),
		UserRole: "viewer",
	}, nil
	
}

func (v *viewerRepo) GetClaims(ctx context.Context, id string) (*models.Claim, error) {

	var viewerId uuid.UUID

	query := `
		SELECT viewer_id FROM viewers WHERE viewer_id = $1 
	`
	v.db.QueryRow(ctx, query, id).Scan(&viewerId)

	return &models.Claim{
		UserID:   viewerId.String(),
		UserRole: "viewer",
	}, nil

}
