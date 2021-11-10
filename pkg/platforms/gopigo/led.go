package gopigo

import (
	"image/color"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dexter/gopigo3"
	"golang.org/x/image/colornames"
)

// LED is the led service object
type LED struct {
	driver *gopigo3.Driver
	color  color.RGBA
	debug  bool
	logger zerolog.Logger
}

type LEDOption func(*LED)

// NewLED creates a new gopigo LED object for manipulating
// the bots lights
func NewLED(driver *gopigo3.Driver, options ...LEDOption) *LED {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	var (
		defaultColor  = colornames.Royalblue
		defaultDebug  = false
		defaultLogger = zerolog.New(os.Stderr)
	)

	led := &LED{
		driver: driver,
		color:  defaultColor,
		debug:  defaultDebug,
		logger: defaultLogger,
	}

	for _, option := range options {
		option(led)
	}
	return led
}

// WithDebug enables debug logging for the led driver
func WithDebug() LEDOption {
	return func(l *LED) {
		l.debug = true
	}
}

// ApplyColor sets the color of the gopigo's LED
func (l *LED) ApplyColor(c color.RGBA, light gopigo3.Led) error {
	l.color = c
	log.Debug().Str("func", "applycolor").Msgf("setting led %q to color %q", light, c)
	err := l.driver.SetLED(light, l.color.R, l.color.G, l.color.B)
	if err != nil {
		return err
	}

	return nil
}

// Blink the lights on the robot at a specified interval
func (l *LED) Blink(frequency time.Duration, done chan bool) error {
	on := uint8(0xFF)
	ticker := gobot.Every(frequency, func() {
		log.Debug().Str("func", "blink").Msg("blinking left led")
		err := l.driver.SetLED(gopigo3.LED_EYE_LEFT, 0x00, 0x00, on)
		if err != nil {
			panic(err)
		}

		log.Debug().Str("func", "blink").Msg("blinking right led")
		err = l.driver.SetLED(gopigo3.LED_EYE_RIGHT, 0x00, 0x00, on)
		if err != nil {
			panic(err)
		}
		on = ^on
	})

	<-done
	ticker.Stop()

	return nil
}

// Wink a light on the robot at a specified interval
func (l *LED) Wink(frequency time.Duration, light gopigo3.Led) error {
	on := uint8(0xFF)

	log.Debug().Str("func", "wink").Msgf("setting %q to wink at interval %q", light, frequency)
	gobot.Every(frequency, func() {
		err := l.driver.SetLED(light, 0x00, 0x00, on)
		if err != nil {
			panic(err)
		}
		on = ^on
	})

	return nil
}
