package model

import (
	"time"

	"github.com/gocraft/dbr"
)

type Snippet struct {
	SnippetId    int64  `form:"snippet_id"`
	SnippetTitle string `form:"snippet_title"`
	SnippetText  string `form:"snippet_text"`
	CreatedAt    int64  `form:"created_at"`
	UpdatedAt    int64  `form:"updated_at"`
}

func CreateSnippet(snippetTitle string, snippetText string) *Snippet {
	return &Snippet{
		SnippetTitle: snippetTitle,
		SnippetText:  snippetText,
		CreatedAt:    time.Now().Unix(),
	}
}

func (s *Snippet) Save(tx *dbr.Tx) error {
	_, err := tx.InsertInto("snippet").
		Columns("snippet_title", "snippet_text", "created_at").
		Record(s).
		Exec()

	return err
}
