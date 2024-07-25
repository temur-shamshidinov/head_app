package v1

import (
	"head_app/models"
	"head_app/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

// ! subcategory

func (h *handlers) CreateSubCategory(ctx *gin.Context) {
	var reqBody models.CreateSubCategoryReq

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		h.log.Error("error in binding req body", logger.Error(err))
		return
	}

	subCategory := &models.SubCategory{}

	helpers.DataParser(reqBody, &subCategory)

	subCategory, err = h.storage.GetSubCategoryRepo().CreateSubCategory(ctx, subCategory)
	if err != nil {
		h.log.Error("error on creating new subcategory", logger.Error(err))
		return
	}

	ctx.JSON(201, subCategory)
}

func (h *handlers) GetSubCategoriesList(ctx *gin.Context) {

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	subCategories, err := h.storage.GetSubCategoryRepo().GetSubCategories(ctx, helpers.GetPage(page), helpers.GetLimit(limit))
	if err != nil {
		h.log.Error("error on creating new subcategory", logger.Error(err))
		return
	}

	ctx.JSON(200, subCategories)
}

func (h *handlers) GetSubCategory(ctx *gin.Context) {

	id := ctx.Param("id")

	subCategory := &models.SubCategory{}

	subCategory, err := h.storage.GetSubCategoryRepo().GetSubCategory(ctx, id)
	if err != nil {
		h.log.Error("error on getting subcategory", logger.Error(err))
		return
	}

	ctx.JSON(201, subCategory)
}

func (h *handlers) UpdateSubCategory(ctx *gin.Context) {

	id := ctx.Param("id")

	subCategory := &models.SubCategory{}

	subCategory, err := h.storage.GetSubCategoryRepo().GetSubCategory(ctx, id)
	if err != nil {
		h.log.Error("error on getting subcategory", logger.Error(err))
		return
	}

	ctx.JSON(201, subCategory)
}

func (h *handlers) DeleteSubCategory(ctx *gin.Context) {

	id := ctx.Param("id")

	subCategory := &models.SubCategory{}

	subCategory, err := h.storage.GetSubCategoryRepo().GetSubCategory(ctx, id)
	if err != nil {
		h.log.Error("error on getting category", logger.Error(err))
		return
	}

	ctx.JSON(201, subCategory)
}
