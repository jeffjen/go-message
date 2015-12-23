package cli

import (
	cli "github.com/codegangsta/cli"

	tmpl "text/template"
)

var (
	NotificationTmpl = "{{.}}"

	FuncMap tmpl.FuncMap
)

var (
	Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "url",
			EnvVar: "SLACK_URL",
			Usage:  "Web Integration endpoint for Slack POST message",
		},
		cli.StringFlag{
			Name:   "channel",
			Value:  "#random",
			EnvVar: "SLACK_CHANNEL",
			Usage:  "Slack channel to POST message to",
		},
		cli.StringFlag{
			Name:   "tmpl",
			EnvVar: "SLACK_TMPL",
			Usage:  "Slack message template",
		},
	}
)
