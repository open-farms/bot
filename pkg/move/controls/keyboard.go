package controls

import (
	"github.com/open-farms/bot/pkg/move"
	"github.com/open-farms/bot/pkg/pubsub"
	"gobot.io/x/gobot/platforms/keyboard"
)

type Keyboard struct {
	Client *pubsub.Client
	*keyboard.Driver
}

func NewKeyboard(c *pubsub.Client) *Keyboard {
	return &Keyboard{
		Client: c,
		Driver: keyboard.NewDriver(),
	}
}

func (k *Keyboard) Move(d move.Direction) {
	k.Client.Publish(pubsub.TopicControl, d.String())
}
