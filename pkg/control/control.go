package control

import (
	"github.com/open-farms/bot/pkg/pubsub"
)

const (
	MoveFront = "front"
	MoveBack  = "back"
	MoveLeft  = "left"
	MoveRight = "right"
)

func Front(client *pubsub.Client) {
	client.Publish(pubsub.TopicControl, MoveFront)
}

func Back(client *pubsub.Client) {
	client.Publish(pubsub.TopicControl, MoveBack)
}

func Left(client *pubsub.Client) {
	client.Publish(pubsub.TopicControl, MoveLeft)
}

func Right(client *pubsub.Client) {
	client.Publish(pubsub.TopicControl, MoveRight)
}
