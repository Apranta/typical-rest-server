package generate

const readmeTemplate = `<!-- Autogenerated by TypicalGo; Modify '_typical/generate/readme_template.go' to add more content -->
# {{ .Context.Name}}

{{ .Context.Description}}

## Configuration

{{ .Configuration}}
`