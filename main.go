package main

import (
	cli "github.com/codegangsta/cli"

	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-message"
	app.Usage = "Very simple message push helper for Go"
	app.Authors = []cli.Author{
		cli.Author{"Yi-Hung Jen", "yihungjen@gmail.com"},
	}
	app.Commands = []cli.Command{
		NewSlackCommand(),
		NewPushBulletCommand(),
	}
	app.Run(os.Args)
}
