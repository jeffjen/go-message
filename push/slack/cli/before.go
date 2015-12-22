package cli

import (
	"github.com/jeffjen/go-message/push/slack"

	cli "github.com/codegangsta/cli"

	tmpl "text/template"
)

var (
	DefaultSlack *slack.Slack
)

func Before(c *cli.Context) error {
	slack_url := c.String("url")
	if slack_url == "" {
		return nil
	}

	channel := c.String("channel")

	slack_msg_tmpl := c.String("tmpl")
	if slack_msg_tmpl == "" {
		slack_msg_tmpl = DefaultNotificationTmpl
	}
	gen, err := tmpl.New("slack").Parse(slack_msg_tmpl)
	if err != nil {
		return err
	}

	DefaultSlack = slack.NewSlack(slack_url, channel, gen)
	return nil
}
