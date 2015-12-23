package slack

import (
	ctx "golang.org/x/net/context"
	http "golang.org/x/net/context/ctxhttp"

	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	tmpl "text/template"
)

type Payload struct {
	Channel string `json:"channel"`
	Message string `json:"text"`
}

type Slack struct {
	// Slack URL endpoint
	SlackPushURL string

	// Template for push notification
	NotificationTmpl *tmpl.Template

	// Channel to push
	Channel string
}

func NewSlack(url, channel string, gen *tmpl.Template) *Slack {
	return &Slack{
		SlackPushURL:     url,
		NotificationTmpl: gen,
		Channel:          channel,
	}
}

func (s *Slack) Push(c ctx.Context, data interface{}) error {
	buf := new(bytes.Buffer)
	if err := s.NotificationTmpl.Execute(buf, data); err != nil {
		return err
	}
	msg := Payload{Channel: s.Channel, Message: buf.String()}
	return s.PushData(c, msg)
}

func (s *Slack) PushData(c ctx.Context, data Payload) error {
	if data.Channel == "" {
		data.Channel = s.Channel
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return err
	}
	resp, err := http.Post(c, nil, s.SlackPushURL, "application/json", buf)
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

func (s *Slack) Funcs(fm tmpl.FuncMap) *tmpl.Template {
	return s.NotificationTmpl.Funcs(fm)
}

func (s *Slack) Template() *tmpl.Template {
	return s.NotificationTmpl
}
