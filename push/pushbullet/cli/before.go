package cli

import (
	"github.com/jeffjen/go-message/push/pushbullet"

	cli "github.com/codegangsta/cli"

	tmpl "text/template"
)

var (
	DefaultPushBullet *pushbullet.PushBullet
)

func Before(c *cli.Context) error {
	token := c.String("token")
	if token == "" {
		return nil
	}

	pushbullet_msg_tmpl := c.String("tmpl")
	gen, err := tmpl.New("pushbullet").Parse(pushbullet_msg_tmpl)
	if err != nil {
		return err
	}

	iden := pushbullet.Identity{
		AccessToken: token,
		Email:       c.String("email"),
		Iden:        c.String("device-iden"),
	}

	DefaultPushBullet = pushbullet.NewPushBullet(iden, c.String("title"), gen)
	return nil
}
