package v1

import (
	"head_app/models"
	"head_app/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *handlers) CreateCategory(ctx *gin.Context) {
	var reqBody models.CreatedCategoryReq

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		h.log.Error("error in binding req body", logger.Error(err))
		return
	}

	category := &models.Category{}

	helpers.DataParser(reqBody, &category)

	category, err = h.storage.GetCategoryRepo().CreateCategory(ctx, category)
	if err != nil {
		h.log.Error("error on creating new category", logger.Error(err))
		return
	}

	ctx.JSON(201, category)
}

func (h *handlers) GetCategoriesList(ctx *gin.Context) {

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	categories, err := h.storage.GetCategoryRepo().GetCategories(ctx, helpers.GetPage(page), helpers.GetLimit(limit))
	if err != nil {
		h.log.Error("error on creating new category", logger.Error(err))
		return
	}

	ctx.JSON(200, categories)
}

func (h *handlers) GetCategory(ctx *gin.Context) {

	id := ctx.Param("id")

	category := &models.Category{}

	category, err := h.storage.GetCategoryRepo().GetCategory(ctx, id)
	if err != nil {
		h.log.Error("error on getting category", logger.Error(err))
		return
	}

	ctx.JSON(201, category)
}

func (h *handlers) UpdateCategory(ctx *gin.Context) {

	id := ctx.Param("id")

	category := &models.Category{}

	category, err := h.storage.GetCategoryRepo().GetCategory(ctx, id)
	if err != nil {
		h.log.Error("error on getting category", logger.Error(err))
		return
	}

	ctx.JSON(201, category)
}

func (h *handlers) DeleteCategory(ctx *gin.Context) {

	id := ctx.Param("id")

	category := &models.Category{}

	category, err := h.storage.GetCategoryRepo().GetCategory(ctx, id)
	if err != nil {
		h.log.Error("error on getting category", logger.Error(err))
		return
	}

	ctx.JSON(201, category)
}
