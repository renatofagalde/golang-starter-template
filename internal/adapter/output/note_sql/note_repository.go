package note_sql

import (
	"gorm.io/gorm"
)

type noteRepository struct {
	database *gorm.DB
}

func NewNoteRepository(database *gorm.DB) *noteRepository {
	return &noteRepository{database}
}
