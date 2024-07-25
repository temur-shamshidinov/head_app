package postgres

import (
	"context"
	"head_app/models"
	log "head_app/pkg/logger"
	repoi "head_app/storage/repoI"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/saidamir98/udevs_pkg/logger"
)

type subcategoryRepo struct {
	db  *pgx.Conn
	log log.Log
}

func NewSubCategoryRepo(db *pgx.Conn, log log.Log) repoi.SubCategoryRepoI {
	return &subcategoryRepo{db, log}
}

// SubCategory
func (s *subcategoryRepo) CreateSubCategory(ctx context.Context, subcategory *models.SubCategory) (*models.SubCategory, error) {
	s.log.Debug("request in CreateSubCategory.")

	subcategory.SubCategoryID = uuid.New()

	query := `
		INSERT INTO 
			sub_categories (
				sub_category_id,
				name,
				category_id
		) VALUES ($1,$2,$3)
	`
	_, err := s.db.Exec(ctx, query, subcategory.SubCategoryID, subcategory.Name, subcategory.CategoryID)
	if err != nil {
		s.log.Error("error on Creating new SubCategory", logger.Error(err))
		return nil, err
	}

	subcat, err := s.GetSubCategory(ctx, subcategory.SubCategoryID.String())
	if err != nil {
		s.log.Error("error on Getting new SubCategory", logger.Error(err))
		return nil, err
	}
	return subcat, nil
}
func (s *subcategoryRepo) GetSubCategories(ctx context.Context, page, limit int32) (*models.GetSubCategoriesLisResp, error) {
	s.log.Debug("request in GetSubCategories.")

	query := `
		SELECT
			* 
		FROM 
			sub_categories 
		LIMIT 
			$1
		OFFSET
			$2
	`
	offset := (page - 1) * limit
	rows, err := s.db.Query(ctx, query, limit, offset)
	if err != nil {
		s.log.Error("error on Getting all SubCategory ", logger.Error(err))
		return nil, err
	}

	defer rows.Close()

	var subcategories []*models.SubCategory

	for rows.Next() {
		var subcategory models.SubCategory

		err := rows.Scan(
			&subcategory.SubCategoryID,
			&subcategory.Name,
			&subcategory.CreatedAt,
			&subcategory.CategoryID,
		)
		if err != nil {
			s.log.Error("error on scaning  SubCategory ", logger.Error(err))
			return nil, err
		}

		subcategories = append(subcategories, &subcategory)
	}

	var count int

	err = s.db.QueryRow(ctx, "SELECT count(*) FROM sub_categories ").Scan(&count)
	if err != nil {
		s.log.Error("error on scaning  SubCategory count ", logger.Error(err))
		return nil, err
	}

	return &models.GetSubCategoriesLisResp{
		SubCategories: subcategories,
		Count:         count,
	}, nil
}
func (s *subcategoryRepo) GetSubCategory(ctx context.Context, id string) (*models.SubCategory, error) {
	s.log.Debug("request in GetSubCategory.")

	var subcategory models.SubCategory

	query := `SELECT
				 * 
			  FROM 
			  	sub_categories 
			  WHERE 
			  	sub_category_id = $1 `

	err := s.db.QueryRow(ctx, query, id).Scan(
		&subcategory.SubCategoryID,
		&subcategory.Name,
		&subcategory.CreatedAt,
	)

	if err != nil {
		s.log.Error("error on Getting  SubCategory by id", logger.Error(err))
		return nil, err
	}
	return &subcategory, nil
}
func (s *subcategoryRepo) UpdateSubCategory(ctx context.Context, category *models.SubCategory) (*models.SubCategory, error) {
	return nil, nil
}
func (s *subcategoryRepo) DeleteSubCategory(ctx context.Context, id string) error {
	return nil
}
