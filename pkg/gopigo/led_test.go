package gopigo_test

import (
	"time"

	"github.com/open-farms/bot/pkg/gopigo"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
)

func ExampleNewLED_debug() {
	rpi := raspi.NewAdaptor()
	driver := gopigo3.NewDriver(rpi)
	done := make(chan bool, 1)

	gobot.After(5*time.Second, func() {
		done <- true
	})

	work := func() {
		led := gopigo.NewLED(driver,
			gopigo.WithDebug(),
		)
		err := led.Blink(1*time.Second, done)
		if err != nil {
			panic(err)
		}
	}

	robot := gobot.NewRobot("gopigo",
		[]gobot.Connection{rpi},
		[]gobot.Device{driver},
		work,
	)

	robot.Start()
}
