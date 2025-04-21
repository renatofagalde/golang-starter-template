package output

import (
	domain_request "bootstrap/internal/application/domain/request"
	domain_response "bootstrap/internal/application/domain/response"
	"bootstrap/internal/config/rest_err"
	"context"
)

type NotePort interface {
	GetNotesPort(context.Context, domain_request.NoteRequest) (*domain_response.NoteResponseDomain, *rest_err.RestErr)
}
