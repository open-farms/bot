package main

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/open-farms/bot/internal/logger"
	"github.com/open-farms/bot/pkg/move"
	"github.com/open-farms/bot/pkg/platforms/gopigo"
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
			payload := string(m.Payload())
			logger.Log.Info().Msg(payload)

			switch payload {
			case move.Forward.String():
				client.Publish(pubsub.TopicControl, "received move forward")
			case move.Backward.String():
				client.Publish(pubsub.TopicControl, "received move backward")
			case move.Backward.String():
				client.Publish(pubsub.TopicControl, "received move left")
			case move.Right.String():
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
