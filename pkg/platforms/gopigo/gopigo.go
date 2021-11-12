package gopigo

import (
	"gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
)

type Gopigo struct {
	Driver  *gopigo3.Driver
	Adaptor *raspi.Adaptor

	*LED
	*Motor
}

const (
	MOTOR_ERR      = "motor_error"
	MOTOR_STOP     = "motor_stop"
	MOTOR_BACKWARD = "motor_backward"
	MOTOR_FORWARD  = "motor_forward"
	MOTOR_LEFT     = "motor_left"
	MOTOR_RIGHT    = "motor_right"

	LED_ERR = "led_error"
	LED_OFF = "led_off"
	LED_ON  = "led_on"

	BLINKER_ERR = "blinker_error"
	BLINKER_OFF = "blinker_off"
	BLINKER_ON  = "blinker_on"
)

func New() *Gopigo {
	rpi := raspi.NewAdaptor()
	driver := gopigo3.NewDriver(rpi)
	return &Gopigo{
		Driver:  driver,
		Adaptor: rpi,
		LED:     NewLED(driver),
		Motor:   NewMotor(driver),
	}
}
