package output

import _ "embed"

var PrintTemplate = `{{ range . -}}
{{ .File }} lines({{.LinesHit}}/{{.LinesFound}}) functions({{.FunctionsHit}}/{{.FunctionsFound}}) branches({{.BranchesHit}}/{{.BranchesFound}})
{{ end }}`

//go:embed html.template
var HtmlTemplate string
