package request

import "time"

type NoteRequest struct {
	Subject string    `form:"subject" binding:"required,min=4" `
	From    time.Time `form:"from" binding:"required" time_format:"1980-08-28"`
}
