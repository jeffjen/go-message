package push

import (
	ctx "golang.org/x/net/context"
)

type Pusher interface {
	Push(c ctx.Context, data interface{}) error
}
