package main

import (
	cli "github.com/codegangsta/cli"
)

func init() {
	cli.AppHelpTemplate = `Usage: {{.Name}} [COMMANDS]

{{.Usage}}

Version: {{.Version}}

Commands:
	{{range .Commands}}{{.Name}}{{ "\t " }}{{.Usage}}
	{{end}}
`

	cli.CommandHelpTemplate = `Usage: {{.Name}} [OPTIONS]

{{.Usage}}

Options:
	{{range .Flags}}{{.}}
	{{end}}
`
}
