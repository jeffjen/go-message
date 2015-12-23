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
			Name:   "token",
			EnvVar: "PUSHBULLET_ACCESS_TOKEN",
			Usage:  "Access token for push action",
		},
		cli.StringFlag{
			Name:   "title",
			Value:  "ping",
			EnvVar: "PUSHBULLET_TITLE",
			Usage:  "Push message title",
		},
		cli.StringFlag{
			Name:   "tmpl",
			EnvVar: "PUSHBULLET_TMPL",
			Usage:  "PushBullet message template",
		},
		cli.StringFlag{
			Name:   "email",
			EnvVar: "PUSHBULLET_EMAIL",
			Usage:  "Email identity for Pushbullet to send to",
		},
		cli.StringFlag{
			Name:   "device-iden",
			EnvVar: "PUSHBULLET_DEVICE_IDEN",
			Usage:  "Device identity for Pushbullet to send to",
		},
	}
)
