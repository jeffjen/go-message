package push

import (
	ctx "golang.org/x/net/context"

	tmpl "text/template"
)

type Pusher interface {
	// Push data to configured pusher
	Push(c ctx.Context, data interface{}) error

	// Template manipulation tool.  see text/template for instructions
	Template() *tmpl.Template
	Funcs(fm tmpl.FuncMap) *tmpl.Template
}

type NullPusher struct{}

func (n *NullPusher) Push(c ctx.Context, data interface{}) error {
	return nil
}

func (n *NullPusher) Template() *tmpl.Template {
	return nil
}

func (n *NullPusher) Funcs(fm tmpl.FuncMap) *tmpl.Template {
	return nil
}
