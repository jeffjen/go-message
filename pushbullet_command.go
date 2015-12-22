package main

import (
	pcli "github.com/jeffjen/go-message/push/pushbullet/cli"

	cli "github.com/codegangsta/cli"
)

func NewPushBulletCommand() cli.Command {
	return cli.Command{
		Name:   "pushbullet",
		Usage:  "Send message using pushbullet",
		Flags:  append(commonflag, pcli.Flags...),
		Before: pushbullet_prepare,
		Action: process,
	}
}

func pushbullet_prepare(c *cli.Context) error {
	if err := pcli.Before(c); err != nil {
		return err
	} else {
		handler = pcli.DefaultPushBullet
		return nil
	}
}
