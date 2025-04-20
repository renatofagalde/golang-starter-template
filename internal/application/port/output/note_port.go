package output

import (
	domain_request "bootstrap/internal/application/domain/request"
	domain_response "bootstrap/internal/application/domain/response"
	"bootstrap/internal/config/rest_err"
)

type NotePort interface {
	GetNotesPort(domain_request.NoteRequest) (*domain_response.NoteResponseDomain, *rest_err.RestErr)
}
