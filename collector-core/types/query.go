package types

import (
	"bytes"
	"text/template"
)

type Query struct {
	Command   string
	Type      string
	Condition func(*Context, interface{}) bool
}

type QueryTemplate struct {
	Template *template.Template
}

func (q *QueryTemplate) Render(name string, query string) {
	q.Template = template.New(name)
	q.Template = template.Must(q.Template.Parse(query))
}

func (q *QueryTemplate) Bind(qp *Context) string {
	var tpl bytes.Buffer
	q.Template.Execute(&tpl, qp)
	return tpl.String()
}

func MakeQueryTemplate(name string, query string) *QueryTemplate {
	template := QueryTemplate{}
	template.Render(name, query)
	return &template
}

/*
changelogQueryset := [1]core.Query{}
	changelogQueryset[0] = core.Query{Type: "Output", Command: changelogQuery}
*/

type QuerySet []Query
