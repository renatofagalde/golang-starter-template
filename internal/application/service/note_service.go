package service

import (
	"bootstrap/internal/adapter/output/factory"
	domain_request "bootstrap/internal/application/domain/request"
	domain_response "bootstrap/internal/application/domain/response"
	"bootstrap/internal/config/logger"
	"bootstrap/internal/config/rest_err"
	"bootstrap/internal/constants"
	"context"
	"fmt"
	"go.uber.org/zap"
)

type noteService struct {
	noteFactory *factory.NoteFactory
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewNoteService(noteFactory *factory.NoteFactory) *noteService {
	return &noteService{noteFactory: noteFactory}
}

func (ns *noteService) ListNotesService(ctx context.Context, noteDomainRequest domain_request.NoteRequest) (*domain_response.NoteResponseDomain, *rest_err.RestErr) {

	logger.Info(fmt.Sprintf("ListNotesService, subject=%s, from=%s, to=%s",
		noteDomainRequest.Subject, noteDomainRequest.From),
		zap.String(constants.H.Stage, "service"),
		zap.String(constants.H.Journey, constants.H.GetJourney(ctx)),
		zap.String(constants.H.TraceID, constants.H.GetTraceID(ctx)))

	noteFactory := ns.noteFactory.GetNotePort(noteDomainRequest.Action)

	noteDomainResponse, err := noteFactory.GetNotesPort(ctx, noteDomainRequest)
	return noteDomainResponse, err

}
