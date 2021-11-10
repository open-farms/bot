package gopigo

import (
	"gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
)

type Robot struct {
	Driver  *gopigo3.Driver
	Adaptor *raspi.Adaptor

	LED   *LED
	Motor *Motor
}

func NewRobot() *Robot {
	rpi := raspi.NewAdaptor()
	driver := gopigo3.NewDriver(rpi)

	return &Robot{
		Driver:  driver,
		Adaptor: rpi,
		LED:     NewLED(driver),
		Motor:   NewMotor(driver),
	}
}
