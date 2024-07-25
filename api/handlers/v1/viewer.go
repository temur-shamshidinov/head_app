package v1

import (
	"encoding/json"
	"head_app/models"
	"head_app/pkg/helpers"
	"head_app/pkg/mail"
	"head_app/pkg/token"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *handlers) CheckUser(ctx *gin.Context) {
	var reqBody *models.CheckViewer

	ctx.Bind(&reqBody)

	isExists, err := h.storage.GetCommonRepo().CheckIsExists(ctx, &models.Common{
		TableName:  "viewers",
		ColumnName: "gmail",
		ExpValue:   reqBody.Gmail,
	})

	if err != nil {
		h.log.Error("error on checking viewer ", logger.Error(err))
		return
	}

	if isExists {
		ctx.JSON(201, models.CheckExistsResp{
			IsExists: isExists,
			Status:   "log-in",
		})
		return
	}

	otp := models.OtpData{
		Gmail: reqBody.Gmail,
		Otp:   mail.GenerateOtp(6),
	}

	otpDataByt, err := json.Marshal(otp)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	err = h.cache.Set(ctx, reqBody.Gmail, string(otpDataByt), 60)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	err = mail.SendMail([]string{reqBody.Gmail}, otp.Otp)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(201, models.CheckExistsResp{
		IsExists: isExists,
		Status:   "register",
	})
}

func (h *handlers) CheckOTP(ctx *gin.Context) {

	var reqBody models.OtpData

	err := ctx.Bind(&reqBody)
	if err != nil {
		h.log.Error("error binding reqBody", logger.Error(err))
		return
	}

	gmail := reqBody.Gmail

	data, err := h.cache.GetDel(ctx, gmail)
	if err != nil {
		h.log.Error("error on getting data from cache", logger.Error(err))
		return
	}

	if data == "" {
		ctx.JSON(201, "otp is expired")
	}

	var cacheData models.OtpData

	json.Unmarshal([]byte(data), &cacheData)

	ctx.JSON(201, models.CheckOTPResp{
		IsRight: cacheData.Otp == reqBody.Otp,
	})

}

func (h *handlers) SignUp(ctx *gin.Context) {
	var regReqBody models.ViewerRegReq

	err := ctx.Bind(&regReqBody)
	if err != nil {
		return
	}

	otpStrData, err := h.cache.GetDel(ctx, regReqBody.Gmail)
	if err != nil {
		return
	}

	if otpStrData == "" {
		ctx.JSON(201, "otp is expired")
	}

	var otpData models.OtpData

	err = json.Unmarshal([]byte(otpStrData), &otpData)
	if err != nil {
		return
	}

	if otpData.Otp != regReqBody.Otp {
		ctx.JSON(405, "otp is incorrect")
		return
	}

	var viewer = &models.Viewer{}

	err = helpers.DataParser(regReqBody, &viewer)
	if err != nil {
		return
	}

	viewer.ViewerID = uuid.New()
	viewer.Password, err = helpers.HashPassword(viewer.Password)
	if err != nil {
		return
	}

	claim, err := h.storage.GetViewerRepo().CreateViewer(ctx, viewer)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	accesstoken, err := token.GenerateJWT(*claim)
	if err != nil {
		ctx.JSON(201, "registired")
		return
	}

	ctx.JSON(201, &models.RespAuth{AccessToken: accesstoken})
}

func (h *handlers) SignIn(ctx *gin.Context) {

	var reqBody *models.LogInViewer

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		return
	}

	claim, err := h.storage.GetViewerRepo().LogIn(ctx, reqBody)
	if err != nil {
		if err.Error() == "password in incorrect" {
			ctx.JSON(405, err)
			return
		}
		ctx.JSON(500, err)
		return
	}

	accesstoken, err := token.GenerateJWT(*claim)
	if err != nil {
		return
	}

	ctx.JSON(201, &models.RespAuth{AccessToken: accesstoken})

}

func (h *handlers) AddComment(ctx *gin.Context) {

	var reqBody models.CreateCommentReq

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		return
	}

	var comment = &models.Comment{}

	helpers.DataParser(reqBody, &comment)

	comment, err = h.storage.GetViewerRepo().AddComment(ctx,comment)
	if err != nil {
		return
	}

	ctx.JSON(201,comment)
}
