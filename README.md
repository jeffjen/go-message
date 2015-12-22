# go-message
Very simple message push helper for Go

### tl;dr
- `go-message slack --url [WEB_INTEGRATION_URL] --msg "Call of the postbox"`

- `go-message pushbullet --token [ACCESS_TOKEN] --msg "Reverse call of the void"`

### About template
Using `text/template` for our template engine

### Options
```
Usage: go-message [COMMANDS]

Very simple message push helper for Go

Version: 0.0.0

Commands:
        slack            Send message using slack
        pushbullet       Send message using pushbullet
        help             Shows a list of commands or help for one command
```

```
Usage: slack [OPTIONS]

Send message using slack

Options:
    --json [--json option --json option]    JSON message
    --msg [--msg option --msg option]       Plain text message
    --url                                   Web Integration endpoint for Slack POST message [$SLACK_URL]
    --channel "#general"                    Slack channel to POST message to [$SLACK_CHANNEL]
    --tmpl "{{.}}"                          Slack message template [$SLACK_TMPL]
```

```
Usage: pushbullet [OPTIONS]

Send message using pushbullet

Options:
    --json [--json option --json option]    JSON message
    --msg [--msg option --msg option]       Plain text message
    --token                                 Access token for push action [$PUSHBULLET_ACCESS_TOKEN]
    --title "ping"                          Push message title [$PUSHBULLET_TITLE]
    --tmpl "{{.}}"                          PushBullet message template [$PUSHBULLET_TMPL]
    --email                                 Email identity for Pushbullet to send to [$PUSHBULLET_EMAIL]
    --device-iden                           Device identity for Pushbullet to send to [$PUSHBULLET_DEVICE_IDEN]
```
