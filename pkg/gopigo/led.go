package gopigo

import (
	"context"
	"image/color"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dexter/gopigo3"
	"golang.org/x/image/colornames"
)

type LED struct {
	driver *gopigo3.Driver
	led    gopigo3.Led
	color  color.RGBA
}

// NewLED creates a new gopigo LED object for manipulating
// the bots lights
func NewLED(driver *gopigo3.Driver, led gopigo3.Led) *LED {
	return &LED{
		driver: driver,
		led:    led,
		color:  colornames.Royalblue,
	}
}

// ApplyColor sets the color of the gopigo's LED
func (l *LED) ApplyColor(c color.RGBA) error {
	l.color = c
	err := l.driver.SetLED(l.led, l.color.R, l.color.G, l.color.B)
	if err != nil {
		return err
	}

	return nil
}

// Blink the lights on the robot at a specified interval
func (l *LED) Blink(ctx context.Context, frequency time.Duration) error {
	on := uint8(0xFF)
	gobot.Every(frequency, func() {
		err := l.driver.SetLED(gopigo3.LED_EYE_RIGHT, 0x00, 0x00, on)
		if err != nil {
			panic(err)
		}

		err = l.driver.SetLED(gopigo3.LED_EYE_RIGHT, 0x00, 0x00, on)
		if err != nil {
			panic(err)
		}
		on = ^on
	})

	return nil
}
