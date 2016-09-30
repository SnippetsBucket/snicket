package model

import (
	"github.com/russross/blackfriday"
)

type (
	PreviewJson struct {
		Markdown string `json:"markdown"`
	}

	ParsedMarkdown struct {
		ParsedStr string `json:"parsedStr"`
	}
)

func ParseMarkdown(rq *PreviewJson) string {
	b := blackfriday.MarkdownCommon([]byte(rq.Markdown))
	return string(b)
}

func BuildPreviewJson(parsedMd string) *ParsedMarkdown {
	return &ParsedMarkdown{
		ParsedStr: parsedMd,
	}
}
