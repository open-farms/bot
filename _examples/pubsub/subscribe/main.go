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
	gpg := gopigo.New()
	done := make(chan bool, 1)

	client := pubsub.NewClient(pubsub.PUBLIC_BROKER, 1883, nil)
	work := func() {
		events := gpg.Subscribe()
		go func() {
			event := <-events
			logger.Log.Info().Str("service", "bot").Str("event", event.Name).Interface("speed", event.Data).Send()
			client.Publish(pubsub.TOPIC_INFO, event.Name)
		}()

		client.Subscribe(pubsub.TOPIC_CONTROL, func(c mqtt.Client, m mqtt.Message) {
			payload := string(m.Payload())
			switch payload {
			case move.Forward.String():
				gpg.Motor.Forward(360)

			case move.Backward.String():
				gpg.Motor.Backward(360)

			case move.Left.String():
				gpg.Motor.Left(360)

			case move.Right.String():
				gpg.Motor.Right(360)

			case move.Stop.String():
				gpg.Motor.Stop()

			default:
				return
			}
		})
		<-done
	}

	robot := gobot.NewRobot("gopigo3",
		[]gobot.Connection{gpg.Adaptor},
		[]gobot.Device{gpg.Driver},
		work,
	)

	if err := robot.Start(); err != nil {
		log.Fatal(err)
	}
}
