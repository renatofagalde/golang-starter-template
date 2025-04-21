package model_request

import "time"

type NoteModelRequest struct {
	Subject string    `form:"subject" binding:"required,min=4" `
	From    time.Time `form:"from" binding:"required" time_format:"2006-01-02"`
	Action  string    `json:"action" validate:"required,oneof=http sql"`
}
