package controls

// import (
// 	"github.com/open-farms/bot/pkg/move"
// 	"github.com/open-farms/bot/pkg/pubsub"
// 	"gobot.io/x/gobot/platforms/joystick"
// )

// type Joystick struct {
// 	Client *pubsub.Client
// 	*joystick.Driver
// }

// func NewJoystick(c *pubsub.Client, config string) *Joystick {
// 	return &Joystick{
// 		Client: c,
// 		Driver: joystick.NewDriver(joystick.NewAdaptor(), config),
// 	}
// }

// func (j *Joystick) Move(d move.Direction) {
// 	j.Client.Publish(pubsub.TopicControl, d.String())
// }
