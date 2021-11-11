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

	client := pubsub.NewClient(pubsub.PublicBroker, 1883, nil)
	work := func() {
		client.Subscribe(pubsub.TopicControl, func(c mqtt.Client, m mqtt.Message) {
			payload := string(m.Payload())
			logger.Log.Info().Msg(payload)
			switch payload {
			case move.Forward.String():
				bot.Motor.Forward(360)
			case move.Backward.String():
				bot.Motor.Backward(360)
			case move.Left.String():
				bot.Motor.Left(360)
			case move.Right.String():
				bot.Motor.Right(360)
			case move.Stop.String():
				bot.Motor.Stop()
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

	if err := robot.Start(); err != nil {
		log.Fatal(err)
	}
}
