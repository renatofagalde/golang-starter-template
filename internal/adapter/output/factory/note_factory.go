package factory

import (
	"bootstrap/internal/adapter/output/note_sql"
	"bootstrap/internal/application/port/output"
	"gorm.io/gorm"
)

type NoteFactory struct {
	implementations map[string]output.NotePort
}

func NewNoteFactory(databae *gorm.DB, httpClient output.NotePort) *NoteFactory {
	implementations := map[string]output.NotePort{
		"sql":  note_sql.NewNoteRepository(databae),
		"http": httpClient,
	}

	return &NoteFactory{
		implementations: implementations,
	}
}

func (nf *NoteFactory) GetNotePort(action string) output.NotePort {
	implementation, exists := nf.implementations[action]

	if !exists {
		return nf.implementations["sql"]
	}

	return implementation
}
