package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"

	"github.com/open-farms/bot/pkg/move"
	"github.com/open-farms/bot/pkg/move/controls"
	"github.com/open-farms/bot/pkg/pubsub"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/keyboard"
)

func setupTLS(pemfile string) *tls.Config {
	caPEM, err := os.ReadFile(pemfile)
	if err != nil {
		panic(err)
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caPEM)
	if !ok {
		panic("failed to parse root certificate")
	}

	return &tls.Config{RootCAs: roots}
}

func main() {
	// cert := setupTLS("./config/certs/ca.pem")
	client := pubsub.NewClient(pubsub.PUBLIC_BROKER, 1883, nil)
	k := controls.NewKeyboard(client)
	ctl := move.New(k)
	work := func() {
		k.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)
			switch key.Key {
			case keyboard.W:
				ctl.Move(move.Forward)
			case keyboard.ArrowUp:
				ctl.Move(move.Forward)
			case keyboard.S:
				ctl.Move(move.Backward)
			case keyboard.ArrowDown:
				ctl.Move(move.Backward)
			case keyboard.A:
				ctl.Move(move.Left)
			case keyboard.ArrowLeft:
				ctl.Move(move.Left)
			case keyboard.D:
				ctl.Move(move.Right)
			case keyboard.ArrowRight:
				ctl.Move(move.Right)
			case keyboard.Spacebar:
				ctl.Move(move.Stop)
			case keyboard.Escape:
				ctl.Move(move.Stop)
			}
		})
	}

	robot := gobot.NewRobot("keyboard",
		[]gobot.Connection{},
		[]gobot.Device{k.Driver},
		work,
	)

	if err := robot.Start(); err != nil {
		log.Fatal(err)
	}
}
