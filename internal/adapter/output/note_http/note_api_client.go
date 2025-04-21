package note_http

import (
	domain_request "bootstrap/internal/application/domain/request"
	domain_response "bootstrap/internal/application/domain/response"
	"bootstrap/internal/config/env"
	"bootstrap/internal/config/logger"
	"bootstrap/internal/config/rest_err"
	"bootstrap/internal/constants"
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

var (
	client *resty.Client
)

type noteClient struct {
}

func NewNoteClient() *noteClient {
	//client := resty.New().SetBaseURL("https://newsapi.org/v2")
	return &noteClient{}
}

func (nc *noteClient) GetNotesPort(ctx context.Context, noteDomainRequest domain_request.NoteRequest) (*domain_response.NoteResponseDomain, *rest_err.RestErr) {

	logger.Info(fmt.Sprintf("Getting notes from SQL database, subject=%s, from=%s",
		noteDomainRequest.Subject, noteDomainRequest.From),
		zap.String(constants.H.Stage, "note_http"),
		zap.String(constants.H.Journey, constants.H.GetJourney(ctx)),
		zap.String(constants.H.TraceID, constants.H.GetTraceID(ctx)))

	//var noteModelResponse domain_response.NoteResponseDomain = domain_response.NoteResponseDomain{}
	noteModelResponse := &domain_response.NoteResponseDomain{}
	client := resty.New().SetBaseURL("https://newsapi.org/v2")

	_, err := client.R().
		SetQueryParams(map[string]string{
			"q":      noteDomainRequest.Subject,
			"from":   noteDomainRequest.From.Format("2006-01-20"),
			"apikey": env.GetNewsTokenAPI(),
		}).SetResult(noteModelResponse).Get("/everything")
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to call notes with params")
	}

	notesDomainResponse := &domain_response.NoteResponseDomain{}
	copier.Copy(notesDomainResponse, noteModelResponse)

	return notesDomainResponse, nil
}
