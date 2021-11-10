package main

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/open-farms/bot/pkg/control"
	"github.com/open-farms/bot/pkg/gopigo"
	"github.com/open-farms/bot/pkg/pubsub"
	"gobot.io/x/gobot"
)

func main() {
	bot := gopigo.NewRobot()
	done := make(chan bool, 1)
	client, err := pubsub.NewClient(pubsub.PublicBroker, 1883)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(200, done)

	work := func() {
		client.Subscribe(pubsub.TopicControl, func(c mqtt.Client, m mqtt.Message) {
			pubsub.Logger.Info().Msg((string(m.Payload())))
			switch string(m.Payload()) {
			case control.MoveFront:
				client.Publish(pubsub.TopicControl, "received move front")
			case control.MoveBack:
				client.Publish(pubsub.TopicControl, "received move back")
			case control.MoveLeft:
				client.Publish(pubsub.TopicControl, "received move left")
			case control.MoveRight:
				client.Publish(pubsub.TopicControl, "received move right")
			default:
				return
			}
		})
		<-done
	}

	robot := gobot.NewRobot("gopigo3",
		[]gobot.Connection{bot.Adaptor},
		[]gobot.Device{bot.Driver},
		work,
	)

	err = robot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
