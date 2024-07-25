package v1

import (
	"head_app/models"
	"head_app/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

// ! article

func (h *handlers) CreateArticle(ctx *gin.Context) {
	var reqBody models.CreateArticleReq

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		h.log.Error("error in binding req body", logger.Error(err))
		return
	}

	article := &models.Article{}

	helpers.DataParser(reqBody, &article)

	article, err = h.storage.GetArticleRepo().CreateArticle(ctx, article)
	if err != nil {
		h.log.Error("error on creating new article", logger.Error(err))
		return
	}

	ctx.JSON(201, article)
}

func (h *handlers) GetArticleList(ctx *gin.Context) {

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	articles, err := h.storage.GetArticleRepo().GetArticles(ctx, helpers.GetPage(page), helpers.GetLimit(limit))
	if err != nil {
		h.log.Error("error on getting  articles", logger.Error(err))
		return
	}

	ctx.JSON(200, articles)
}

func (h *handlers) GetArticle(ctx *gin.Context) {

	id := ctx.Param("id")

	article, err := h.storage.GetArticleRepo().GetArticle(ctx, id)
	if err != nil {
		h.log.Error("error on getting article", logger.Error(err))
		return
	}

	ctx.JSON(201, article)
}
