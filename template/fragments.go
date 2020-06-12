package template

import (
	"text/template"
)

// This file contains all the dynamic template fragments for each language.

// language templates
var templates = map[Language]langTemplates{
	Elm:        elmTemplates,
	Flow:       flowTemplates,
	Typescript: tsTemplates,
}

type langTemplates struct {
	header          string
	arrayOpen       string
	arrayClose      string
	arrayShortOpen  string
	arrayShortClose string
	basic           string
	fieldDocComment string
	declaration     string
	fieldClose      string
	lastFieldClose  string
	fieldName       string
	mapClose        string
	mapKey          string
	mapValue        string
	structClose     string
	structOpen      string
	timeType        string
}

// newTemplate returns the template string for a language and a string
func newTemplate(tpl string) *template.Template {
	return template.Must(template.New("dummy").
		Funcs(funcMap).
		Parse(tpl))
}

var elmTemplates = langTemplates{
	header: `-- Automatically generated by typewriter. Do not edit.
-- http://www.github.com/natdm/typewriter

`,
	arrayOpen:       ` List`,
	arrayClose:      ``,
	arrayShortOpen:  ` List`,
	arrayShortClose: ``,
	basic:           ` {{updateElmType .Type}}`,
	fieldDocComment: `{{elmMultilineComment .DocComment 1}}`,
	declaration: `
{{elmMultilineComment .Comment 0}}type alias {{.Name}} : `,
	fieldClose: `,{{elmComment .LineComment}}
`,
	lastFieldClose: `{{elmComment .LineComment}}
`, // Elm has no trailing comma support
	fieldName: `	{{.Name}} :`,
	mapClose:    ``,
	mapKey:      `Dict `,
	mapValue:    ` `,
	structClose: `}`,
	structOpen: `
{`,
	timeType: "Date",
}

var flowTemplates = langTemplates{
	header: `// @flow
// Automatically generated by typewriter. Do not edit.
// http://www.github.com/natdm/typewriter

`,
	arrayOpen:       `Array<`,
	arrayClose:      `>`,
	arrayShortOpen:  ``,
	arrayShortClose: `[]`,
	basic:           `{{if .Pointer}}?{{end}}{{updateFlowType .Type}}`,
	fieldDocComment: `{{flowMultilineComment .DocComment 1}}`,
	declaration: `
{{flowMultilineComment .Comment 0}}export type {{.Name}} = `,
	fieldClose: `,{{flowComment .LineComment}}
`,
	fieldName: `	{{.Name}}: `,
	mapClose:    ` }`,
	mapKey:      `{ [key: `,
	mapValue:    `]: `,
	structClose: `{{ range .Embedded}}{{.}} & {{end}}{{if .Strict}}|}{{else}}}{{end}}`,
	structOpen: `{{"{"}}{{if .Strict}}|{{end}}
`,
	timeType: "Date",
}

var tsTemplates = langTemplates{
	header: `// Automatically generated by typewriter. Do not edit.
// http://www.github.com/natdm/typewriter

`,
	arrayOpen:       `Array<`,
	arrayClose:      `>`,
	arrayShortOpen:  ``,
	arrayShortClose: `[]`,
	basic:           `{{updateTSType .Type}}{{if .Pointer}} | undefined{{end}}`,
	fieldDocComment: `{{tsMultilineComment .DocComment 1}}`,
	declaration: `
{{tsMultilineComment .Comment 0}}type {{.Name}} = `,
	fieldClose: `,{{tsComment .LineComment}}
`,
	fieldName: `	{{.Name}}{{if .Type.IsPointer}}?{{end}}: `,
	mapClose:    ` }`,
	mapKey:      `{ [key: `,
	mapValue:    `]: `,
	structClose: `}`,
	structOpen: `{{ range .Embedded}}{{ . }} & {{end}}{
`,
	timeType: "Date",
}
