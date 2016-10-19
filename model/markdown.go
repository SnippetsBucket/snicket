package model

import "github.com/shurcooL/github_flavored_markdown"

type (
	PreviewJson struct {
		Markdown string `json:"markdown"`
	}

	ParsedMarkdown struct {
		ParsedStr string `json:"parsedStr"`
	}
)

func ParseMarkdown(rq *PreviewJson) string {
	b := github_flavored_markdown.Markdown([]byte(rq.Markdown))
	return string(b)
}

func BuildPreviewJson(parsedMd string) *ParsedMarkdown {
	return &ParsedMarkdown{
		ParsedStr: parsedMd,
	}
}
