package service

import (
	domain_request "bootstrap/internal/application/domain/request"
	domain_response "bootstrap/internal/application/domain/response"
	"bootstrap/internal/application/port/output"
	"bootstrap/internal/config/logger"
	"bootstrap/internal/config/rest_err"
	"bootstrap/internal/constants"
	"context"
	"fmt"
	"go.uber.org/zap"
)

type noteService struct {
	notePort output.NotePort
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewNoteService(notePort output.NotePort) *noteService {
	return &noteService{notePort: notePort}
}

func (ns *noteService) ListNotesService(ctx context.Context, noteDomainRequest domain_request.NoteRequest) (*domain_response.NoteResponseDomain, *rest_err.RestErr) {

	logger.Info(fmt.Sprintf("ListNotesService, subject=%s, from=%s, to=%s",
		noteDomainRequest.Subject, noteDomainRequest.From),
		zap.String(constants.H.Stage, "service"),
		zap.String(constants.H.Journey, ctx.Value(constants.H.Journey).(string)),
		zap.String(constants.H.TraceID, ctx.Value(constants.H.TraceID).(string)))

	noteDomainResponse, err := ns.notePort.GetNotesPort(noteDomainRequest)
	return noteDomainResponse, err
}
