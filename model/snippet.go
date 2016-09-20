package model

import (
	"time"

	"github.com/gocraft/dbr"
)

type Snippet struct {
	SnippetId    int64  `db:"snippet_id"`
	SnippetTitle string `db:"snippet_title"`
	SnippetText  string `db:"snippet_text"`
	UserId       int64  `db:"user_id"`
	CreatedAt    int64  `db:"created_at"`
	UpdatedAt    int64  `db:"updated_at"`
}

func CreateSnippet(snippetTitle string, snippetText string) *Snippet {
	return &Snippet{
		SnippetTitle: snippetTitle,
		SnippetText:  snippetText,
		CreatedAt:    time.Now().Unix(),
	}
}

func Save(tx *dbr.Tx, snippetTitle string, snippetText string) error {
	_, err := tx.InsertInto("snippet").
		Columns("snippet_title", "snippet_text", "created_at").
		Values(snippetTitle, snippetText).
		Exec()

	return err
}

// func (s *Snippet) Save(tx *dbr.Tx) error {
// 	_, err := tx.InsertInto("snippet").
// 		Columns("snippet_title", "snippet_text", "created_at").
// 		Record(s).
// 		Exec()

// 	return err
// }
