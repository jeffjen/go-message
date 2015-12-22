package main

import (
	"github.com/jeffjen/go-message/push"

	log "github.com/Sirupsen/logrus"
	cli "github.com/codegangsta/cli"
	ctx "golang.org/x/net/context"

	"encoding/json"
	"fmt"
)

var (
	commonflag = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "json",
			Usage: "JSON message",
		},
		cli.StringSliceFlag{
			Name:  "msg",
			Usage: "Plain text message",
		},
	}

	handler push.Pusher
)

func process(c *cli.Context) {
	var (
		jsonlist = c.StringSlice("json")

		msglist = c.StringSlice("msg")
	)

	if len(jsonlist) != 0 {
		for _, d := range jsonlist {
			var v interface{}
			if err := json.Unmarshal([]byte(d), &v); err != nil {
				log.WithFields(log.Fields{"err": err, "d": d}).Fatal("abort")
			}
			if err := handler.Push(ctx.Background(), v); err != nil {
				log.WithFields(log.Fields{"err": err, "d": d}).Fatal("abort")
			}
		}
	} else if len(msglist) != 0 {
		for _, d := range msglist {
			if err := handler.Push(ctx.Background(), d); err != nil {
				log.WithFields(log.Fields{"err": err, "d": d}).Fatal("abort")
			}
		}
	} else {
		fmt.Println("nothing to do")
	}
}
