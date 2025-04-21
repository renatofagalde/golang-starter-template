package entity

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	ID        int64     `json:"id" gorm:"primaryKey;column:id"`
	IDE       uuid.UUID `json:"ide" gorm:"type:uuid;default:gen_random_uuid();column:ide;uniqueIndex"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	Deleted   bool      `json:"deleted" gorm:"default:false;column:deleted"`
}

type ArticleSource struct {
	Base
	SourceID *string   `json:"source_id" gorm:"column:source_id"`
	Name     string    `json:"name" gorm:"not null;column:name"`
	Articles []Article `json:"articles" gorm:"foreignKey:SourceID;references:ID"`
}

func (ArticleSource) TableName() string {
	return "note.article_source"
}

type Article struct {
	Base
	SourceID    int64         `json:"source_id" gorm:"not null;column:source_id"`
	Author      *string       `json:"author" gorm:"column:author"`
	Title       string        `json:"title" gorm:"not null;column:title"`
	Description *string       `json:"description" gorm:"type:text;column:description"`
	URL         string        `json:"url" gorm:"not null;uniqueIndex;column:url;size:512"`
	URLToImage  *string       `json:"url_to_image" gorm:"column:url_to_image;size:512"`
	PublishedAt *time.Time    `json:"published_at" gorm:"column:published_at"`
	Content     *string       `json:"content" gorm:"type:text;column:content"`
	Source      ArticleSource `json:"-" gorm:"foreignKey:SourceID;references:ID"`
	Notes       []Note        `json:"notes" gorm:"many2many:note.note_article;foreignKey:ID;joinForeignKey:article_id;References:ID;joinReferences:note_id"`
}

func (Article) TableName() string {
	return "note.article"
}

type Note struct {
	Base
	Status       string    `json:"status" gorm:"not null;column:status;size:100"`
	TotalResults int       `json:"total_results" gorm:"not null;default:0;column:total_results"`
	QueryText    *string   `json:"query_text" gorm:"column:query_text"`
	Articles     []Article `json:"articles" gorm:"many2many:note.note_article;foreignKey:ID;joinForeignKey:note_id;References:ID;joinReferences:article_id"`
}

func (Note) TableName() string {
	return "note.note"
}

type NoteArticle struct {
	Base
	NoteID    int64   `json:"note_id" gorm:"not null;column:note_id"`
	ArticleID int64   `json:"article_id" gorm:"not null;column:article_id"`
	Note      Note    `json:"-" gorm:"foreignKey:NoteID;references:ID"`
	Article   Article `json:"-" gorm:"foreignKey:ArticleID;references:ID"`
}

// TableName define o nome da tabela
func (NoteArticle) TableName() string {
	return "note.note_article"
}
