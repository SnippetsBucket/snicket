package model

import (
	"time"

	"github.com/gocraft/dbr"
)

type (
	Snippet struct {
		SnippetId    int64     `db:"snippet_id"`
		SnippetTitle string    `db:"snippet_title"`
		SnippetText  string    `db:"snippet_text"`
		UserId       int64     `db:"user_id"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
	}

	SnippetJson struct {
		SnippetTitle string `json:"snippetTitle"`
		SnippetText  string `json:"snippetText"`
	}
)

// FIXME: `UserId` is not literal
func CreateSnippet(snippetTitle string, snippetText string) *Snippet {
	return &Snippet{
		SnippetId:    0,
		SnippetTitle: snippetTitle,
		SnippetText:  snippetText,
		UserId:       1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (s *Snippet) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("snippets").
		Columns("snippet_id", "snippet_title", "snippet_text", "user_id", "created_at", "updated_at").
		Record(s).
		Exec()

	return err
}
