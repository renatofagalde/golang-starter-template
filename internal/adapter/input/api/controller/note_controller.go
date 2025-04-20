package controller

import (
	"bootstrap/internal/adapter/input/model/request"
	domain_request "bootstrap/internal/application/domain/request"
	"bootstrap/internal/application/port/input"
	"bootstrap/internal/config/logger"
	"bootstrap/internal/config/validation"
	"bootstrap/internal/constants"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type noteController struct {
	notesUseCase input.NotesUseCase
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewNoteController(notesUseCase input.NotesUseCase) *noteController {
	return &noteController{notesUseCase}
}

func (nc *noteController) ListNotes(ctx *gin.Context) {

	logger.Info("ListNotesController",
		zap.String(constants.H.Stage, "controller"),
		zap.String(constants.H.Journey, ctx.GetHeader(constants.H.Journey)),
		zap.String(constants.H.TraceID, ctx.GetHeader(constants.H.TraceID)))

	var noteRequest = &request.NoteRequest{}

	if err := ctx.ShouldBindQuery(noteRequest); err != nil {
		logger.Error("Error trying to validate fields from request", err)
		errRest := validation.ValidateError(err)
		ctx.JSON(errRest.Code, errRest)
		return
	}

	noteRequestDomain := domain_request.NoteRequest{
		Subject: noteRequest.Subject,
		From:    noteRequest.From,
	}

	notes, err := nc.notesUseCase.ListNotesService(ctx.Request.Context(), noteRequestDomain)
	if err != nil {
		ctx.JSON(err.Code, err)
	}
	ctx.JSON(http.StatusOK, notes)
}
