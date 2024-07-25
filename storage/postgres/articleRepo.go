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

type articleRepo struct {
	db  *pgx.Conn
	log log.Log
}

func NewArticleRepo(db *pgx.Conn, log log.Log) repoi.ArticleRepoI {
	return &articleRepo{db, log}
}

// Article
func (a *articleRepo) CreateArticle(ctx context.Context, article *models.Article) (*models.Article, error) {
	a.log.Debug("request in CreateArticle.")

	article.ArticleID = uuid.New()

	query := `
		INSERT INTO 
			articles (
				article_id,     
				title,     
				content,          
				category_id,     
				sub_category_id 
		) VALUES ($1,$2,$3,$4,$5)
	`
	_, err := a.db.Exec(
		ctx, query,
		article.ArticleID,
		article.Title,
		article.Content,
		article.CategoryID,
		article.SubCategoryID,
	)
	if err != nil {
		a.log.Error("error on Creating new Article", logger.Error(err))
		return nil, err
	}

	artc, err := a.GetArticle(ctx, article.ArticleID.String())
	if err != nil {
		a.log.Error("error on Getting new Article", logger.Error(err))
		return nil, err
	}
	return artc, nil
}
func (a *articleRepo) GetArticles(ctx context.Context, page, limit int32) (*models.GetArticleListResp, error) {

	a.log.Debug("request in GetArticles.")

	query := `
		SELECT
			 article_id,    
 			 title,          
 			 content,        
 			 created_at,     
 			 updated_at,    
 			 category_id,    
 			 sub_category_id
		FROM 
			articles 
		LIMIT 
			$1
		OFFSET
			$2
		WHERE 
			deleted_at IS  NULL
	`
	offset := (page - 1) * limit
	rows, err := a.db.Query(ctx, query, limit, offset)
	if err != nil {
		a.log.Error("error on Getting all Article ", logger.Error(err))
		return nil, err
	}

	defer rows.Close()

	var articles []*models.Article

	for rows.Next() {
		var article models.Article

		err := rows.Scan(
			&article.ArticleID,
			&article.Title,
			&article.Content,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.CategoryID,
			&article.SubCategoryID,
		)
		if err != nil {
			a.log.Error("error on scaning  Article ", logger.Error(err))
			return nil, err
		}

		articles = append(articles, &article)
	}

	var count int

	err = a.db.QueryRow(ctx, "SELECT count(*) FROM articles").Scan(&count)
	if err != nil {
		a.log.Error("error on scaning  Articles count ", logger.Error(err))
		return nil, err
	}

	return &models.GetArticleListResp{
		Articles: articles,
		Count:    count,
	}, nil
}
func (a *articleRepo) GetArticle(ctx context.Context, id string) (*models.Article, error) {

	a.log.Debug("request in GetArticle.")

	var article models.Article

	query := `SELECT
				 * 
			  FROM 
			  	articles 
			  WHERE 
			  	article_id = $1 `

	err := a.db.QueryRow(ctx, query, id).Scan(
		&article.CategoryID,
		&article.Title,
		&article.Content,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.CategoryID,
		&article.SubCategoryID,

	)

	if err != nil {
		a.log.Error("error on Getting  Article by id", logger.Error(err))
		return nil, err
	}
	return &article, nil
}
func (a *articleRepo) UpdateArticle(ctx context.Context, category *models.Article) (*models.Article, error) {
	return nil, nil
}
func (a *articleRepo) DeleteArticle(ctx context.Context, id string) error {
	return nil
}
