package note_sql

import (
	"bootstrap/internal/adapter/output/note_sql/entity"
	domain_request "bootstrap/internal/application/domain/request"
	domain_response "bootstrap/internal/application/domain/response"
	"bootstrap/internal/config/logger"
	"bootstrap/internal/config/rest_err"
	"bootstrap/internal/constants"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

func (nr *noteRepository) GetNotesPort(ctx context.Context, noteRequest domain_request.NoteRequest) (*domain_response.NoteResponseDomain, *rest_err.RestErr) {
	logger.Info(fmt.Sprintf("Getting notes from SQL database, subject=%s, from=%s",
		noteRequest.Subject, noteRequest.From),
		zap.String(constants.H.Stage, "repository"),
		zap.String(constants.H.Journey, constants.H.GetJourney(ctx)),
		zap.String(constants.H.TraceID, constants.H.GetTraceID(ctx)))

	var articles []entity.Article = make([]entity.Article, 0)
	result := nr.database.WithContext(ctx).
		Table("note.article").
		Find(&articles)
	if result.Error != nil {
		return nil, rest_err.NewInternalServerError(result.Error.Error())
	}

	articleResponseDomain := []domain_response.ArticleResponseDomain{}
	copier.Copy(&articleResponseDomain, &articles)
	notesDomain := domain_response.NoteResponseDomain{
		Status:       "ok",
		TotalResults: len(articleResponseDomain),
		Articles:     articleResponseDomain,
	}

	return &notesDomain, nil
}
