package controls

import (
	"github.com/open-farms/bot/pkg/move"
	"github.com/open-farms/bot/pkg/pubsub"
	"gobot.io/x/gobot/platforms/joystick"
)

type Joystick struct {
	Client *pubsub.Client
	*joystick.Driver
}

type JoystickVariant string

const (
	Dualshock3 JoystickVariant = "dualshock3"
	Dualshock4 JoystickVariant = "dualshock4"
	Shield     JoystickVariant = "shield"
	Xbox360    JoystickVariant = "xbox360"
)

func NewJoystick(c *pubsub.Client, variant JoystickVariant) *Joystick {
	return &Joystick{
		Client: c,
		Driver: joystick.NewDriver(joystick.NewAdaptor(), string(variant)),
	}
}

func (j *Joystick) Move(d move.Direction) {
	j.Client.Publish(pubsub.TopicControl, d.String())
}
