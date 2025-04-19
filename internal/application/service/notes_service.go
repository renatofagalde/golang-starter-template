package service

import (
	domain_request "bootstrap/internal/application/domain/request"
	domain_response "bootstrap/internal/application/domain/response"
	"bootstrap/internal/config/logger"
	"bootstrap/internal/config/rest_err"
	"bootstrap/internal/constants"
	"context"
	"fmt"
	"go.uber.org/zap"
)

type noteService struct{}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewNoteService() *noteService {
	return &noteService{}
}

func (ns *noteService) ListNotesService(ctx context.Context, request domain_request.NoteRequest) (*domain_response.NoteResponseDomain, *rest_err.RestErr) {

	logger.Info(fmt.Sprintf("ListNotesService, subject=%s, from=%s, to=%s",
		request.Subject, request.From),
		zap.String(constants.H.Stage, "service"),
		zap.String(constants.H.Journey, ctx.Value(constants.H.Journey).(string)),
		zap.String(constants.H.TraceID, ctx.Value(constants.H.TraceID).(string)))

	return nil, nil
}
