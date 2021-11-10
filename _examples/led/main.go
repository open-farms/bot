package main

import (
	"log"
	"time"

	"github.com/open-farms/bot/pkg/gopigo"
	"gobot.io/x/gobot"
)

func main() {
	bot := gopigo.NewRobot()
	work := func() {
		done := make(chan bool, 1)
		gobot.After(5*time.Second, func() {
			done <- true
		})
		err := bot.LED.Blink(1000*time.Millisecond, done)
		if err != nil {
			panic(err)
		}
	}

	robot := gobot.NewRobot("gopigo3",
		[]gobot.Connection{bot.Adaptor},
		[]gobot.Device{bot.Driver},
		work,
	)

	err := robot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
