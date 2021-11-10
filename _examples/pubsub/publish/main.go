package main

import (
	"log"

	"github.com/open-farms/bot/pkg/pubsub"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/keyboard"
)

func main() {
	client, err := pubsub.NewClient(pubsub.PublicBroker, 1883)
	if err != nil {
		log.Fatal(err)
	}

	keys := keyboard.NewDriver()
	work := func() {
		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)
			switch key.Key {
			case keyboard.W:
				client.Publish(pubsub.TopicControl, "front")
			case keyboard.S:
				client.Publish(pubsub.TopicControl, "back")
			case keyboard.A:
				client.Publish(pubsub.TopicControl, "left")
			case keyboard.D:
				client.Publish(pubsub.TopicControl, "right")
			}
		})
	}

	robot := gobot.NewRobot("keyboard",
		[]gobot.Connection{},
		[]gobot.Device{keys},
		work,
	)

	err = robot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
