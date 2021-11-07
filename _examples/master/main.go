package main

import (
	"time"

	"github.com/open-farms/bot/pkg/gopigo"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
)

func setupRobot(master *gobot.Master) *gobot.Robot {
	rpi := raspi.NewAdaptor()
	driver := gopigo3.NewDriver(rpi)
	done := make(chan bool, 1)

	gobot.After(5*time.Second, func() {
		done <- true
	})

	work := func() {
		led := gopigo.NewLED(driver)
		err := led.Blink(1000*time.Millisecond, done)
		if err != nil {
			panic(err)
		}
	}

	robot := gobot.NewRobot("gopigo3",
		[]gobot.Connection{rpi},
		[]gobot.Device{driver},
		work,
	)

	return master.AddRobot(robot)
}

func setupAPI(master *gobot.Master) *api.API {
	a := api.NewAPI(master)
	a.Start()
	return a
}

func main() {
	master := gobot.NewMaster()
	setupRobot(master)
	setupAPI(master)

	master.Start()
}
