package entity

type NoteEntity struct {
	Source      string
	Author      string
	Title       string
	Description string
	URL         string
	URLToImage  string
	PublishedAt string
	Content     string
}

func (NoteEntity) TableName() string {
	return "note.note"
}
