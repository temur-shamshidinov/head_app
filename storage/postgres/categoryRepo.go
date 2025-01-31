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

type categoryRepo struct {
	db  *pgx.Conn
	log log.Log
}

func NewCategoryRepo(db *pgx.Conn, log log.Log) repoi.CategoryRepoI {
	return &categoryRepo{db, log}
}

//Category

func (c *categoryRepo) CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	
	c.log.Debug("request in CreateCategory.")

	category.CategoryID = uuid.New()

	query := `
		INSERT INTO 
			categories (
				category_id,
				name
		) VALUES ($1,$2)
	`
	_, err := c.db.Exec(ctx, query, category.CategoryID, category.Name)
	if err != nil {
		c.log.Error("error on Creating new Category", logger.Error(err))
		return nil, err
	}

	cat, err := c.GetCategory(ctx, category.CategoryID.String())
	if err != nil {
		c.log.Error("error on Getting new Category", logger.Error(err))
		return nil, err
	}
	return cat, nil
}

func (c *categoryRepo) GetCategories(ctx context.Context, page, limit int32) (*models.GetCategoriesListResp, error) {
	c.log.Debug("request in GetCategories.")

	query := `
		SELECT
			* 
		FROM 
			categories 
		LIMIT 
			$1
		OFFSET
			$2
	`
	offset := (page - 1) * limit
	rows, err := c.db.Query(ctx, query, limit, offset)
	if err != nil {
		c.log.Error("error on Getting all Category ", logger.Error(err))
		return nil, err
	}

	defer rows.Close()

	var categories []*models.Category

	for rows.Next() {
		var category models.Category

		err := rows.Scan(
			&category.CategoryID,
			&category.Name,
			&category.CreatedAt,
		)
		if err != nil {
			c.log.Error("error on scaning  Category ", logger.Error(err))
			return nil, err
		}

		categories = append(categories, &category)
	}

	var count int32

	err = c.db.QueryRow(ctx, "SELECT count(*) FROM categories").Scan(&count)
	if err != nil {
		c.log.Error("error on scaning  Category count ", logger.Error(err))
		return nil, err
	}

	return &models.GetCategoriesListResp{
		Categories: categories,
		Count:      count,
	}, nil
}
func (c *categoryRepo) GetCategory(ctx context.Context, id string) (*models.Category, error) {
	
	c.log.Debug("request in GetCategory.")

	var category models.Category

	query := `SELECT
				 * 
			  FROM 
			  	categories 
			  WHERE 
			  	category_id = $1 `

	err := c.db.QueryRow(ctx, query, id).Scan(
		&category.CategoryID,
		&category.Name,
		&category.CreatedAt,
	)

	if err != nil {
		c.log.Error("error on Getting  Category by id", logger.Error(err))
		return nil, err
	}
	return &category, nil
}
func (c *categoryRepo) UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	return nil, nil
}
func (c *categoryRepo) DeleteCategory(ctx context.Context, id string) error {
	return nil
}
