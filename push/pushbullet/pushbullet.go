package pushbullet

import (
	ctx "golang.org/x/net/context"
	http "golang.org/x/net/context/ctxhttp"

	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	net_http "net/http"
	tmpl "text/template"
)

const (
	PushBulletURL = "https://api.pushbullet.com/v2/pushes"
)

type Payload struct {
	Type       string `json:"type"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Email      string `json:"email,omitempty"`
	DeviceIden string `json:"device_iden,omitempty"`
}

type Identity struct {
	// AccessToken for pushing
	AccessToken string

	// Identity for target to receive
	Iden string

	// Email for target to receive, either to their email, or to their account
	Email string
}

type PushBullet struct {
	Iden Identity

	// Title for the push
	Title string

	// Template for push notification
	NotificationTmpl *tmpl.Template
}

func NewPushBullet(iden Identity, title string, gen *tmpl.Template) *PushBullet {
	return &PushBullet{
		Iden:             iden,
		Title:            title,
		NotificationTmpl: gen,
	}
}

func (p *PushBullet) Push(c ctx.Context, data interface{}) error {
	buf := new(bytes.Buffer)
	if err := p.NotificationTmpl.Execute(buf, data); err != nil {
		return err
	}
	msg := Payload{
		Type:  "note",
		Title: p.Title,
		Body:  buf.String(),
	}
	if p.Iden.Iden != "" {
		msg.DeviceIden = p.Iden.Iden
	} else if p.Iden.Email != "" {
		msg.Email = p.Iden.Email
	} else {
		// NOTE: default to  broadcast mode
		// probably not the best but allowd by pushbullet API
	}
	return p.PushData(c, msg)
}

func (p *PushBullet) PushData(c ctx.Context, data Payload) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return err
	}
	req, err := net_http.NewRequest("POST", PushBulletURL, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Access-Token", p.Iden.AccessToken)

	resp, err := http.Do(c, nil, req)
	if err != nil {
		return err
	} else {
		go func() {
			defer resp.Body.Close()
			io.Copy(ioutil.Discard, resp.Body)
		}()
		return nil
	}
}

func (p *PushBullet) Funcs(fm tmpl.FuncMap) *tmpl.Template {
	return p.NotificationTmpl.Funcs(fm)
}

func (p *PushBullet) Template() *tmpl.Template {
	return p.NotificationTmpl
}
