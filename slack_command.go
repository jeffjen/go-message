package main

import (
	scli "github.com/jeffjen/go-message/push/slack/cli"

	cli "github.com/codegangsta/cli"
)

func NewSlackCommand() cli.Command {
	return cli.Command{
		Name:   "slack",
		Usage:  "Send message using slack",
		Flags:  append(commonflag, scli.Flags...),
		Before: slack_prepare,
		Action: process,
	}
}

func slack_prepare(c *cli.Context) error {
	if err := scli.Before(c); err != nil {
		return err
	} else {
		handler = scli.DefaultSlack
		return nil
	}
}
