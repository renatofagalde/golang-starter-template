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

	var noteEntities []entity.NoteEntity
	query := nr.database.Model(&entity.NoteEntity{})

	if err := query.First(&noteEntities).Error; err != nil {
		logger.Error("Error getting notes from database", err,
			zap.String(constants.H.Stage, "repository"),
			zap.String(constants.H.Journey, constants.H.GetJourney(nil)),
			zap.String(constants.H.TraceID, constants.H.GetTraceID(nil)))
		return nil, rest_err.NewInternalServerError("Error getting notes from database")
	}

	noteDomain := domain_response.ArticleResponseDomain{
		Author:      noteEntities[0].Author,
		Title:       noteEntities[0].Title,
		Description: noteEntities[0].Description,
		URL:         noteEntities[0].URL,
		URLToImage:  noteEntities[0].URLToImage,
		PublishedAt: noteEntities[0].PublishedAt,
		Content:     noteEntities[0].Content,
	}

	//var n [1]domain_response.ArticleSourceResponseDomain = [1]domain_response.ArticleSourceResponseDomain
	//n[0] = noteDomain
	//domain_response.NoteResponseDomain{
	//	Status:       "ok",
	//	TotalResults: 0,
	//	Articles:     n,
	//}

	copier.Copy(&noteDomain, &noteEntities)

	return &noteDomain, nil
}
