//
// URL:  https://github.com/sfmunoz/golang-playground
// Date: Sun Jul 27 03:59:18 PM UTC 2025
//

package templates

import (
	"os"
	"text/template"
)

type tpl struct {
	Name  string
	Value string
}

var headerTpl string = `{{ define "header" }}Header: {{ .Header }}{{ end }}`

var footerTpl string = `{{ define "footer" }}Footer: {{ .Footer }}{{ end }}`

var mainTpl string = `{{ template "header" . }}
{{ block "title" . }}Title: {{ .Title }}{{ end }}
{{ range .People -}}
- {{ .Name }}: {{ .Age }}
{{ end -}}
{{ template "footer" . }}
`

type Person struct {
	Name string
	Age  int
}

func newTemplate() (*template.Template, error) {
	t, err := template.New("main").Parse(mainTpl)
	if err != nil {
		return nil, err
	}
	t, err = t.Parse(headerTpl) // name defined by the template
	if err != nil {
		return nil, err
	}
	t, err = t.Parse(footerTpl) // name defined by the template
	if err != nil {
		return nil, err
	}
	t = t.Lookup("main") // not needed: it's already active
	return t, nil
}

func run() error {
	t, err := newTemplate()
	if err != nil {
		return err
	}
	return t.Execute(
		os.Stdout,
		map[string]any{
			"Header": "Name and Age",
			"Title":  "People list",
			"Footer": "the end",
			"People": []Person{{"Jane", 24}, {"John", 23}, {"Sam", 22}},
		},
	)
}

func Main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
