package main

import (
	"bytes"
	"log/slog"
	"text/template"
)

const mdTemplate = `# Weekly Report for {{.RepoName}}

## Commits :tada:

{{range .Commits}}
- {{.Author}}: {{.Commits}} commits
{{end}}

## Contributions :arigato:

{{range .Contributions}}
- {{.Author}}: {{.Commits}} commits
{{end}}

## Pull Requests :blob_bongo:

{{range .PullRequests}}
- [{{.Title}}]({{.URL}}) - Created at: {{.CreatedAt}}
{{end}}

## Issues

{{range .Issues}}
- [{{.Title}}]({{.URL}}) - Created at: {{.CreatedAt}}
{{end}}`

func parse() (string, error) {
	t, err := template.New("TemplateName").Parse(mdTemplate)
	if err != nil {
		slog.Error("%d", err)
	}

	var b *bytes.Buffer

	t.Execute(b, "")

	text := b.String()

	return text, nil
}
