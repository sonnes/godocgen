package godocgen

var pkgTemplate = `
# {{ .Name }}

* [Overview](#pkg-overview)
* [Index](#pkg-index){{if $.Examples}}
* [Examples](#pkg-examples){{- end}}

## <a name="pkg-overview">Overview</a>
{{comment_md .Doc}}

## <a name="pkg-index">Index</a>

{{ if .Consts }}* [Constants](#pkg-constants){{ end }}{{ if .Vars }}
* [Variables](#pkg-variables){{ end }}{{ range .Funcs }}
* [{{ .Snippet }}](#{{  .Name | slugify  }}){{ end }}{{ range .Types }}{{$typeName := .Name}}
* [type {{ .Name }}](#{{  .Name | slugify  }}){{ range .Funcs }}
	* [{{ .Snippet }}](#{{  .Name | slugify  }}){{ end }}{{ range .Methods }}
	* [{{ .Snippet }}](#{{ $typeName | slugify }}-{{  .Name | slugify  }}){{ end }}{{ end }}
{{if $.Examples}}
#### <a name="pkg-examples">Examples</a>{{ range .Examples }}
* [{{.Name}}](#example-{{ .Name | slugify }}){{ end }}{{ end }}

{{with .Consts}}## <a name="pkg-constants">Constants</a>
{{range .}}{{comment_md .Doc}}
{{ .Snippet | pre}}{{end}}{{end}}
{{with .Vars}}## <a name="pkg-variables">Variables</a>
{{range .}}{{comment_md .Doc}}{{ .Snippet | pre}}{{end}}{{end}}

{{range .Funcs}}## <a name="{{ .Name | slugify }}">func {{ .Name }}</a>
{{ .Snippet | pre}}
{{comment_md .Doc}}{{ range .Examples }}
#### <a name="example-{{ .Name | slugify }}">Example {{ .Suffix }}</a>
Code:
{{ .Snippet | pre}}{{ if .Output }}
Output:
{{ .Output | pre}}{{ end }}
{{ comment_md .Doc}}{{ end }}{{ end }}

{{range .Types}}{{ $typeName := .Name }}## <a name="{{ .Name | slugify }}">type {{ .Name }}</a>
{{ .Snippet | pre}}
{{comment_md .Doc}}
{{range .Funcs}}{{$funcName := .Name}}### <a name="{{ $funcName | slugify}}">func {{ $funcName }}</a>
{{ .Snippet | pre }}
{{comment_md .Doc}}{{ range .Examples }}
#### <a name="example-{{ .Name | slugify }}">Example {{ .Suffix }}</a>
Code:
{{ .Snippet | pre}}{{ if .Output }}
Output:
{{ .Output | pre}}{{ end }}
{{ comment_md .Doc}}{{ end }}{{end}}
{{range .Methods}}{{$funcName := .Name}}### <a name="{{ $typeName | slugify }}-{{ $funcName | slugify}}">func ({{ .Recv }}) {{ .Name }}</a>
{{ .Snippet | pre }}
{{comment_md .Doc}}{{ range .Examples }}
#### <a name="example-{{ .Name | slugify }}">Example {{ .Suffix }}</a>
Code:
{{ .Snippet | pre}}{{ if .Output }}
Output:
{{ .Output | pre}}{{ end }}
{{ comment_md .Doc}}{{ end }}{{end}}{{ end }}{{ with .Examples }}
{{ range . }}
## <a name="example-{{ .Name | slugify }}">Example {{ .Suffix }}</a>
Code:
{{ .Snippet | pre}}{{ if .Output }}
Output:
{{ .Output | pre}}{{ end }}
{{ comment_md .Doc}}{{ end }}{{end}}
`
