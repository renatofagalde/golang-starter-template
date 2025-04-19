package controller

import (
	"bootstrap/internal/config/logger"
	"bootstrap/internal/config/validation"
	"bootstrap/internal/constants"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/adapter/input/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type noteController struct{}

func NewNoteController() *noteController {
	return &noteController{}
}

func (*noteController) ListNotes(ctx *gin.Context) {

	logger.Info("ListNotes",
		zap.String(constants.H.Journey, ctx.GetHeader(constants.H.Journey)),
		zap.String(constants.H.TraceID, ctx.GetHeader(constants.H.TraceID)))

	var noteRequest = &request.NewsRequest{}

	if err := ctx.ShouldBindQuery(noteRequest); err != nil {
		logger.Error("Error trying to validate fields from request", err)
		errRest := validation.ValidateUserError(err)
		ctx.JSON(errRest.Code, errRest)
		return
	}
}
