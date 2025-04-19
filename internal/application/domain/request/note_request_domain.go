package domain_request

import "time"

type NoteRequest struct {
	Subject string
	From    time.Time
}
